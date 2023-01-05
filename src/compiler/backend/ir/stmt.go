package ir

import (
	"latte/compiler/frontend/typecheck"
	"latte/compiler/frontend/types"
	"latte/parser"
	"log"
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

// func (v *Visitor) VisitSEmpty(ctx *parser.SEmptyContext) interface{} {
// 	return doesReturn{}
// }

// func (v *Visitor) VisitSBlockStmt(ctx *parser.SBlockStmtContext) interface{} {
// 	return v.Visit(ctx.Block())
// }

func (v *Visitor) VisitSDecl(ctx *parser.SDeclContext) interface{} {
	log.Println("sdecl")
	t, _ := v.Visitor.EvalType(ctx.Nvtype_())

	for _, item := range ctx.AllItem() {
		item, ok := item.(*parser.ItemContext)
		if !ok {
			panic("unexpected antlr behaviour: item is not an ItemContext")
		}
		ident := item.ID().GetText()

		src := LConst{
			Type_: t,
		}

		if item.Expr() != nil {
			src.Value = v.evalSExp(item.Expr())
		} else {
			src.Value = t.DefaultValue()
		}

		dst, drop := v.ShadowLocal(ident, t)
		v.EmitQuad(QMov{
			Src: src,
			Dst: dst,
		})

		v.DropperStack = append(v.DropperStack, typecheck.VarDropper{
			Drop:  drop,
			Depth: v.Depth,
		})
	}

	return nil
}

// func (v *Visitor) VisitSAss(ctx *parser.SAssContext) interface{} {
// 	t, err := v.evalLVType(ctx.Lvalue())
// 	if err != nil {
// 		return err
// 	}
// 	if err := v.ExpectType(t, ctx.Expr()); err != nil {
// 		return err
// 	}
// 	return doesReturn{}
// }

// func (v *Visitor) VisitSIncr(ctx *parser.SIncrContext) interface{} {
// 	t, err := v.evalLVType(ctx.Lvalue())
// 	if err != nil {
// 		return err
// 	}

// 	if !SameType(t, TInt{}) {
// 		return UnexpectedTypeError{
// 			Expr:     ctx,
// 			Expected: TInt{},
// 			Got:      t,
// 		}
// 	}

// 	return doesReturn{}
// }

// func (v *Visitor) VisitSDecr(ctx *parser.SDecrContext) interface{} {
// 	t, err := v.evalLVType(ctx.Lvalue())
// 	if err != nil {
// 		return err
// 	}

// 	if !SameType(t, TInt{}) {
// 		return UnexpectedTypeError{
// 			Expr:     ctx,
// 			Expected: TInt{},
// 			Got:      t,
// 		}
// 	}

// 	return doesReturn{}
// }

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

// func (v *Visitor) VisitSVRet(ctx *parser.SVRetContext) interface{} {
// 	if !SameType(v.CurFun.Result, TVoid{}) {
// 		return MissingReturnValueError{
// 			Ctx:      ctx,
// 			Expected: v.CurFun.Result,
// 		}
// 	}
// 	return doesReturn{
// 		always:    true,
// 		sometimes: true,
// 	}
// }

// func (v *Visitor) evalNonDeclStmt(ctx parser.IStmtContext) (doesReturn, error) {
// 	if _, ok := ctx.(*parser.SDeclContext); ok {
// 		return doesReturn{}, DeclarationWithoutBlockError{
// 			Ctx: ctx,
// 		}
// 	}

// 	return v.evalStmt(ctx)
// }

// func (v *Visitor) VisitSCond(ctx *parser.SCondContext) interface{} {
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

// 	blockReturns, err := v.evalNonDeclStmt(ctx.Stmt())
// 	if err != nil {
// 		return err
// 	}

// 	// If it always executes we are sure of the return.
// 	if b.Value != nil && *b.Value {
// 		return blockReturns
// 	}

// 	return doesReturn{
// 		always:    false,
// 		sometimes: blockReturns.sometimes,
// 	}
// }

// func (v *Visitor) VisitSCondElse(ctx *parser.SCondElseContext) interface{} {
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

// 	retTrue, err := v.evalNonDeclStmt(ctx.Stmt(0))
// 	if err != nil {
// 		return err
// 	}

// 	retFalse, err := v.evalNonDeclStmt(ctx.Stmt(1))
// 	if err != nil {
// 		return err
// 	}

// 	if b.Value != nil {
// 		if *b.Value {
// 			return retTrue
// 		} else {
// 			return retFalse
// 		}
// 	}

// 	return doesReturn{
// 		always:    retFalse.always || retTrue.always,
// 		sometimes: retFalse.sometimes || retTrue.sometimes,
// 	}
// }

// func (v *Visitor) VisitSWhile(ctx *parser.SWhileContext) interface{} {
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

// 	returns, err := v.evalNonDeclStmt(ctx.Stmt())
// 	if err != nil {
// 		return err
// 	}

// 	if b.Value != nil && *b.Value {
// 		return doesReturn{
// 			always:    returns.sometimes,
// 			sometimes: returns.sometimes,
// 		}
// 	}

// 	return doesReturn{
// 		always:    false,
// 		sometimes: returns.sometimes,
// 	}
// }

// func (v *Visitor) VisitSFor(ctx *parser.SForContext) interface{} {
// 	t, err := v.evalType(ctx.Type_())
// 	if err != nil {
// 		return err
// 	}

// 	arrType, err := v.evalExpr(ctx.Expr())
// 	if err != nil {
// 		return err
// 	}

// 	arr, ok := arrType.(TArray)
// 	if !ok {
// 		return NotAnArrayError{
// 			Ctx:  ctx,
// 			Expr: ctx.Expr(),
// 			Type: arrType,
// 		}
// 	}

// 	if !SameType(t, arr.Elem) && !v.isSubClass(t, arr.Elem) {
// 		return UnexpectedTypeError{.Visit(ctx.Expr())
// 			Got:      t,
// 		}
// 	}

// 	defer v.ShadowLocal(ctx.ID().GetText(), arr.Elem)()
// 	returns, err := v.evalNonDeclStmt(ctx.Stmt())
// 	if err != nil {
// 		return err
// 	}

// 	return doesReturn{
// 		always:    false,
// 		sometimes: returns.sometimes,
// 	}
// }

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
