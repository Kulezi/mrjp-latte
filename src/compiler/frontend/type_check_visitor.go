package frontend

import (
	"fmt"
	"latte/parser"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type varDropper struct {
	drop  func()
	depth int
}

type doesReturn struct {
	always, sometimes bool
}

// typeCheckVisitor methods evaluating statements return doesReturn or an error.
// Methods evaluating expressions return Type or an error.
type typeCheckVisitor struct {
	parser.BaseLatteVisitor
	state        *state
	depth        int
	dropperStack []varDropper
	curClass     *TClass
	curFun       *TFun
}

func makeTypeCheckVisitor(s *state) *typeCheckVisitor {
	return &typeCheckVisitor{state: s}
}

func (v *typeCheckVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *typeCheckVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	res := make([]interface{}, 0)
	for _, child := range ctx.AllTopDef() {
		if err, ok := v.Visit(child).(error); ok {
			return err
		}
	}

	return res
}

func (v *typeCheckVisitor) isSubClass(expected, child Type) bool {
	class, ok := child.(TClass)
	if !ok {
		return false
	}

	for !sameType(expected, class) && class.Parent != nil {
		parentRef := *class.Parent
		parent, _ := v.TypeOfGlobal(parentRef.String())
		class = parent.Type.(TClass)
	}

	return sameType(expected, class)
}

func (v *typeCheckVisitor) ExpectType(expected Type, ctx parser.IExprContext) error {
	if ctx == nil {
		return nil
	}

	got, err := v.evalType(ctx)
	if err != nil {
		return err
	}
	if !sameType(expected, got) && !v.isSubClass(expected, got) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: expected,
			Got:      got,
		}
	}

	return nil
}

func (v *typeCheckVisitor) EnterClass(signature TClass) (exit func()) {
	// fmt.Printf("enter class %s\n", signature)
	v.depth++
	// Put all fields into enviroment.
	var exits []func()
	for ident, typ := range signature.Fields {
		exits = append(exits, v.ShadowLocal(ident, typ))
	}

	return func() {
		// fmt.Printf("exit class %s\n", signature)
		for _, exit := range exits {
			exit()
		}
		v.depth--
	}
}

func (v *typeCheckVisitor) VisitTopDef(ctx *parser.TopDefContext) interface{} {
	if fun := ctx.Fundef(); fun != nil {
		return v.Visit(fun)
	}

	return v.Visit(ctx.Classdef())
}

func (v *typeCheckVisitor) ConflictLocal(ident string) (TypeInfo, bool) {
	return v.state.signatures.ConflictLocal(ident, v.depth)
}

func (v *typeCheckVisitor) ShadowLocal(ident string, typ Type) (drop func()) {
	return v.state.signatures.ShadowLocal(ident, typ, v.depth)
}

func (v *typeCheckVisitor) TypeOfLocal(ident string) (TypeInfo, bool) {
	res, ok := v.state.signatures.Locals[ident]
	return res, ok
}

func (v *typeCheckVisitor) TypeOfGlobal(ident string) (TypeInfo, bool) {
	res, ok := v.state.signatures.Globals[ident]
	return res, ok
}

func (v *typeCheckVisitor) TypeOf(ident string) (TypeInfo, bool) {
	if res, ok := v.TypeOfLocal(ident); ok {
		return res, ok
	}

	return v.TypeOfGlobal(ident)
}

func (v *typeCheckVisitor) VisitFunDef(ctx *parser.FunDefContext) interface{} {
	ident := ctx.ID().GetText()

	t, ok := v.TypeOf(ident)
	if !ok {
		panic(fmt.Sprintf("undeclared identifier %s found at %s", ident, posFromToken(ctx.GetStart())))
	}

	signature, ok := t.Type.(TFun)
	if !ok {
		panic(fmt.Sprintf("typecheck: identifier %s is not a function/method, at %s", ident, posFromToken(ctx.GetStart())))
	}

	if _, ok := signature.Result.(TClassRef); ok {
		if _, ok := v.TypeOfGlobal(signature.Result.String()); !ok {
			return UnknownClassError{
				Type: signature.Result,
			}
		}
	}

	v.depth++
	for _, arg := range signature.Args {
		defer v.ShadowLocal(arg.Ident, arg.Type)()
	}
	v.depth--

	v.curFun = &signature
	defer func() { v.curFun = nil }()
	res := v.Visit(ctx.Block())
	if err, ok := res.(error); ok {
		return err
	}

	returns := res.(doesReturn)
	if !returns.always && !sameType(TVoid{}, signature.Result) {
		return MissingReturnError{
			Fun: signature,
		}
	}

	return nil
}

