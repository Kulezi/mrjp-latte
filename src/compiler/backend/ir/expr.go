package ir

import (
	"latte/parser"
)

func (v *Visitor) evalExpr(tree parser.IExprContext) Location {
	ret := v.Visit(tree).(Location)
	return ret
}

func (v *Visitor) VisitEFieldAccess(ctx *parser.EFieldAccessContext) interface{} {
	panic("can't access field - classes are not yet supported")
	// t, err := v.evalExpr(ctx.Expr(0))
	// if err != nil {
	// 	return err
	// }

	// class, ok := t.(TClass)
	// if ok {
	// 	defer v.EnterClass(class)()
	// } else if _, ok = t.(TArray); ok {
	// 	defer v.ShadowLocal("length", TInt{})()
	// }

	// // Evaluate right-hand side in this environment.
	// return v.Visit(ctx.Expr(1))
}

func (v *Visitor) VisitEArrayRef(ctx *parser.EArrayRefContext) interface{} {
	panic("arrays are not yet supported")
	// t, err := v.evalExpr(ctx.Expr(0))
	// if err != nil {
	// 	return err
	// }

	// arr, ok := t.(TArray)
	// if !ok {
	// 	return ExpectedArrayError{
	// 		Expr: ctx,
	// 		Got:  t,
	// 	}
	// }

	// if err := v.ExpectType(TInt{}, ctx.Expr(1)); err != nil {
	// 	return err
	// }

	// return arr.Elem
}

func (v *Visitor) VisitENegOp(ctx *parser.ENegOpContext) interface{} {
	arg := v.evalExpr(ctx.Expr())
	dst := v.FreshTemp(arg.Type())

	v.EmitQuad(QUnOp{
		Op:  "-",
		Dst: dst,
		Arg: arg,
	})

	return dst
}

// func (v *Visitor) VisitENotOp(ctx *parser.ENotOpContext) interface{} {
// 	t, err := v.evalExpr(ctx.Expr())
// 	if err != nil {
// 		return err
// 	}

// 	b, ok := t.(TBool)
// 	if !ok {
// 		return UnexpectedTypeError{
// 			Expr:     ctx,
// 			Expected: TBool{},
// 			Got:      t,
// 		}
// 	}

// 	return TBool{
// 		StartToken: ctx.GetStart(),
// 		Value:      EvalConstBoolNotOp(b),
// 	}
// }

func (v *Visitor) VisitEMulOp(ctx *parser.EMulOpContext) interface{} {
	panic("TODO")
	// t1, err := v.evalExpr(ctx.Expr(0))
	// t2, err := v.evalExpr(ctx.Expr(1))
	// if err != nil {
	// 	return err
	// }

	// if !SameType(t1, t2) {
	// 	return ArgTypeMismatchError{
	// 		Expr:  ctx,
	// 		Type1: t1,
	// 		Type2: t2,
	// 	}
	// }

	// if _, ok := validMulOpArg[t1.String()]; !ok {
	// 	return InvalidOpArgsError{
	// 		Expr:       ctx,
	// 		Type:       t1,
	// 		ValidTypes: validMulOpArg,
	// 	}
	// }

	// cv, ok := EvalConstIntBinOp(ctx.MulOp().GetText(), t1, t2)
	// if !ok {
	// 	return ZeroDivisionError{
	// 		Ctx: ctx,
	// 	}
	// }

	// return TInt{
	// 	StartToken: ctx.GetStart(),
	// 	Value:      cv,
	// }
}

// func (v *Visitor) VisitEAddOp(ctx *parser.EAddOpContext) interface{} {
// 	t1, err := v.evalExpr(ctx.Expr(0))
// 	if err != nil {
// 		return err
// 	}

// 	t2, err := v.evalExpr(ctx.Expr(1))
// 	if err != nil {
// 		return err
// 	}

