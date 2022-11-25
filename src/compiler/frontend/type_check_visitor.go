package frontend

import (
	"fmt"
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type varDropper struct {
	drop  func()
	depth int
}

type typeCheckVisitor struct {
	parser.BaseLatteVisitor
	state        *state
	depth        int
	dropperStack []varDropper
	curClass     *TClass
}

func makeTypeCheckVisitor(s *state) *typeCheckVisitor {
	return &typeCheckVisitor{state: s}
}

func (v *typeCheckVisitor) Run() error {
	return v.Visit(v.state.tree).(error)
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

func (v *typeCheckVisitor) ExpectType(expected Type, ctx parser.IExprContext) error {
	if ctx == nil {
		return nil
	}

	got, err := v.evalType(ctx)
	if err != nil {
		return err
	}

	if !sameType(expected, got) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: expected,
			Got:      got,
		}
	}

	return nil
}

func (v *typeCheckVisitor) EnterClass(signature TClass) (exit func()) {
	v.depth++
	// Put all fields into enviroment.
	var exits []func()
	for ident, typ := range signature.Fields {
		exits = append(exits, v.ShadowLocal(ident, typ))
	}

	return func() {
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

	v.depth++
	for _, arg := range signature.Args {
		defer v.ShadowLocal(arg.Ident, arg.Type)()
	}
	v.depth--

	return v.Visit(ctx.Block())
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
	return nil
}

func (v *typeCheckVisitor) VisitClassMethodDef(ctx *parser.ClassMethodDefContext) interface{} {
	return v.Visit(ctx.Fundef())
}

func (v *typeCheckVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.depth++
	defer func() { v.depth-- }()
	for _, stmt := range ctx.AllStmt() {
		if err, ok := v.Visit(stmt).(error); ok {
			return err
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

	return nil
}

func (v *typeCheckVisitor) VisitSBlockStmt(ctx *parser.SBlockStmtContext) interface{} {
	return v.Visit(ctx.Block())
}

func (v *typeCheckVisitor) VisitSDecl(ctx *parser.SDeclContext) interface{} {
	t := v.Visit(ctx.Type_())
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
	return nil
}

func (v *typeCheckVisitor) VisitLVField(ctx *parser.LVFieldContext) interface{} {
	t, err := v.evalType(ctx.Expr())
	if err != nil {
		return err
	}

	if class, ok := t.(TClass); ok {
		defer v.EnterClass(class)()
	} else if _, ok = t.(TArray); ok {
		defer v.ShadowLocal("length", TInt{})()
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
	t, ok := v.TypeOf(ctx.ID().GetText())
	if !ok {
		return UndeclaredIdentifierError{
			Ident: ctx.ID(),
		}
	}

	return t.Type
}

func (v *typeCheckVisitor) VisitSAss(ctx *parser.SAssContext) interface{} {
	res := v.Visit(ctx.Lvalue())
	if err, ok := res.(error); ok {
		return err
	}

	t := res.(Type)
	return v.ExpectType(t, ctx.Expr())
}

func (v *typeCheckVisitor) VisitSExp(ctx *parser.SExpContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *typeCheckVisitor) evalType(tree parser.IExprContext) (Type, error) {
	res := v.Visit(tree)
	if err, ok := res.(error); ok {
		return nil, err
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
	return TBool{ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitENegOp(ctx *parser.ENegOpContext) interface{} {
	err := v.ExpectType(TInt{}, ctx.Expr())
	if err != nil {
		return err
	}
	return TInt{ctx.GetStart()}
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

	return t1
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

	return t1
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

	if !sameType(t1, t2) {
		return ArgTypeMismatchError{
			Expr:  ctx,
			Type1: t1,
			Type2: t2,
		}
	}

	if _, ok := validRelOpArg[t1.String()]; !ok {
		return InvalidOpArgsError{
			Expr:       ctx,
			Type:       t1,
			ValidTypes: validRelOpArg,
		}
	}

	return TBool{ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitEAnd(ctx *parser.EAndContext) interface{} {
	for _, e := range ctx.AllExpr() {
		if err := v.ExpectType(TBool{}, e); err != nil {
			return err
		}
	}

	return TBool{ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitEOr(ctx *parser.EOrContext) interface{} {
	for _, e := range ctx.AllExpr() {
		if err := v.ExpectType(TBool{}, e); err != nil {
			return err
		}
	}

	return TBool{ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitENew(ctx *parser.ENewContext) interface{} {
	t := v.Visit(ctx.Singular_type_())
	if err, ok := t.(error); ok {
		return err
	}

	typ := t.(Type)
	for _, e := range ctx.AllExpr() {
		idxType, err := v.evalType(e)
		if err != nil {
			return err
		}

		if _, ok := idxType.(TInt); !ok {
			return ArrayIndexTypeError{
				Expr: e,
				Type: idxType,
			}
		}

		typ = TArray{
			Elem: typ,
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
	if typ, ok := v.TypeOf(ident); ok {
		return typ.Type
	}

	return UndeclaredIdentifierError{Ident: ctx.ID()}
}

func (v *typeCheckVisitor) VisitEInt(ctx *parser.EIntContext) interface{} {
	return TInt{ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitETrue(ctx *parser.ETrueContext) interface{} {
	return TBool{ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitEFalse(ctx *parser.EFalseContext) interface{} {
	return TBool{ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitEStr(ctx *parser.EStrContext) interface{} {
	return TString{ctx.GetStart()}
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
	class, ok := v.TypeOfGlobal(classRef.String())
	if !ok {
		return UnknownClassError{classRef}
	}

	return class
}

func (v *typeCheckVisitor) VisitEParen(ctx *parser.EParenContext) interface{} {
	return v.Visit(ctx.Expr())
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
			Type: t,
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
	t := v.Visit(ctx.Type_())
	if err, ok := t.(error); ok {
		return err
	}

	return TArray{
		StartToken: ctx.GetStart(),
		Elem:       t.(Type),
	}
}
