package ir

import (
	"latte/compiler/frontend/types"
	"latte/parser"
	"log"
)

func (v *Visitor) VisitLVFieldArrayRef(ctx *parser.LVFieldArrayRefContext) interface{} {
	Unimplemented("classes are not yet supported\n\t%s", types.PosFromToken(ctx.GetStart()))
	return nil

}

func (v *Visitor) VisitLVFieldMethodCall(ctx *parser.LVFieldMethodCallContext) interface{} {
	Unimplemented("classes are not yet supported\n\t%s", types.PosFromToken(ctx.GetStart()))
	return nil
}

func (v *Visitor) VisitLVField(ctx *parser.LVFieldContext) interface{} {
	lhs := v.evalExpr(ctx.Expr())
	ident := ctx.ID().GetText()
	log.Printf("\n####\n%#v\n$$$\n, %#v\n####\n", lhs, lhs.Type())

	var class types.TClass
	switch t := lhs.Type().(type) {
	case types.TClassRef:
		class = v.EvalClass(t.ID.GetText())
	case types.TClass:
		class = t
	default:
		panic("field access as lvalue happened on non class type")
	}

	fieldInfo := class.Fields[ident]
	dst := v.FreshTemp("class_field", fieldInfo.Type)
	v.EmitQuad(QArrayAccess{
		Array: lhs,
		Index: LConst{Type_: types.TInt{}, Value: fieldInfo.Offset - 1},
		Dst:   dst,
	})
	return LMem{Type_: dst.Type(), Addr: dst}
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

	return LMem{Type_: array.Type().BaseType(), Addr: dst}
}

func (v *Visitor) VisitLVId(ctx *parser.LVIdContext) interface{} {
	ident := ctx.ID().GetText()
	return v.GetLocal(ident)
}

func (v *Visitor) evalLV(ctx parser.ILvalueContext) (loc Location) {
	return v.Visit(ctx).(Location)
}
