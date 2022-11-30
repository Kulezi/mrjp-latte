package typecheck

import (
	. "latte/compiler/frontend/types"
	"latte/parser"
	"strconv"
)

func (v *visitor) evalExpr(tree parser.IExprContext) (Type, error) {
	res := v.Visit(tree)
	if err, ok := res.(error); ok {
		return nil, err
	}

	if classRef, ok := res.(TClassRef); ok {
		class, _ := v.TypeOfGlobal(classRef.String())
		return class.Type, nil
	}

	if t, ok := res.(Type); ok {
		return t, nil
	}

	panic("evalExpr called in wrong context, visiting the tree should return a type")
}

func (v *visitor) VisitEFieldAccess(ctx *parser.EFieldAccessContext) interface{} {
	t, err := v.evalExpr(ctx.Expr(0))
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

func (v *visitor) VisitEArrayRef(ctx *parser.EArrayRefContext) interface{} {
	t, err := v.evalExpr(ctx.Expr(0))
	if err != nil {
		return err
	}

	arr, ok := t.(TArray)
	if !ok {
		return ExpectedArrayError{
			Expr: ctx,
			Got:  t,
		}
	}

	if err := v.ExpectType(TInt{}, ctx.Expr(1)); err != nil {
		return err
	}

	return arr.Elem
}

func (v *visitor) VisitENegOp(ctx *parser.ENegOpContext) interface{} {
	t, err := v.evalExpr(ctx.Expr())
	if err != nil {
		return err
	}

	n, ok := t.(TInt)
	if !ok {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TInt{},
			Got:      t,
		}
	}

	return TInt{
		StartToken: ctx.GetStart(),
		Value:      EvalConstIntNegOp(n),
	}
}

func (v *visitor) VisitENotOp(ctx *parser.ENotOpContext) interface{} {
	t, err := v.evalExpr(ctx.Expr())
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

	return TBool{
		StartToken: ctx.GetStart(),
		Value:      EvalConstBoolNotOp(b),
	}
}

func (v *visitor) VisitEMulOp(ctx *parser.EMulOpContext) interface{} {
	t1, err := v.evalExpr(ctx.Expr(0))
	if err != nil {
		return err
	}

	t2, err := v.evalExpr(ctx.Expr(1))
	if err != nil {
		return err
	}

	if !SameType(t1, t2) {
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
		Value:      EvalConstIntBinOp(ctx.MulOp().GetText(), t1, t2),
	}
}

func (v *visitor) VisitEAddOp(ctx *parser.EAddOpContext) interface{} {
	t1, err := v.evalExpr(ctx.Expr(0))
	if err != nil {
		return err
	}

	t2, err := v.evalExpr(ctx.Expr(1))
	if err != nil {
		return err
	}

	if !SameType(t1, t2) {
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
			Value:      EvalConstIntBinOp(ctx.AddOp().GetText(), t1, t2),
		}
	case TString:
		return TString{
			StartToken: ctx.GetStart(),
			Value:      EvalConstStringBinOp(ctx.AddOp().GetText(), t1, t2),
		}
	default:
		panic("unexpected addOp type")
	}
}

func (v *visitor) VisitERelOp(ctx *parser.ERelOpContext) interface{} {
	t1, err := v.evalExpr(ctx.Expr(0))
	if err != nil {
		return err
	}

	t2, err := v.evalExpr(ctx.Expr(1))
	if err != nil {
		return err
	}

	if !SameType(t1, t2) && !(v.isSubClass(t1, t2) || v.isSubClass(t2, t1)) {
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
		Value:      EvalConstBoolBinOp(ctx.RelOp().GetText(), t1, t2),
	}
}

func (v *visitor) VisitEAnd(ctx *parser.EAndContext) interface{} {
	t1, err := v.evalExpr(ctx.Expr(0))
	if err != nil {
		return err
	}

	t2, err := v.evalExpr(ctx.Expr(1))
	if err != nil {
		return err
	}

	if !SameType(t1, TBool{}) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TBool{},
			Got:      t1,
		}
	}

	if !SameType(t2, TBool{}) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TBool{},
			Got:      t2,
		}
	}

	return TBool{
		StartToken: ctx.GetStart(),
		Value:      EvalConstBoolBinOp("&&", t1, t2),
	}
}

func (v *visitor) VisitEOr(ctx *parser.EOrContext) interface{} {
	t1, err := v.evalExpr(ctx.Expr(0))
	if err != nil {
		return err
	}

	t2, err := v.evalExpr(ctx.Expr(1))
	if err != nil {
		return err
	}

	if !SameType(t1, TBool{}) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TBool{},
			Got:      t1,
		}
	}

	if !SameType(t2, TBool{}) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: TBool{},
			Got:      t2,
		}
	}
	return TBool{
		StartToken: ctx.GetStart(),
		Value:      EvalConstBoolBinOp("||", t1, t2),
	}
}

func (v *visitor) VisitENewArray(ctx *parser.ENewArrayContext) interface{} {
	t, err := v.evalType(ctx.Singular_type_())
	if err != nil {
		return err
	}

	e := ctx.Expr()
	idxType, err := v.evalExpr(e)
	if err != nil {
		return err
	}

	if _, ok := idxType.(TInt); !ok {
		return ArraySizeTypeError{
			Expr: e,
			Type: idxType,
		}
	}

	t = TArray{
		Elem: t,
	}

	return t
}

func (v *visitor) VisitENew(ctx *parser.ENewContext) interface{} {
	t, err := v.evalType(ctx.Singular_type_())
	if err != nil {
		return err
	}

	class, ok := t.(TClass)
	if !ok {
		return UnknownClassError{
			Type: t,
		}
	}

	return class
}

func (v *visitor) VisitESelf(ctx *parser.ESelfContext) interface{} {
	return v.curClass
}

func (v *visitor) VisitEId(ctx *parser.EIdContext) interface{} {
	ident := ctx.ID().GetText()
	if t, ok := v.TypeOfLocal(ident); ok {
		return t.Type
	}

	return UndeclaredIdentifierError{Ident: ctx.ID()}
}

func (v *visitor) VisitEInt(ctx *parser.EIntContext) interface{} {
	n, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		return ConstOutOfRangeError{
			Ctx: ctx,
		}
	}

	return TInt{
		StartToken: ctx.GetStart(),
		Value:      &n,
	}
}

func (v *visitor) VisitETrue(ctx *parser.ETrueContext) interface{} {
	b := true
	return TBool{
		StartToken: ctx.GetStart(),
		Value:      &b,
	}
}

func (v *visitor) VisitEFalse(ctx *parser.EFalseContext) interface{} {
	b := false
	return TBool{
		StartToken: ctx.GetStart(),
		Value:      &b,
	}
}

func (v *visitor) VisitEFunCall(ctx *parser.EFunCallContext) interface{} {
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

func (v *visitor) VisitEStr(ctx *parser.EStrContext) interface{} {
	withBraces := ctx.STR().GetText()
	s := withBraces[1 : len(withBraces)-1]
	return TString{
		StartToken: ctx.GetStart(),
		Value:      &s,
	}
}

func (v *visitor) VisitENull(ctx *parser.ENullContext) interface{} {
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

func (v *visitor) VisitEParen(ctx *parser.EParenContext) interface{} {
	return v.Visit(ctx.Expr())
}

var validInequalityOpArg = map[string]struct{}{
	"int":    {},
	"string": {},
}

var validAddOpArg = map[string]struct{}{
	"int":    {},
	"string": {},
}
var validSubOpArg = map[string]struct{}{
	"int": {},
}

var validMulOpArg = map[string]struct{}{
	"int": {},
}
