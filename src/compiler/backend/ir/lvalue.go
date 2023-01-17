package ir

import (
	"latte/compiler/frontend/types"
	"latte/parser"
)

func (v *Visitor) VisitLVField(ctx *parser.LVFieldContext) interface{} {
	Unimplemented("classes are not yet supported\n\t%s", types.PosFromToken(ctx.GetStart()))
	return nil
}

func (v *Visitor) VisitLVArrayRef(ctx *parser.LVArrayRefContext) interface{} {
	array := v.evalExpr(ctx.Expr(0))
	index := v.evalExpr(ctx.Expr(1))
	dst := v.FreshTemp("array_ref", array.Type().BaseType())
	v.EmitQuad(QArrayAccess{
		Array: array,
		Index: index,
		Dst:   dst,
	})

	return dst
}

func (v *Visitor) VisitLVId(ctx *parser.LVIdContext) interface{} {
	ident := ctx.ID().GetText()
	return v.GetLocal(ident)
}

func (v *Visitor) evalLV(ctx parser.ILvalueContext) (loc Location) {
	return v.Visit(ctx).(Location)
}
