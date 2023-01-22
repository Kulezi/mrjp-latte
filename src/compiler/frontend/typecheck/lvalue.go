package typecheck

import (
	. "latte/compiler/frontend/types"
	"latte/parser"
)

func (v *Visitor) VisitLVArrayRef(ctx *parser.LVArrayRefContext) interface{} {
	t, err := v.EvalExpr(ctx.Expr(0))
	if err != nil {
		return err
	}

	arr, ok := t.(TArray)
	if !ok {
		return NotAnArrayError{Ctx: ctx, Expr: ctx.Expr(0), Type: t}
	}

	if err := v.ExpectType(TInt{}, ctx.Expr(1)); err != nil {
		return err
	}

	return arr.Elem
}

func (v *Visitor) VisitLVFieldArrayRef(ctx *parser.LVFieldArrayRefContext) interface{} {
	lhs, err := v.EvalExpr(ctx.Expr(0))
	if err != nil {
		return err
	}

	drop := v.EnterType(lhs, true)
	t, ok := v.TypeOfLocal(ctx.ID().GetText())
	if !ok {
		return UndeclaredIdentifierError{
			Ident: ctx.ID(),
		}
	}
	drop()

	arr, ok := t.Type.(TArray)
	if !ok {
		return NotAnArrayError{Ctx: ctx, Expr: ctx.Expr(0), Type: t}
	}

	if err := v.ExpectType(TInt{}, ctx.Expr(1)); err != nil {
		return err
	}

	return arr.Elem
}

func (v *Visitor) VisitLVField(ctx *parser.LVFieldContext) interface{} {
	lhs, err := v.EvalExpr(ctx.Expr())
	if err != nil {
		return err
	}

	drop := v.EnterType(lhs, true)
	t, ok := v.TypeOfLocal(ctx.ID().GetText())
	if !ok {
		return UndeclaredIdentifierError{
			Ident: ctx.ID(),
		}
	}
	drop()

	return t.Type
}

func (v *Visitor) VisitLVId(ctx *parser.LVIdContext) interface{} {
	t, ok := v.TypeOfLocal(ctx.ID().GetText())
	if !ok {
		return UndeclaredIdentifierError{
			Ident: ctx.ID(),
		}
	}

	return t.Type
}

func (v *Visitor) evalLVType(ctx parser.ILvalueContext) (Type, error) {
	res := v.Visit(ctx)
	if err, ok := res.(error); ok {
		return nil, err
	}

	return res.(Type), nil
}
