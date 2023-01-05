package ir

import (
	"latte/parser"
)

func (v *Visitor) VisitLVField(ctx *parser.LVFieldContext) interface{} {
	panic("classes are not yet supported")
	// t, err := v.EvalExpr(ctx.Expr())
	// if err != nil {
	// 	return err
	// }

	// if class, ok := t.(TClass); ok {
	// 	defer v.EnterClass(class)()
	// }

	// return v.Visit(ctx.Lvalue())
}

func (v *Visitor) VisitLVArrayRef(ctx *parser.LVArrayRefContext) interface{} {
	panic("arrays are not yet suppoted")
	// t, err := v.EvalExpr(ctx.Expr(0))
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

func (v *Visitor) VisitLVId(ctx *parser.LVIdContext) interface{} {
	ident := ctx.ID().GetText()
	return v.GetLocal(ident)
}

func (v *Visitor) evalLV(ctx parser.ILvalueContext) Location {
	return v.Visit(ctx).(Location)
}
