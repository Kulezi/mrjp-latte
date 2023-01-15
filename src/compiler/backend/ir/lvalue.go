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
	Unimplemented("arrays are not yet suppoted\n\t%s", types.PosFromToken(ctx.GetStart()))
	return nil
}

func (v *Visitor) VisitLVId(ctx *parser.LVIdContext) interface{} {
	ident := ctx.ID().GetText()
	old := v.GetLocal(ident)
	new := v.FreshTemp("lv_new", old.Type())
	new.Variable = ident
	new.Index = v.variableLocations[ident].(LReg).Index
	v.variableLocations[ident] = new

	return locPair{
		old: old,
		new: new,
	}
}

type locPair struct {
	old, new Location
}

func (v *Visitor) evalLV(ctx parser.ILvalueContext) (old, new Location) {
	locs := v.Visit(ctx).(locPair)
	return locs.old, locs.new
}
