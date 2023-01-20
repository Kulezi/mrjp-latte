package typecheck

import (
	. "latte/compiler/frontend/types"
	"latte/parser"
)

func (v *Visitor) VisitLVField(ctx *parser.LVFieldContext) interface{} {
	t, err := v.EvalExpr(ctx.Expr())
	if err != nil {
		return err
	}

	if class, ok := t.(TClass); ok {
		defer v.EnterClass(class)()
	} else if _, ok := t.(TArray); ok {
		defer v.ShadowLocal("length", TReadOnly{Type: TInt{StartToken: ctx.GetStart()}})()
	}

	return v.Visit(ctx.Lvalue())
}

func (v *Visitor) VisitLVArrayRef(ctx *parser.LVArrayRefContext) interface{} {
	t, err := v.EvalExpr(ctx.Expr(0))
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
