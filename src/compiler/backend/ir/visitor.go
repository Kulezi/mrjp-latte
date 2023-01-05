package ir

import (
	"latte/compiler/frontend/typecheck"
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Addr = uint
type Label = string

// Visitor for intermediate representation generation
type Visitor struct {
	*typecheck.Visitor
	variableLocations map[string]Location
	totalAddresses    uint

	CFG    map[Label]BasicBlock
	Blocks []BasicBlock

	totalLabels uint
	curBlock    BasicBlock
	lastQuad    Quadruple
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	res := make([]interface{}, 0)
	for _, child := range ctx.AllTopDef() {
		if err, ok := v.Visit(child).(error); ok {
			return err
		}
	}

	return res
}

func MakeVisitor(v *typecheck.Visitor) *Visitor {
	return &Visitor{
		Visitor:           v,
		variableLocations: make(map[string]Location),
		CFG:               make(map[string]BasicBlock),
	}
}
