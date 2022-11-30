package typecheck

import (
	"latte/parser"

	. "latte/compiler/frontend/types"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type varDropper struct {
	drop  func()
	depth int
}

type doesReturn struct {
	always, sometimes bool
}

// visitor for typechecking:
//
// Visiting topdefs returns nil on successful typecheck or a meaningful error on a failed one.
// Visiting statements returns doesReturn or an error
// Visiting expressions and types returns Type or an error
type visitor struct {
	parser.BaseLatteVisitor
	// Holds type signatures of identifiers in current scope.
	signatures Signatures
	// Depth of the current scope.
	depth int
	// Stack of variables to unshadow.
	dropperStack []varDropper
	// Signature of the currently checked class
	curClass *TClass
	// Signature of the currently checked function/method.
	curFun *TFun
}

func MakeVisitor(s Signatures) *visitor {
	return &visitor{signatures: s}
}

func (v *visitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	res := make([]interface{}, 0)
	for _, child := range ctx.AllTopDef() {
		if err, ok := v.Visit(child).(error); ok {
			return err
		}
	}

	return res
}

func (v *visitor) VisitClassMethodDef(ctx *parser.ClassMethodDefContext) interface{} {
	// Check method validity.
	return v.Visit(ctx.Fundef())
}
