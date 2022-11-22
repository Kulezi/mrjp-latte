package compiler

import (
	"fmt"
	"latte/parser"
)

type TypeCheckVisitor struct {
	parser.BaseLatteVisitor
}

func (v *TypeCheckVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	fmt.Println("todo: typecheck")
	return v.VisitChildren(ctx)
}