func (v *typeCheckVisitor) VisitBaseClassDef(ctx *parser.BaseClassDefContext) interface{} {
	t, ok := v.TypeOfGlobal(ctx.ID().GetText())
	if !ok {
		panic(fmt.Sprintf("typecheck: found undeclared class %s at %s", ctx.ID().GetText(), posFromToken(ctx.GetStart())))
	}

	signature, ok := t.Type.(TClass)
	if !ok {
		panic(fmt.Sprintf("typecheck: identifier %s is not a class, at %s", ctx.ID().GetText(), posFromToken(ctx.GetStart())))
	}

	defer v.EnterClass(signature)()
	v.curClass = &signature
	// Evaluate methods
	for _, field := range ctx.AllField() {
		if err, ok := v.Visit(field).(error); ok {
			return err
		}
	}

	v.curClass = nil
	return nil
}

func (v *typeCheckVisitor) VisitDerivedClassDef(ctx *parser.DerivedClassDefContext) interface{} {
	t, ok := v.TypeOfGlobal(ctx.ID(0).GetText())
	if !ok {
		panic(fmt.Sprintf("typecheck: found undeclared class %s at %s", ctx.ID(0).GetText(), posFromToken(ctx.GetStart())))
	}

	signature, ok := t.Type.(TClass)
	if !ok {
		panic(fmt.Sprintf("typecheck: identifier %s is not a class, at %s", ctx.ID(0).GetText(), posFromToken(ctx.GetStart())))
	}

	defer v.EnterClass(signature)()

	v.curClass = &signature
	// Evaluate methods
	for _, field := range ctx.AllField() {
		if err, ok := v.Visit(field).(error); ok {
			return err
		}
	}

	v.curClass = nil
	return nil
}

func (v *typeCheckVisitor) VisitClassFieldDef(ctx *parser.ClassFieldDefContext) interface{} {
	return v.Visit(ctx.Nvtype_())
}

func (v *typeCheckVisitor) VisitClassMethodDef(ctx *parser.ClassMethodDefContext) interface{} {
	return v.Visit(ctx.Fundef())
}

func (v *typeCheckVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.depth++
	defer func() { v.depth-- }()
	returns := doesReturn{}
	for _, stmt := range ctx.AllStmt() {
		stmtReturns, err := v.evalStmt(stmt)
		if err != nil {
			return err
		}
		returns = doesReturn{
			always:    returns.always || stmtReturns.always,
			sometimes: returns.sometimes || stmtReturns.sometimes,
		}
	}

	for len(v.dropperStack) > 0 {
		dropper := v.dropperStack[len(v.dropperStack)-1]
		if dropper.depth != v.depth {
			break
		}

		dropper.drop()
		v.dropperStack = v.dropperStack[:len(v.dropperStack)-1]
	}

	return returns
}

func (v *typeCheckVisitor) evalStmt(ctx parser.IStmtContext) (returns doesReturn, err error) {
	ret := v.Visit(ctx)
	if err, ok := ret.(error); ok {
		return doesReturn{}, err
	}

	b, ok := ret.(doesReturn)
	if !ok {
		panic("all stmt's should evaluate to doesReturn")
	}

	return b, nil
}

func (v *typeCheckVisitor) VisitSEmpty(ctx *parser.SEmptyContext) interface{} {
	return doesReturn{}
}

func (v *typeCheckVisitor) VisitSBlockStmt(ctx *parser.SBlockStmtContext) interface{} {
	return v.Visit(ctx.Block())
}

