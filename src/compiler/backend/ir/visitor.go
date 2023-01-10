package ir

import (
	"latte/compiler/config"
	"latte/compiler/frontend/typecheck"
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Addr = uint
type Label struct {
	IsFunction bool
	Name       string
}

// Visitor for intermediate representation generation
type Visitor struct {
	*typecheck.Visitor

	config config.Config

	variableLocations map[string]Location
	allAddresses      map[string]struct{}

	CFG    map[Label]BasicBlock
	Blocks []BasicBlock

	FunInfo  FunInfo
	varCount int

	totalLabels uint
	curBlock    BasicBlock
	lastQuad    Quadruple

	// For boolean expression short circuit evaluation.
	lTrue, lFalse, lNext Label
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

func MakeVisitor(v *typecheck.Visitor, config config.Config) *Visitor {
	return &Visitor{
		Visitor:           v,
		config:            config,
		variableLocations: make(map[string]Location),
		allAddresses:      make(map[string]struct{}),
		CFG:               make(map[Label]BasicBlock),
		FunInfo:           make(FunInfo),
	}
}