// 	if !SameType(t1, t2) {
// 		return ArgTypeMismatchError{
// 			Expr:  ctx,
// 			Type1: t1,
// 			Type2: t2,
// 		}
// 	}

// 	validTypes := validAddOpArg
// 	if ctx.AddOp().GetText() == "-" {
// 		validTypes = validSubOpArg
// 	}

// 	if _, ok := validTypes[t1.String()]; !ok {
// 		return InvalidOpArgsError{
// 			Expr:       ctx,
// 			Type:       t1,
// 			ValidTypes: validTypes,
// 		}
// 	}

// 	switch t1.(type) {
// 	case TInt:
// 		cv, _ := EvalConstIntBinOp(ctx.AddOp().GetText(), t1, t2)
// 		return TInt{
// 			StartToken: ctx.GetStart(),
// 			Value:      cv,
// 		}
// 	case TString:
// 		return TString{
// 			StartToken: ctx.GetStart(),
// 			Value:      EvalConstStringBinOp(ctx.AddOp().GetText(), t1, t2),
// 		}
// 	default:
// 		panic("unexpected addOp type")
// 	}
// }

// func (v *Visitor) VisitERelOp(ctx *parser.ERelOpContext) interface{} {
// 	t1, err := v.evalExpr(ctx.Expr(0))
// 	if err != nil {
// 		return err
// 	}

// 	t2, err := v.evalExpr(ctx.Expr(1))
// 	if err != nil {
// 		return err
// 	}

// 	if !SameType(t1, t2) && !(v.isSubClass(t1, t2) || v.isSubClass(t2, t1)) {
// 		return ArgTypeMismatchError{
// 			Expr:  ctx,
// 			Type1: t1,
// 			Type2: t2,
// 		}
// 	}

// 	switch ctx.RelOp().GetText() {
// 	case "<", ">", "<=", ">=":
// 		if _, ok := validInequalityOpArg[t1.String()]; !ok {
// 			return InvalidOpArgsError{
// 				Expr:       ctx,
// 				Type:       t1,
// 				ValidTypes: validInequalityOpArg,
// 			}
// 		}
// 	}

// 	return TBool{
// 		StartToken: ctx.GetStart(),
// 		Value:      EvalConstBoolBinOp(ctx.RelOp().GetText(), t1, t2),
// 	}
// }

// func (v *Visitor) VisitEAnd(ctx *parser.EAndContext) interface{} {
// 	t1, err := v.evalExpr(ctx.Expr(0))
// 	if err != nil {
// 		return err
// 	}

// 	t2, err := v.evalExpr(ctx.Expr(1))
// 	if err != nil {
// 		return err
// 	}

// 	if !SameType(t1, TBool{}) {
// 		return UnexpectedTypeError{
// 			Expr:     ctx,
// 			Expected: TBool{},
// 			Got:      t1,
// 		}
// 	}

// 	if !SameType(t2, TBool{}) {
// 		return UnexpectedTypeError{
// 			Expr:     ctx,
// 			Expected: TBool{},
// 			Got:      t2,
// 		}
// 	}

// 	return TBool{
// 		StartToken: ctx.GetStart(),
// 		Value:      EvalConstBoolBinOp("&&", t1, t2),
// 	}
// }

// func (v *Visitor) VisitEOr(ctx *parser.EOrContext) interface{} {
// 	t1, err := v.evalExpr(ctx.Expr(0))
// 	if err != nil {
// 		return err
// 	}

// 	t2, err := v.evalExpr(ctx.Expr(1))
// 	if err != nil {
// 		return err
// 	}

// 	if !SameType(t1, TBool{}) {
// 		return UnexpectedTypeError{
// 			Expr:     ctx,
// 			Expected: TBool{},
// 			Got:      t1,
// 		}
// 	}

// 	if !SameType(t2, TBool{}) {
// 		return UnexpectedTypeError{
// 			Expr:     ctx,
// 			Expected: TBool{},
// 			Got:      t2,
// 		}
// 	}
// 	return TBool{
// 		StartToken: ctx.GetStart(),
// 		Value:      EvalConstBoolBinOp("||", t1, t2),
// 	}
// }