func (v *typeCheckVisitor) VisitSDecl(ctx *parser.SDeclContext) interface{} {
	t := v.Visit(ctx.Nvtype_())
	if err, ok := t.(error); ok {
		return err
	}

	typ := t.(Type)
	for _, item := range ctx.AllItem() {
		item, ok := item.(*parser.ItemContext)
		if !ok {
			panic("unexpected antlr behaviour: item is not an ItemContext")
		}
		ident := item.ID()
		if clashing, ok := v.ConflictLocal(ident.GetText()); ok {
			return DuplicateIdentifierError{
				Ident: ident.GetText(),
				Pos1:  clashing.Position(),
				Pos2:  posFromToken(ident.GetSymbol()),
			}
		}

		if item.Expr() != nil {
			if err := v.ExpectType(typ, item.Expr()); err != nil {
				return err
			}
		}
		v.dropperStack = append(v.dropperStack, varDropper{
			drop:  v.ShadowLocal(ident.GetText(), typ),
			depth: v.depth,
		})
	}
	return doesReturn{}
}

func (v *typeCheckVisitor) VisitLVField(ctx *parser.LVFieldContext) interface{} {
	t, err := v.evalType(ctx.Expr())
	if err != nil {
		return err
	}

	if class, ok := t.(TClass); ok {
		defer v.EnterClass(class)()
	}

	return v.Visit(ctx.Lvalue())
}

func (v *typeCheckVisitor) VisitLVArrayRef(ctx *parser.LVArrayRefContext) interface{} {
	typ, err := v.evalType(ctx.Expr(0))
	if err != nil {
		return err
	}

	arr, ok := typ.(TArray)
	if !ok {
		return ExpectedArrayError{
			Expr: ctx,
			Got:  typ,
		}
	}

	if err := v.ExpectType(TInt{}, ctx.Expr(1)); err != nil {
		return err
	}
	return arr.Elem
}

func (v *typeCheckVisitor) VisitLVId(ctx *parser.LVIdContext) interface{} {
	t, ok := v.TypeOfLocal(ctx.ID().GetText())
	if !ok {
		return UndeclaredIdentifierError{
			Ident: ctx.ID(),
		}
	}

	return t.Type
}

func (v *typeCheckVisitor) VisitSAss(ctx *parser.SAssContext) interface{} {
	t, err := v.evalLVType(ctx.Lvalue())
	if err != nil {
		return err
	}
	if err := v.ExpectType(t, ctx.Expr()); err != nil {
		return err
	}
	return doesReturn{}
}

func (v *typeCheckVisitor) VisitSIncr(ctx *parser.SIncrContext) interface{} {
	t, err := v.evalLVType(ctx.Lvalue())
	if err != nil {
		return err
	}

	if !sameType(t, TInt{}) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TInt{},
			Got:      t,
		}
	}

	return doesReturn{}
}

func (v *typeCheckVisitor) VisitSDecr(ctx *parser.SDecrContext) interface{} {
	t, err := v.evalLVType(ctx.Lvalue())
	if err != nil {
		return err
	}

	if !sameType(t, TInt{}) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TInt{},
			Got:      t,
		}
	}

	return doesReturn{}
}

func (v *typeCheckVisitor) VisitSRet(ctx *parser.SRetContext) interface{} {
	if sameType(v.curFun.Result, TVoid{}) {
		return VoidReturnWithValueError{
			Ctx: ctx,
			Fun: *v.curFun,
		}
	}

	if err := v.ExpectType(v.curFun.Result, ctx.Expr()); err != nil {
		return err
	}

	return doesReturn{
		always:    true,
		sometimes: true,
	}
}

func (v *typeCheckVisitor) VisitSVRet(ctx *parser.SVRetContext) interface{} {
	if !sameType(v.curFun.Result, TVoid{}) {
		return MissingReturnValueError{
			Ctx:      ctx,
			Expected: v.curFun.Result,
		}
	}
	return doesReturn{
		always:    true,
		sometimes: true,
	}
}

