package frontend

import (
	"fmt"
	"latte/parser"
)

type typeCheckVisitor struct {
	parser.BaseLatteVisitor
}

func (v *typeCheckVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	fmt.Println("todo: typecheck")
	return v.VisitChildren(ctx)
}
