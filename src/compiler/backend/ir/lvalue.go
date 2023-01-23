package ir

import (
	"latte/compiler/frontend/types"
	"latte/parser"
)

func (v *Visitor) VisitLVFieldArrayRef(ctx *parser.LVFieldArrayRefContext) interface{} {
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

	return LMem{Type_: array.Type().BaseType(), Addr: ptr}
}

func (v *Visitor) VisitLVField(ctx *parser.LVFieldContext) interface{} {
	lhs := v.evalExpr(ctx.Expr())
	ident := ctx.ID().GetText()

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