func (v *typeCheckVisitor) evalNonDeclStmt(ctx parser.IStmtContext) (doesReturn, error) {
	if _, ok := ctx.(*parser.SDeclContext); ok {
		return doesReturn{}, DeclarationWithoutBlockError{
			Ctx: ctx,
		}
	}

	return v.evalStmt(ctx)
}

func (v *typeCheckVisitor) VisitSCond(ctx *parser.SCondContext) interface{} {
	t, err := v.evalType(ctx.Expr())
	if err != nil {
		return err
	}

	b, ok := t.(TBool)
	if !ok {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TBool{},
			Got:      t,
		}
	}

	blockReturns, err := v.evalNonDeclStmt(ctx.Stmt())
	if err != nil {
		return err
	}

	// If it always executes we are sure of the return.
	if b.constValue != nil && *b.constValue {
		return blockReturns
	}

	return doesReturn{
		always:    false,
		sometimes: blockReturns.sometimes,
	}
}

func (v *typeCheckVisitor) VisitSCondElse(ctx *parser.SCondElseContext) interface{} {
	t, err := v.evalType(ctx.Expr())
	if err != nil {
		return err
	}

	b, ok := t.(TBool)
	if !ok {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TBool{},
			Got:      t,
		}
	}

	retTrue, err := v.evalNonDeclStmt(ctx.Stmt(0))
	if err != nil {
		return err
	}

	retFalse, err := v.evalNonDeclStmt(ctx.Stmt(1))
	if err != nil {
		return err
	}

	if b.constValue != nil {
		if *b.constValue {
			return retTrue
		} else {
			return retFalse
		}
	}

	return doesReturn{
		always:    retFalse.always || retTrue.always,
		sometimes: retFalse.sometimes || retTrue.sometimes,
	}
}

func (v *typeCheckVisitor) VisitSWhile(ctx *parser.SWhileContext) interface{} {
	t, err := v.evalType(ctx.Expr())
	if err != nil {
		return err
	}

	b, ok := t.(TBool)
	if !ok {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TBool{},
			Got:      t,
		}
	}

	returns, err := v.evalNonDeclStmt(ctx.Stmt())
	if err != nil {
		return err
	}

	if b.constValue != nil && *b.constValue {
		return doesReturn{
			always:    returns.sometimes,
			sometimes: returns.sometimes,
		}
	}

	return doesReturn{
		always:    false,
		sometimes: returns.sometimes,
	}
}

func (v *typeCheckVisitor) VisitSFor(ctx *parser.SForContext) interface{} {
	typ := v.Visit(ctx.Type_())
	if err, ok := typ.(error); ok {
		return err
	}

	t := typ.(Type)

	arrType, err := v.evalType(ctx.Expr())
	if err != nil {
		return err
	}

	arr, ok := arrType.(TArray)
	if !ok {
		return NotAnArrayError{
			Ctx:  ctx,
			Expr: ctx.Expr(),
			Type: arrType,
		}
	}

	if !sameType(t, arr.Elem) && !v.isSubClass(t, arr.Elem) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: arr.Elem,
			Got:      t,
		}
	}

	defer v.ShadowLocal(ctx.ID().GetText(), arr.Elem)()
	returns, err := v.evalNonDeclStmt(ctx.Stmt())
	if err != nil {
		return err
	}

	return doesReturn{
		always:    false,
		sometimes: returns.sometimes,
	}
}

func (v *typeCheckVisitor) VisitSExp(ctx *parser.SExpContext) interface{} {
	if _, err := v.evalType(ctx.Expr()); err != nil {
		return err
	}
	return doesReturn{
		always:    false,
		sometimes: false,
	}
}

func (v *typeCheckVisitor) evalLVType(ctx parser.ILvalueContext) (Type, error) {
	res := v.Visit(ctx)
	if err, ok := res.(error); ok {
		return nil, err
	}

	return res.(Type), nil
}

