package ir

import (
	"latte/compiler/frontend/types"
	"latte/parser"
	"strconv"
)

func (v *Visitor) evalExpr(tree parser.IExprContext) Location {
	return v.Visit(tree).(Location)
}

func (v *Visitor) GetBoolLoc() Location {
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

func (v *Visitor) VisitEMethodCall(ctx *parser.EMethodCallContext) interface{} {
	classPtr := v.evalExpr(ctx.Expr(0))
	class := classPtr.Type().(types.TClass)

	ident := ctx.ID().GetText()
	fieldInfo := class.Fields[ident]
	signature := fieldInfo.Type.(types.TFun)

	var classPtrHolder Location

	var args []Location
	var drop func()
	classPtrHolder, drop = v.ShadowLocal("_method_self", types.TInt{})
	defer drop()

	v.EmitQuad(QMov{Src: classPtr, Dst: classPtrHolder})
	v.EmitQuad(QPush{Src: classPtrHolder})

	args = append(args, classPtr)

	for i, e := range ctx.AllExpr()[1:] {
		pop := func() {}
		if _, ok := signature.Args[i].Type.(types.TBool); ok {
			lTrue, lFalse := v.FreshLabel("bool_argTrue"), v.FreshLabel("bool_argFalse")
			pop = v.PushLabels(lTrue, lFalse, lTrue)
		}
		arg := v.evalExpr(e)
		switch arg.(type) {
		case LUnassigned:
			arg = v.GetBoolLoc()
		case LDrop:
		default:
			v.EmitQuad(QPush{
				Src: arg,
			})
		}
		pop()
		args = append(args, arg)
	}

	dst := v.FreshTemp("call_tmp", signature.Result)

	vtablePtr := v.FreshTemp("vtab_ptr", types.TInt{})
	v.EmitQuad(QDeref{
		Src: classPtrHolder,
		Dst: vtablePtr,
	})
	funAddr := v.FreshTemp("vtab_addr", types.TInt{})
	v.EmitQuad(QArrayDeref{
		Array: vtablePtr,
		Index: LConst{Type_: types.TInt{}, Value: class.Fields[ident].Offset - 1},
		Dst:   funAddr,
	})

	v.EmitQuad(QCallMethod{
		Signature: signature,
		Label:     funAddr,
		Dst:       dst,
		Args:      args,
	})

	if _, ok := signature.Result.(types.TBool); ok {
		v.EmitQuad(QRelOp{
			Op:     "==",
			LFalse: v.lFalse,
			LTrue:  v.lTrue,
			LNext:  v.lNext,
			Lhs:    dst,
			Rhs:    LConst{Type_: types.TBool{}, Value: true},
		})
		return LUnassigned{Type_: types.TBool{}}
	}

	return dst

}

func (v *Visitor) VisitEFieldArrayAccess(ctx *parser.EFieldArrayAccessContext) interface{} {
	lhs := v.evalExpr(ctx.Expr(0))
	ident := ctx.ID().GetText()
	var array Location
	switch t := lhs.Type().(type) {
	case types.TClassRef:
		class := v.EvalClass(t.ID.GetText())
		fieldInfo := class.Fields[ident]
		ptr := v.FreshTemp("class_array_field_ref", fieldInfo.Type)
		array = v.FreshTemp("class_array_field", fieldInfo.Type)
		v.EmitQuad(QArrayAccess{
			Array: lhs,
			Index: LConst{Type_: types.TInt{}, Value: fieldInfo.Offset - 1},
			Dst:   ptr,
		})
		v.EmitQuad(QDeref{
			Src: ptr,
			Dst: array,
		})
	case types.TClass:
		fieldInfo := t.Fields[ident]
		ptr := v.FreshTemp("class_array_field_ref", fieldInfo.Type)
		array = v.FreshTemp("class_array_field", fieldInfo.Type)
		v.EmitQuad(QArrayAccess{
			Array: lhs,
			Index: LConst{Type_: types.TInt{}, Value: fieldInfo.Offset - 1},
			Dst:   ptr,
		})
		v.EmitQuad(QDeref{
			Src: ptr,
			Dst: array,
		})
	default:
		panic("field access happened on non array/class type")
	}

	index := v.evalExpr(ctx.Expr(1))

	ptr := v.FreshTemp("arr_access", array.Type().BaseType())

	v.EmitQuad(QArrayAccess{
		Array: array,
		Index: index,
		Dst:   ptr,
	})

	value := v.FreshTemp("arr_deref", array.Type().BaseType())
	v.EmitQuad(QDeref{
		Src: ptr,
		Dst: value,
	})

	if loc, ok := v.unassignBool(value); ok {
		return loc
	}

	return value
}

func (v *Visitor) EvalSelfFieldLocation(lhs Location, ident string) Location {
	var dst Location
	switch t := lhs.Type().(type) {
	case types.TClassRef:
		class := v.EvalClass(t.ID.GetText())
		fieldInfo := class.Fields[ident]
		ptr := v.FreshTemp("class_field_ref", fieldInfo.Type)
		dst = v.FreshTemp("class_field", fieldInfo.Type)
		v.EmitQuad(QArrayAccess{
			Array: lhs,
			Index: LConst{Type_: types.TInt{}, Value: fieldInfo.Offset - 1},
			Dst:   ptr,
		})
		v.EmitQuad(QDeref{
			Src: ptr,
			Dst: dst,
		})
	case types.TClass:
		fieldInfo := t.Fields[ident]
		ptr := v.FreshTemp("class_field_ref", fieldInfo.Type)
		dst = v.FreshTemp("class_field", fieldInfo.Type)
		v.EmitQuad(QArrayAccess{
			Array: lhs,
			Index: LConst{Type_: types.TInt{}, Value: fieldInfo.Offset - 1},
			Dst:   ptr,
		})
		v.EmitQuad(QDeref{
			Src: ptr,
			Dst: dst,
		})
	case types.TArray:
		// It must be a length call.
		dst = v.FreshTemp("array_length", types.TInt{})
		v.EmitQuad(QDeref{
			Src: lhs,
			Dst: dst,
		})
	default:
		panic("field access happened on non array/class type")
	}

	if loc, ok := v.unassignBool(dst); ok {
		return loc
	}

	return dst
}

func (v *Visitor) VisitEFieldAccess(ctx *parser.EFieldAccessContext) interface{} {
	lhs := v.evalExpr(ctx.Expr())
	ident := ctx.ID().GetText()
	return v.EvalSelfFieldLocation(lhs, ident)
}

func (v *Visitor) VisitEArrayRef(ctx *parser.EArrayRefContext) interface{} {
	array := v.evalExpr(ctx.Expr(0))
	index := v.evalExpr(ctx.Expr(1))
	ptr := v.FreshTemp("arr_access", array.Type().BaseType())

	v.EmitQuad(QArrayAccess{
		Array: array,
		Index: index,
		Dst:   ptr,
	})

	value := v.FreshTemp("arr_deref", array.Type().BaseType())
	v.EmitQuad(QDeref{
		Src: ptr,
		Dst: value,
	})

	if loc, ok := v.unassignBool(value); ok {
		return loc
	}

	return value
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

	return dst
}

func (v *Visitor) VisitERelOp(ctx *parser.ERelOpContext) interface{} {
	op := ctx.RelOp().GetText()

	lTrue, lFalse := v.FreshLabel("lRelTrue"), v.FreshLabel("lRelFalse")
	pop := v.PushLabels(lTrue, lFalse, lTrue)
	lhs := v.evalExpr(ctx.Expr(0))
	if _, ok := lhs.(LUnassigned); ok {
		lhs = v.GetBoolLoc()
	}
	pop()

	lTrue, lFalse = v.FreshLabel("lRelTrue"), v.FreshLabel("lRelFalse")
	pop = v.PushLabels(lTrue, lFalse, lTrue)
	rhs := v.evalExpr(ctx.Expr(1))
	if _, ok := rhs.(LUnassigned); ok {
		rhs = v.GetBoolLoc()
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
	t, _ := v.EvalType(ctx.Singular_type_())
	size := v.evalExpr(ctx.Expr())
	dst := v.FreshTemp("new_array", types.TArray{StartToken: ctx.GetStart(), Elem: t})
	v.EmitQuad(QNewArray{
		Type: t,
		Dst:  dst,
		Size: size,
	})
	return dst
}

func (v *Visitor) VisitENew(ctx *parser.ENewContext) interface{} {
	t, _ := v.EvalType(ctx.Singular_type_())
	class := t.(types.TClass)
	dst := v.FreshTemp("new_class", class)
	v.EmitQuad(QNewClass{
		Class: class,
		Dst:   dst,
	})
	return dst
}

func (v *Visitor) VisitESelf(ctx *parser.ESelfContext) interface{} {
	loc := v.GetLocal("self")
	if _, ok := loc.(LMem); ok {
		value := v.FreshTemp("eid_deref", loc.Type())
		v.EmitQuad(QMov{
			Src: loc,
			Dst: value,
		})
		return value
	}

	return loc
}

func (v *Visitor) unassignBool(loc Location) (Location, bool) {
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

		return LUnassigned{}, true
	}

	return nil, false
}

func (v *Visitor) VisitEId(ctx *parser.EIdContext) interface{} {
	ident := ctx.ID().GetText()
	loc := v.GetLocal(ident)

	if _, ok := loc.(LMem); ok {
		value := v.FreshTemp("eid_deref", loc.Type())
		v.EmitQuad(QMov{
			Src: loc,
			Dst: value,
		})
		return value
	}

	if _, ok := loc.(LSelfField); ok {
		lhs := v.VisitESelf(nil).(Location)
		return v.EvalSelfFieldLocation(lhs, ident)
	}

	if loc, ok := v.unassignBool(loc); ok {
		return loc
	}

	return loc
}

func (v *Visitor) VisitEFunCall(ctx *parser.EFunCallContext) interface{} {
	ident := ctx.ID().GetText()
	t, _ := v.TypeOf(ident)

	signature := t.Type.(types.TFun)

	var args []Location
	if signature.IsMethod {
		self := v.GetLocal("self")
		v.EmitQuad(QPush{Src: self})
		args = append(args, self)
	}

	for i, e := range ctx.AllExpr() {
		pop := func() {}
		if _, ok := signature.Args[i].Type.(types.TBool); ok {
			lTrue, lFalse := v.FreshLabel("bool_argTrue"), v.FreshLabel("bool_argFalse")
			pop = v.PushLabels(lTrue, lFalse, lTrue)
		}
		arg := v.evalExpr(e)
		switch arg.(type) {
		case LUnassigned:
			arg = v.GetBoolLoc()
		case LDrop:
		default:
			v.EmitQuad(QPush{
				Src: arg,
			})
		}
		pop()
		args = append(args, arg)
	}

	dst := v.FreshTemp("call_tmp", signature.Result)

	if signature.IsMethod {
		class := *v.CurClass

		vtablePtr := v.FreshTemp("vtab_ptr", types.TInt{})
		v.EmitQuad(QDeref{
			Src: v.GetLocal("self"),
			Dst: vtablePtr,
		})
		funAddr := v.FreshTemp("vtab_addr", types.TInt{})
		v.EmitQuad(QArrayDeref{
			Array: vtablePtr,
			Index: LConst{Type_: types.TInt{}, Value: class.Fields[ident].Offset - 1},
			Dst:   funAddr,
		})
		v.EmitQuad(QCallMethod{
			Signature: signature,
			Label:     funAddr,
			Dst:       dst,
			Args:      args,
		})
	} else {
		v.EmitQuad(QCall{
			Signature: signature,
			Label:     v.GetFunctionLabel(ident),
			Dst:       dst,
			Args:      args,
		})
	}

	if _, ok := signature.Result.(types.TBool); ok {
		v.EmitQuad(QRelOp{
			Op:     "==",
			LFalse: v.lFalse,
			LTrue:  v.lTrue,
			LNext:  v.lNext,
			Lhs:    dst,
			Rhs:    LConst{Type_: types.TBool{}, Value: true},
		})
		return LUnassigned{Type_: types.TBool{}}
	}

	return dst
}

func (v *Visitor) VisitEStr(ctx *parser.EStrContext) interface{} {
	withBraces := ctx.STR().GetText()
	s := withBraces[1 : len(withBraces)-1]
	s, _ = strconv.Unquote("\"" + s + "\"")
	return LConst{Type_: types.TString{}, Value: s}
}

func (v *Visitor) VisitENull(ctx *parser.ENullContext) interface{} {
	class := v.EvalClass(ctx.ID().GetText())
	return LConst{Type_: class, Value: 0}
}

func (v *Visitor) VisitEParen(ctx *parser.EParenContext) interface{} {
	return v.Visit(ctx.Expr())
}
