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
			src = v.evalSExp(item.Expr())
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
	src := v.evalSExp(ctx.Expr())
	_, dst := v.evalLV(ctx.Lvalue())
	v.EmitQuad(QMov{
		Src: src,
		Dst: dst,
	})

	return nil
}

func (v *Visitor) VisitSIncr(ctx *parser.SIncrContext) interface{} {
	src, dst := v.evalLV(ctx.Lvalue())
	v.EmitQuad(QBinOp{
		Dst: dst,
		Op:  "+",
		Lhs: src,
		Rhs: LConst{
			Type_: types.TInt{},
			Value: 1,
		},
	})

	return nil
}

func (v *Visitor) VisitSDecr(ctx *parser.SDecrContext) interface{} {
	src, dst := v.evalLV(ctx.Lvalue())
	v.EmitQuad(QBinOp{
		Dst: dst,
		Op:  "-",
		Lhs: src,
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
		v.EmitQuad(QRet{
			Value: v.evalExpr(ctx.Expr()),
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

func (v *Visitor) VisitSFor(ctx *parser.SForContext) interface{} {
	panic("arrays are not supported yet")
	// t, err := v.evalType(ctx.Type_())
	// if err != nil {
	// 	return err
	// }

	// arrType, err := v.evalExpr(ctx.Expr())
	// if err != nil {
	// 	return err
	// }

	// arr, ok := arrType.(TArray)
	// if !ok {
	// 	return NotAnArrayError{
	// 		Ctx:  ctx,
	// 		Expr: ctx.Expr(),
	// 		Type: arrType,
	// 	}
	// }

	// if !SameType(t, arr.Elem) && !v.isSubClass(t, arr.Elem) {
	// 	return UnexpectedTypeError{.Visit(ctx.Expr())
	// 		Got:      t,
	// 	}
	// }

	// defer v.ShadowLocal(ctx.ID().GetText(), arr.Elem)()
	// returns, err := v.evalNonDeclStmt(ctx.Stmt())
	// if err != nil {
	// 	return err
	// }

	// return doesReturn{
	// 	always:    false,
	// 	sometimes: returns.sometimes,
	// }
}

func (v *Visitor) VisitSExp(ctx *parser.SExpContext) interface{} {
	return v.evalSExp(ctx.Expr())
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
