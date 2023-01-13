package ir

import (
	"latte/compiler/frontend/types"
	"latte/parser"
	"strconv"
)

func (v *Visitor) evalExpr(tree parser.IExprContext) Location {
	return v.Visit(tree).(Location)
}

func (v *Visitor) getBoolLoc() Location {
	t := types.TBool{}
	loc := v.FreshTemp("bool_loc", t)
	lEnd := v.FreshLabel("lBEnd")

	v.StartBlock(v.lTrue)
	v.EmitQuad(QMov{
		Dst: loc,
		Src: LConst{Type_: t, Value: true},
	})
	v.EmitQuad(QJmp{Dst: lEnd})

	v.StartBlock(v.lFalse)
	v.EmitQuad(QMov{
		Dst: loc,
		Src: LConst{Type_: t, Value: false},
	})

	v.StartBlock(lEnd)

	return loc
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
	dst := v.FreshTemp("neg_tmp", arg.Type())

	v.EmitQuad(QNeg{
		Dst: dst,
		Arg: arg,
	})

	return dst
}

func (v *Visitor) VisitENotOp(ctx *parser.ENotOpContext) interface{} {
	defer v.PushLabels(v.lFalse, v.lTrue, v.lNext)()
	v.Visit(ctx.Expr())
	return LUnassigned{Type_: types.TBool{}}
}

func (v *Visitor) VisitEMulOp(ctx *parser.EMulOpContext) interface{} {
	lhs := v.evalExpr(ctx.Expr(0))
	rhs := v.evalExpr(ctx.Expr(1))
	dst := v.FreshTemp("mul_tmp", lhs.Type())
	v.EmitQuad(QBinOp{
		Op:  ctx.MulOp().GetText(),
		Dst: dst,
		Lhs: lhs,
		Rhs: rhs,
	})

	return dst
}

func (v *Visitor) VisitEAddOp(ctx *parser.EAddOpContext) interface{} {
	lhs := v.evalExpr(ctx.Expr(0))
	rhs := v.evalExpr(ctx.Expr(1))
	op := ctx.AddOp().GetText()
	dst := v.FreshTemp("add_tmp", lhs.Type())
	v.EmitQuad(QBinOp{
		Op:  op,
		Dst: dst,
		Lhs: lhs,
		Rhs: rhs,
	})

	// log.Printf("%#v\n, %#v\n, %#v", dst, lhs, rhs)

	return dst
}

func (v *Visitor) VisitERelOp(ctx *parser.ERelOpContext) interface{} {
	op := ctx.RelOp().GetText()

	lTrue, lFalse := v.FreshLabel("lRelTrue"), v.FreshLabel("lRelFalse")
	pop := v.PushLabels(lTrue, lFalse, lTrue)
	lhs := v.evalExpr(ctx.Expr(0))
	if _, ok := lhs.(LUnassigned); ok {
		lhs = v.getBoolLoc()
	}
	pop()

	lTrue, lFalse = v.FreshLabel("lRelTrue"), v.FreshLabel("lRelFalse")
	pop = v.PushLabels(lTrue, lFalse, lTrue)
	rhs := v.evalExpr(ctx.Expr(1))
	if _, ok := rhs.(LUnassigned); ok {
		rhs = v.getBoolLoc()
	}
	pop()

	v.EmitQuad(QRelOp{
		Op:     op,
		LTrue:  v.lTrue,
		LFalse: v.lFalse,
		LNext:  v.lNext,
		Lhs:    lhs,
		Rhs:    rhs,
	})

	return LUnassigned{Type_: types.TBool{}}
}

func (v *Visitor) VisitEAnd(ctx *parser.EAndContext) interface{} {
	lp := v.FreshLabel("lANDp")

	pop := v.PushLabels(lp, v.lFalse, lp)
	v.Visit(ctx.Expr(0))
	pop()

	v.StartBlock(lp)
	v.Visit(ctx.Expr(1))

	return LUnassigned{Type_: types.TBool{}}
}

func (v *Visitor) VisitEInt(ctx *parser.EIntContext) interface{} {
	n, _ := strconv.Atoi(ctx.INT().GetText())

	return LConst{Type_: types.TInt{}, Value: n}
}

func (v *Visitor) VisitETrue(ctx *parser.ETrueContext) interface{} {
	if v.lTrue != v.lNext {
		v.EmitQuad(QJmp{Dst: v.lTrue})
	}

	return LUnassigned{Type_: types.TBool{}}
}

func (v *Visitor) VisitEFalse(ctx *parser.EFalseContext) interface{} {
	if v.lFalse != v.lNext {
		v.EmitQuad(QJmp{Dst: v.lFalse})
	}

	return LUnassigned{Type_: types.TBool{}}
}

func (v *Visitor) VisitEOr(ctx *parser.EOrContext) interface{} {
	lp := v.FreshLabel("lORp")

	pop := v.PushLabels(v.lTrue, lp, lp)
	v.Visit(ctx.Expr(0))
	pop()

	v.StartBlock(lp)
	v.Visit(ctx.Expr(1))

	return LUnassigned{Type_: types.TBool{}}
}

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
	loc := v.GetLocal(ident)
	if _, ok := loc.Type().(types.TBool); ok {
		if v.lTrue == v.lNext {
			v.EmitQuad(QJz{
				Value: loc,
				Dst:   v.lFalse,
			})
		} else if v.lFalse == v.lNext {
			v.EmitQuad(QJnz{
				Value: loc,
				Dst:   v.lTrue,
			})
		} else {
			v.EmitQuad(QJz{
				Value: loc,
				Dst:   v.lFalse,
			})
			v.EmitQuad(QJmp{
				Dst: v.lTrue,
			})
		}

		return LUnassigned{}
	}

	return loc
}

func (v *Visitor) VisitEFunCall(ctx *parser.EFunCallContext) interface{} {
	ident := ctx.ID().GetText()
	t, _ := v.TypeOf(ident)

	signature := t.Type.(types.TFun)

	var args []Location
	for _, e := range ctx.AllExpr() {
		arg := v.evalExpr(e)
		if _, ok := arg.(LUnassigned); ok {
			arg = v.getBoolLoc()
		}
		if _, ok := arg.(LConst); ok {
			v.EmitQuad(QPush{
				Src: arg,
			})
		}
		if _, ok := arg.(LReg); ok {
			v.EmitQuad(QPush{
				Src: arg,
			})
		}

		args = append(args, arg)
	}

	dst := v.FreshTemp("call_tmp", signature.Result)
	// Emit call.
	v.EmitQuad(QCall{
		Signature: signature,
		Label:     v.GetFunctionLabel(ident),
		Dst:       dst,
		Args:      args,
	})

	return dst
}

func (v *Visitor) VisitEStr(ctx *parser.EStrContext) interface{} {
	withBraces := ctx.STR().GetText()
	s := withBraces[1 : len(withBraces)-1]
	return LConst{Type_: types.TString{}, Value: s}
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