func (v *typeCheckVisitor) evalType(tree parser.IExprContext) (Type, error) {
	res := v.Visit(tree)
	if err, ok := res.(error); ok {
		return nil, err
	}

	if classRef, ok := res.(TClassRef); ok {
		class, _ := v.TypeOfGlobal(classRef.String())
		return class.Type, nil
	}

	if typ, ok := res.(Type); ok {
		return typ, nil
	}

	panic("evalType called in wrong context, visiting the tree should return a type")
}

func (v *typeCheckVisitor) VisitENotOp(ctx *parser.ENotOpContext) interface{} {
	err := v.ExpectType(TBool{}, ctx.Expr())
	if err != nil {
		return err
	}

	t, _ := v.evalType(ctx.Expr())
	cv, ok := t.ConstValue()
	if ok {
		cv := !cv.(bool)
		return TBool{
			StartToken: ctx.GetStart(),
			constValue: &cv,
		}
	}

	return TBool{
		StartToken: ctx.GetStart(),
	}
}

func (v *typeCheckVisitor) VisitENegOp(ctx *parser.ENegOpContext) interface{} {
	err := v.ExpectType(TInt{}, ctx.Expr())
	if err != nil {
		return err
	}

	t, _ := v.evalType(ctx.Expr())
	cv, ok := t.ConstValue()
	if ok {
		cv := -cv.(int)
		return TInt{
			StartToken: ctx.GetStart(),
			constValue: &cv,
		}
	}

	return TInt{
		StartToken: ctx.GetStart(),
	}
}

func (v *typeCheckVisitor) VisitEMulOp(ctx *parser.EMulOpContext) interface{} {
	t1, err := v.evalType(ctx.Expr(0))
	if err != nil {
		return err
	}

	t2, err := v.evalType(ctx.Expr(1))
	if err != nil {
		return err
	}

	if !sameType(t1, t2) {
		return ArgTypeMismatchError{
			Expr:  ctx,
			Type1: t1,
			Type2: t2,
		}
	}

	if _, ok := validMulOpArg[t1.String()]; !ok {
		return InvalidOpArgsError{
			Expr:       ctx,
			Type:       t1,
			ValidTypes: validMulOpArg,
		}
	}

	return TInt{
		StartToken: ctx.GetStart(),
		constValue: evalConstIntBinOp(ctx.MulOp().GetText(), t1, t2),
	}
}

func (v *typeCheckVisitor) VisitEAddOp(ctx *parser.EAddOpContext) interface{} {
	t1, err := v.evalType(ctx.Expr(0))
	if err != nil {
		return err
	}

	t2, err := v.evalType(ctx.Expr(1))
	if err != nil {
		return err
	}

	if !sameType(t1, t2) {
		return ArgTypeMismatchError{
			Expr:  ctx,
			Type1: t1,
			Type2: t2,
		}
	}

	validTypes := validAddOpArg
	if ctx.AddOp().GetText() == "-" {
		validTypes = validSubOpArg
	}

	if _, ok := validTypes[t1.String()]; !ok {
		return InvalidOpArgsError{
			Expr:       ctx,
			Type:       t1,
			ValidTypes: validTypes,
		}
	}

	switch t1.(type) {
	case TInt:
		return TInt{
			StartToken: ctx.GetStart(),
			constValue: evalConstIntBinOp(ctx.AddOp().GetText(), t1, t2),
		}
	case TString:
		return TString{
			StartToken: ctx.GetStart(),
			constValue: evalConstStringBinOp(ctx.AddOp().GetText(), t1, t2),
		}
	default:
		panic("unexpected addOp type")
	}
}

func (v *typeCheckVisitor) VisitERelOp(ctx *parser.ERelOpContext) interface{} {
	t1, err := v.evalType(ctx.Expr(0))
	if err != nil {
		return err
	}

	t2, err := v.evalType(ctx.Expr(1))
	if err != nil {
		return err
	}

	if !sameType(t1, t2) && !(v.isSubClass(t1, t2) || v.isSubClass(t2, t1)) {
		return ArgTypeMismatchError{
			Expr:  ctx,
			Type1: t1,
			Type2: t2,
		}
	}

	switch ctx.RelOp().GetText() {
	case "<", ">", "<=", ">=":
		if _, ok := validInequalityOpArg[t1.String()]; !ok {
			return InvalidOpArgsError{
				Expr:       ctx,
				Type:       t1,
				ValidTypes: validInequalityOpArg,
			}
		}
	}

	return TBool{
		StartToken: ctx.GetStart(),
		constValue: evalConstBoolBinOp(ctx.RelOp().GetText(), t1, t2),
	}
}

