package ir

import (
	"fmt"
	"latte/compiler/config"
	"latte/compiler/frontend/typecheck"
	"latte/compiler/frontend/types"
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Label struct {
	IsFunction bool
	Name       string
}

func (l Label) String() string {
	return l.Name
}

type fname struct {
	class string
	name  string
}

// Visitor for intermediate representation generation
type Visitor struct {
	*typecheck.Visitor

	config config.Config

	variableLocations map[string]Location
	allAddresses      map[string]struct{}

	functionLabels map[fname]Label

	VTables map[string]Label

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
	visitor := &Visitor{
		Visitor:           v,
		config:            config,
		variableLocations: make(map[string]Location),
		allAddresses:      make(map[string]struct{}),
		functionLabels:    make(map[fname]Label),
		CFG:               make(map[Label]BasicBlock),
		FunInfo:           make(FunInfo),
	}

	// Prealloc labels for all VTables.
	for _, v := range v.Signatures.Globals {
		if class, ok := v.Type.(types.TClass); ok {
			ident := class.ID.GetText()
			visitor.VTables[class.ID.GetText()] = visitor.FreshLabel(fmt.Sprintf("_%s_vtable_", ident))
		}
	}

	return visitor
}
