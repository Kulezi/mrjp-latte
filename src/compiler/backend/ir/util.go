package ir

import (
	"fmt"
	. "latte/compiler/frontend/types"
)

func (v *Visitor) FreshTemp(t Type) LReg {
	addr := v.totalAddresses
	v.totalAddresses++

	return LReg{
		Type_: t,
		Addr:  addr,
	}
}

func (v *Visitor) FreshLabel() Label {
	res := fmt.Sprintf("_%d:", v.totalLabels)
	v.totalLabels++
	return res
}

func (v *Visitor) StartBlock(l Label) {
	v.curBlock = BasicBlock{Label: l}
}

func (v *Visitor) ShadowLocal(ident string, t Type) (location Location) {
	loc := v.FreshTemp(t)
	v.variableLocations[ident] = loc
	v.Signatures.Locals[ident] = TypeInfo{
		Type: t,
	}

	return loc
}

func (v *Visitor) GetLocal(ident string) Location {
	return v.variableLocations[ident]
}

func (v *Visitor) EmitQuad(q Quadruple) {
	v.lastQuad = q
	v.curBlock.Ops = append(v.curBlock.Ops, q)
	if q.IsJump() {
		if v.curBlock.Label == "" {
			v.curBlock.Label = v.FreshLabel()
		}
		v.CFG[v.curBlock.Label] = v.curBlock
		v.Blocks = append(v.Blocks, v.curBlock)
		v.curBlock = BasicBlock{}
	}
}

func (v *Visitor) TypeOfLocal(ident string) (TypeInfo, bool) {
	res, ok := v.Signatures.Locals[ident]
	return res, ok
}

func (v *Visitor) TypeOfGlobal(ident string) (TypeInfo, bool) {
	res, ok := v.Signatures.Globals[ident]
	return res, ok
}

func (v *Visitor) TypeOf(ident string) (TypeInfo, bool) {
	if res, ok := v.TypeOfLocal(ident); ok {
		return res, ok
	}

	return v.TypeOfGlobal(ident)
}