func (v *typeCheckVisitor) VisitEAnd(ctx *parser.EAndContext) interface{} {
	t1, err := v.evalType(ctx.Expr(0))
	if err != nil {
		return err
	}

	t2, err := v.evalType(ctx.Expr(1))
	if err != nil {
		return err
	}

	if !sameType(t1, TBool{}) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TBool{},
			Got:      t1,
		}
	}

	if !sameType(t2, TBool{}) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TBool{},
			Got:      t2,
		}
	}

	return TBool{
		StartToken: ctx.GetStart(),
		constValue: evalConstBoolBinOp("&&", t1, t2),
	}
}

func (v *typeCheckVisitor) VisitEOr(ctx *parser.EOrContext) interface{} {
	t1, err := v.evalType(ctx.Expr(0))
	if err != nil {
		return err
	}

	t2, err := v.evalType(ctx.Expr(1))
	if err != nil {
		return err
	}

	if !sameType(t1, TBool{}) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TBool{},
			Got:      t1,
		}
	}

	if !sameType(t2, TBool{}) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TBool{},
			Got:      t2,
		}
	}
	return TBool{
		StartToken: ctx.GetStart(),
		constValue: evalConstBoolBinOp("||", t1, t2),
	}
}

func (v *typeCheckVisitor) VisitENewArray(ctx *parser.ENewArrayContext) interface{} {
	t := v.Visit(ctx.Singular_type_())
	if err, ok := t.(error); ok {
		return err
	}

	typ := t.(Type)
	e := ctx.Expr()
	idxType, err := v.evalType(e)
	if err != nil {
		return err
	}

	if _, ok := idxType.(TInt); !ok {
		return ArraySizeTypeError{
			Expr: e,
			Type: idxType,
		}
	}

	typ = TArray{
		Elem: typ,
	}

	return typ
}

func (v *typeCheckVisitor) VisitENew(ctx *parser.ENewContext) interface{} {
	t := v.Visit(ctx.Singular_type_())
	if err, ok := t.(error); ok {
		return err
	}

	typ, ok := t.(TClass)
	if !ok {
		return UnknownClassError{
			Type: t.(Type),
		}
	}

	return typ
}

func (v *typeCheckVisitor) VisitEFieldAccess(ctx *parser.EFieldAccessContext) interface{} {
	t, err := v.evalType(ctx.Expr(0))
	if err != nil {
		return err
	}

	class, ok := t.(TClass)
	if ok {
		defer v.EnterClass(class)()
	} else if _, ok = t.(TArray); ok {
		defer v.ShadowLocal("length", TInt{})()
	}

	// Evaluate right-hand side in this environment.
	return v.Visit(ctx.Expr(1))
}

func (v *typeCheckVisitor) VisitEArrayRef(ctx *parser.EArrayRefContext) interface{} {
	typ, err := v.evalType(ctx.Expr(0))
	if err != nil {
		return err
	}

	arr, ok := typ.(TArray)
	if !ok {
		return ExpectedArrayError{
			Expr: ctx,
			Got:  typ,
		}
	}

	if err := v.ExpectType(TInt{}, ctx.Expr(1)); err != nil {
		return err
	}

	return arr.Elem
}

func (v *typeCheckVisitor) VisitESelf(ctx *parser.ESelfContext) interface{} {
	return v.curClass
}

func (v *typeCheckVisitor) VisitEId(ctx *parser.EIdContext) interface{} {
	ident := ctx.ID().GetText()
	if typ, ok := v.TypeOfLocal(ident); ok {
		return typ.Type
	}

	return UndeclaredIdentifierError{Ident: ctx.ID()}
}