func (v *Visitor) VisitENewArray(ctx *parser.ENewArrayContext) interface{} {
	panic("can't use new - arrays are not yet supported")
	// t, err := v.evalType(ctx.Singular_type_())
	// if err != nil {
	// 	return err
	// }

	// e := ctx.Expr()
	// idxType, err := v.evalExpr(e)
	// if err != nil {
	// 	return err
	// }

	// if _, ok := idxType.(TInt); !ok {
	// 	return ArraySizeTypeError{
	// 		Expr: e,
	// 		Type: idxType,
	// 	}
	// }

	// t = TArray{
	// 	Elem: t,
	// }

	// return t
}

func (v *Visitor) VisitENew(ctx *parser.ENewContext) interface{} {
	panic("can't use new - classes are not yet supported")
	// t, err := v.evalType(ctx.Singular_type_())
	// if err != nil {
	// 	return err
	// }

	// class, ok := t.(TClass)
	// if !ok {
	// 	return UnknownClassError{
	// 		Type: t,
	// 	}
	// }

	// return class
}

func (v *Visitor) VisitESelf(ctx *parser.ESelfContext) interface{} {
	panic("can't use self - classes are not yet supported")
	// return v.CurClass
}

func (v *Visitor) VisitEId(ctx *parser.EIdContext) interface{} {
	ident := ctx.ID().GetText()
	return v.GetLocal(ident)
}

// func (v *Visitor) VisitEFunCall(ctx *parser.EFunCallContext) interface{} {
// 	ident := ctx.ID().GetText()
// 	t, ok := v.TypeOf(ident)
// 	if !ok {
// 		return UndeclaredIdentifierError{
// 			Ident: ctx.ID(),
// 		}
// 	}

// 	signature, ok := t.Type.(TFun)
// 	if !ok {
// 		return NotAFunctionError{
// 			Ident: ctx.ID(),
// 			Type:  t.Type,
// 		}
// 	}

// 	if len(signature.Args) != len(ctx.AllExpr()) {
// 		return InvalidFunctionArgumentCountError{
// 			Expr: ctx,
// 			Fun:  signature,
// 		}
// 	}

// 	for i, e := range ctx.AllExpr() {
// 		if err := v.ExpectType(signature.Args[i].Type, e); err != nil {
// 			return err
// 		}
// 	}

// 	if classRef, ok := signature.Result.(TClassRef); ok {
// 		class, _ := v.TypeOfGlobal(classRef.String())
// 		return class.Type.(TClass)
// 	}
// 	return signature.Result
// }

func (v *Visitor) VisitEStr(ctx *parser.EStrContext) interface{} {
	withBraces := ctx.STR().GetText()
	s := withBraces[1 : len(withBraces)-1]
	return LConst{Value: s}
}

func (v *Visitor) VisitENull(ctx *parser.ENullContext) interface{} {
	panic("nulls are not yet supported")
	// classRef := TClassRef{ctx.ID()}
	// t, ok := v.TypeOfGlobal(classRef.String())
	// if !ok {
	// 	return UnknownClassError{classRef}
	// }

	// class, ok := t.Type.(TClass)
	// if !ok {
	// 	return UnknownClassError{
	// 		Type: t,
	// 	}
	// }
	// return class
}

func (v *Visitor) VisitEParen(ctx *parser.EParenContext) interface{} {
	return v.Visit(ctx.Expr())
}

// var validInequalityOpArg = map[string]struct{}{
// 	"int":    {},
// 	"string": {},
// }

// var validAddOpArg = map[string]struct{}{
// 	"int":    {},
// 	"string": {},
// }
// var validSubOpArg = map[string]struct{}{
// 	"int": {},
// }

// var validMulOpArg = map[string]struct{}{
// 	"int": {},
// }
