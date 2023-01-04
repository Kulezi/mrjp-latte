package ir

import (
	"latte/compiler/frontend/typecheck"
)

type Addr = uint
type Label = string

// Visitor for intermediate representation generation
type Visitor struct {
	*typecheck.Visitor
	variableAdresses map[string]Addr
	totalAddresses   uint

	cfg         map[Label]BasicBlock
	totalLabels uint
	curBlock    BasicBlock
}

func MakeVisitor(v *typecheck.Visitor) *Visitor {
	return &Visitor{
		Visitor:          v,
		variableAdresses: make(map[string]uint),
	}
}
