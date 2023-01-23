package ir

import (
	"latte/compiler/frontend/typecheck"
	"latte/compiler/frontend/types"
	"latte/parser"
)

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.Depth++
	defer func() { v.Depth-- }()
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	for len(v.DropperStack) > 0 {
		dropper := v.DropperStack[len(v.DropperStack)-1]
		if dropper.Depth != v.Depth {
			break
		}

		dropper.Drop()
		v.DropperStack = v.DropperStack[:len(v.DropperStack)-1]
	}

	return nil
}

func (v *Visitor) VisitSEmpty(ctx *parser.SEmptyContext) interface{} {
	return nil
}

func (v *Visitor) VisitSBlockStmt(ctx *parser.SBlockStmtContext) interface{} {
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitSDecl(ctx *parser.SDeclContext) interface{} {
	t, _ := v.Visitor.EvalType(ctx.Nvtype_())

	for _, item := range ctx.AllItem() {
		item, ok := item.(*parser.ItemContext)
		if !ok {
			panic("unexpected antlr behaviour: item is not an ItemContext")
		}
		ident := item.ID().GetText()

		var src Location
		if item.Expr() != nil {
			if _, ok := t.(types.TBool); ok {
				lTrue, lFalse := v.FreshLabel("retTrue"), v.FreshLabel("retFalse")
				defer v.PushLabels(lTrue, lFalse, lTrue)()
			}
			src = v.evalSExp(item.Expr())
			if _, ok := src.(LUnassigned); ok {
				src = v.GetBoolLoc()
			}
		} else {
			src = LConst{
				Type_: t,
				Value: t.DefaultValue(),
			}
		}

		dst, drop := v.ShadowLocal(ident, t)
		v.DropperStack = append(v.DropperStack, typecheck.VarDropper{
			Drop:  drop,
			Depth: v.Depth,
		})
		v.EmitQuad(QMov{
			Src: src,
			Dst: dst,
		})
	}

	return nil
}

func (v *Visitor) VisitSAss(ctx *parser.SAssContext) interface{} {
	dst := v.evalLV(ctx.Lvalue())
	if _, ok := dst.Type().(types.TBool); ok {
		lTrue, lFalse := v.FreshLabel("sassTrue"), v.FreshLabel("sassFalse")
		defer v.PushLabels(lTrue, lFalse, lTrue)()
	}
	src := v.evalSExp(ctx.Expr())
	if _, ok := src.(LUnassigned); ok {
		src = v.GetBoolLoc()
	}

	v.EmitQuad(QMov{
		Src: src,
		Dst: dst,
	})

	return nil
}

func (v *Visitor) VisitSIncr(ctx *parser.SIncrContext) interface{} {
	lhs := v.evalLV(ctx.Lvalue())
	dst := v.evalLV(ctx.Lvalue())
	v.EmitQuad(QBinOp{
		Dst: dst,
		Op:  "+",
		Lhs: lhs,
		Rhs: LConst{
			Type_: types.TInt{},
			Value: 1,
		},
	})

	return nil
}

func (v *Visitor) VisitSDecr(ctx *parser.SDecrContext) interface{} {
	lhs := v.evalLV(ctx.Lvalue())
	dst := v.evalLV(ctx.Lvalue())
	v.EmitQuad(QBinOp{
		Dst: dst,
		Op:  "-",
		Lhs: lhs,
		Rhs: LConst{
			Type_: types.TInt{},
			Value: 1,
		},
	})

	return nil
}

func (v *Visitor) VisitSRet(ctx *parser.SRetContext) interface{} {
	if loc, ok := v.evalConstExpr(ctx.Expr()); ok {
		v.EmitQuad(QRet{
			Value: loc,
		})
	} else {
		if _, ok := v.CurFun.Result.(types.TBool); ok {
			lTrue, lFalse := v.FreshLabel("retTrue"), v.FreshLabel("retFalse")
			defer v.PushLabels(lTrue, lFalse, lTrue)()
		}
		ret := v.evalExpr(ctx.Expr())
		if _, ok := ret.(LUnassigned); ok {
			ret = v.GetBoolLoc()
		}
		v.EmitQuad(QRet{
			Value: ret,
		})
	}
	return nil
}

func (v *Visitor) VisitSVRet(ctx *parser.SVRetContext) interface{} {
	v.EmitQuad(QVRet{})
	return nil
}

func (v *Visitor) VisitSCond(ctx *parser.SCondContext) interface{} {
	if b, ok := v.evalConstExpr(ctx.Expr()); ok {
		b := b.(LConst).Value.(bool)
		if b {
			v.Visit(ctx.Stmt())
		}

		return nil
	}

	lTrue := v.FreshLabel("lCTrue")
	lFalse := v.FreshLabel("lCFalse")

	// Generate condition code.
	popLabels := v.PushLabels(lTrue, lFalse, lTrue)
	v.evalExpr(ctx.Expr())
	popLabels()

	// Generate if body.
	v.StartBlock(lTrue)
	v.Visit(ctx.Stmt())

	v.StartBlock(lFalse)
	return nil
}

func (v *Visitor) VisitSCondElse(ctx *parser.SCondElseContext) interface{} {
	if b, ok := v.evalConstExpr(ctx.Expr()); ok {
		b := b.(LConst).Value.(bool)
		if b {
			v.Visit(ctx.Stmt(0))
		} else {
			v.Visit(ctx.Stmt(1))
		}

		return nil
	}

	lTrue := v.FreshLabel("lCETrue")
	lFalse := v.FreshLabel("lCEFalse")
	lEnd := v.FreshLabel("lCEEnd")

	// Generate condition code.
	popLabels := v.PushLabels(lTrue, lFalse, lTrue)
	v.Visit(ctx.Expr())
	popLabels()

	// Generate if body.
	v.StartBlock(lTrue)
	v.Visit(ctx.Stmt(0))
	v.EmitQuad(QJmp{
		Dst: lEnd,
	})

	// Generate else body.
	v.StartBlock(lFalse)
	v.Visit(ctx.Stmt(1))

	v.StartBlock(lEnd)
	return nil
}

func (v *Visitor) VisitSWhile(ctx *parser.SWhileContext) interface{} {
	genBody := true
	genCond := true
	if b, ok := v.evalConstExpr(ctx.Expr()); ok {
		b := b.(LConst).Value.(bool)
		if !b {
			genBody = false
		} else {
			genCond = false
		}
	}

	lBody := v.FreshLabel("lWBody")
	lCond := v.FreshLabel("lWCond")
	lEnd := v.FreshLabel("lWEnd")

	v.EmitQuad(QJmp{
		Dst: lCond,
	})

	v.StartBlock(lBody)
	if genBody {
		v.Visit(ctx.Stmt())
	}

	v.StartBlock(lCond)
	if genCond {
		popLabels := v.PushLabels(lBody, lEnd, lEnd)
		v.evalExpr(ctx.Expr())
		popLabels()
	} else {
		v.EmitQuad(QJmp{
			Dst: lBody,
		})
	}

	v.StartBlock(lEnd)

	return nil
}

func isVariable(loc Location) bool {
	if reg, ok := loc.(LReg); ok {
		return reg.Variable != ""
	}
	return false
}

func (v *Visitor) VisitSFor(ctx *parser.SForContext) interface{} {
	lBody := v.FreshLabel("_for_body")
	lCond := v.FreshLabel("_for_cond")
	lEnd := v.FreshLabel("_for_end")

	arr := v.evalSExp(ctx.Expr())

	arrLoc, drop := v.ShadowLocal("arr_ptr", types.TInt{})
	defer drop()
	v.EmitQuad(QMov{
		Src: arr,
		Dst: arrLoc,
	})

	loc, drop := v.ShadowLocal(ctx.ID().GetText(), arr.Type().BaseType())
	defer drop()
	counter, drop := v.ShadowLocal("for_counter", types.TInt{})
	defer drop()
	v.EmitQuad(QMov{
		Src: LConst{Type_: types.TInt{}, Value: 0},
		Dst: counter,
	})
	v.EmitQuad(QJmp{Dst: lCond})

	v.StartBlock(lBody)
	ptr := v.FreshTemp("for_ptr", types.TInt{})
	v.EmitQuad(QArrayAccess{
		Array: arrLoc,
		Index: counter,
		Dst:   ptr,
	})
	v.EmitQuad(QDeref{
		Src: ptr,
		Dst: loc,
	})
	v.Visit(ctx.Stmt())
	v.EmitQuad(QBinOp{
		Op:  "+",
		Lhs: counter,
		Rhs: LConst{Type_: types.TInt{}, Value: 1},
		Dst: counter,
	})

	v.StartBlock(lCond)
	defer v.PushLabels(lBody, lEnd, lEnd)()
	v.EmitQuad(QRelOp{
		Op:     "==",
		Lhs:    counter,
		Rhs:    LMem{Type_: types.TInt{}, Addr: arrLoc},
		LTrue:  lEnd,
		LNext:  lEnd,
		LFalse: lBody,
	})

	v.StartBlock(lEnd)
	return nil
}

func (v *Visitor) VisitSExp(ctx *parser.SExpContext) interface{} {
	expr := ctx.Expr()
	t := v.Visitor.Visit(expr).(types.Type)
	// If expression is constant we don't need to generate any code as it doesn't do anything to the state.
	if _, ok := t.Const(); ok {
		return nil
	}

	// We don't care about the result but we need to give jumping code some label it can jump to.
	if _, ok := t.(types.TBool); ok {
		label := v.FreshLabel("_sexp_end")
		defer v.PushLabels(label, label, label)()
		v.evalSExp(ctx.Expr())
		v.StartBlock(label)
		return nil
	}

	loc := v.evalSExp(ctx.Expr())
	// We need to pop the result from stack as we won't use it.
	v.EmitQuad(QMov{
		Src: loc,
		Dst: LDrop{Type_: t},
	})
	return nil
}

func (v *Visitor) evalSExp(expr parser.IExprContext) Location {
	if loc, ok := v.evalConstExpr(expr); ok {
		return loc
	}

	return v.evalExpr(expr)
}

func (v *Visitor) evalConstExpr(expr parser.IExprContext) (value Location, ok bool) {
	t := v.Visitor.Visit(expr).(types.Type)
	if v, ok := t.Const(); ok {
		return LConst{
			Type_: t,
			Value: v,
		}, true
	}

	return nil, false
}