func (v *typeCheckVisitor) VisitEInt(ctx *parser.EIntContext) interface{} {
	n, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		return ConstOutOfRangeError{
			Ctx: ctx,
		}
	}

	return TInt{
		StartToken: ctx.GetStart(),
		constValue: &n,
	}
}

func (v *typeCheckVisitor) VisitETrue(ctx *parser.ETrueContext) interface{} {
	b := true
	return TBool{
		StartToken: ctx.GetStart(),
		constValue: &b,
	}
}

func (v *typeCheckVisitor) VisitEFalse(ctx *parser.EFalseContext) interface{} {
	b := false
	return TBool{
		StartToken: ctx.GetStart(),
		constValue: &b,
	}
}

func (v *typeCheckVisitor) VisitEStr(ctx *parser.EStrContext) interface{} {
	withBraces := ctx.STR().GetText()
	s := withBraces[1 : len(withBraces)-1]
	return TString{
		StartToken: ctx.GetStart(),
		constValue: &s,
	}
}

func (v *typeCheckVisitor) VisitEFunCall(ctx *parser.EFunCallContext) interface{} {
	ident := ctx.ID().GetText()
	t, ok := v.TypeOf(ident)
	if !ok {
		return UndeclaredIdentifierError{
			Ident: ctx.ID(),
		}
	}

	signature, ok := t.Type.(TFun)
	if !ok {
		return NotAFunctionError{
			Ident: ctx.ID(),
			Type:  t.Type,
		}
	}

	if len(signature.Args) != len(ctx.AllExpr()) {
		return InvalidFunctionArgumentCountError{
			Expr: ctx,
			Fun:  signature,
		}
	}

	for i, e := range ctx.AllExpr() {
		if err := v.ExpectType(signature.Args[i].Type, e); err != nil {
			return err
		}
	}

	if classRef, ok := signature.Result.(TClassRef); ok {
		class, _ := v.TypeOfGlobal(classRef.String())
		return class.Type.(TClass)
	}
	return signature.Result
}

func (v *typeCheckVisitor) VisitENull(ctx *parser.ENullContext) interface{} {
	classRef := TClassRef{ctx.ID()}
	t, ok := v.TypeOfGlobal(classRef.String())
	if !ok {
		return UnknownClassError{classRef}
	}

	class, ok := t.Type.(TClass)
	if !ok {
		return UnknownClassError{
			Type: t,
		}
	}
	return class
}

func (v *typeCheckVisitor) VisitEParen(ctx *parser.EParenContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *typeCheckVisitor) VisitTNonVoid(ctx *parser.TNonVoidContext) interface{} {
	return v.Visit(ctx.Nvtype_())
}

func (v *typeCheckVisitor) VisitTSingular(ctx *parser.TSingularContext) interface{} {
	return v.Visit(ctx.Singular_type_())
}

func (v *typeCheckVisitor) VisitTInt(ctx *parser.TIntContext) interface{} {
	return TInt{StartToken: ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitTStr(ctx *parser.TStrContext) interface{} {
	return TString{StartToken: ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitTBool(ctx *parser.TBoolContext) interface{} {
	return TBool{StartToken: ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitTVoid(ctx *parser.TVoidContext) interface{} {
	return TVoid{StartToken: ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitTClass(ctx *parser.TClassContext) interface{} {
	ident := ctx.ID().GetText()
	t, ok := v.TypeOfGlobal(ident)
	if !ok {
		return UnknownClassError{
			Type: TClassRef{
				ID: ctx.ID(),
			},
		}
	}

	class, ok := t.Type.(TClass)
	if !ok {
		return UnknownClassError{
			Type: t,
		}
	}

	return class
}

func (v *typeCheckVisitor) VisitTArray(ctx *parser.TArrayContext) interface{} {
	t := v.Visit(ctx.Singular_type_())
	if err, ok := t.(error); ok {
		return err
	}

	return TArray{
		StartToken: ctx.GetStart(),
		Elem:       t.(Type),
	}
}
