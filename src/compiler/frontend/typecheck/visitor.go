package typecheck

import (
	"latte/parser"

	. "latte/compiler/frontend/types"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type varDropper struct {
	Drop  func()
	Depth int
}

type doesReturn struct {
	always, sometimes bool
}

// Visitor for typechecking:
//
// Visiting topdefs returns nil on successful typecheck or a meaningful error on a failed one.
// Visiting statements returns doesReturn or an error
// Visiting expressions and types returns Type or an error
type Visitor struct {
	parser.BaseLatteVisitor
	// Holds type Signatures of identifiers in current scope.
	Signatures Signatures
	// Depth of the current scope.
	Depth int
	// Stack of variables to unshadow.
	DropperStack []varDropper
	// Signature of the currently checked class
	CurClass *TClass
	// Signature of the currently checked function/method.
	CurFun *TFun
}

func MakeVisitor(s Signatures) *Visitor {
	return &Visitor{Signatures: s}
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

func (v *Visitor) VisitClassMethodDef(ctx *parser.ClassMethodDefContext) interface{} {
	// Check method validity.
	return v.Visit(ctx.Fundef())
}
