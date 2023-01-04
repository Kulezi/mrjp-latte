package ir

import (
	. "latte/compiler/frontend/types"
)

func (v *Visitor) AddLocal(ident string, t Type) (addr Addr, drop func()) {
	addr = v.totalAddresses
	v.variableAdresses[ident] = addr
	v.totalAddresses++

	drop = v.ShadowLocal(ident, t)
	return
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
