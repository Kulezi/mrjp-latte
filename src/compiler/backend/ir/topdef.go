package ir

import (
	. "latte/compiler/frontend/types"
	"latte/parser"
)

func (v *Visitor) VisitTopDef(ctx *parser.TopDefContext) interface{} {
	if fun := ctx.Fundef(); fun != nil {
		return v.Visit(fun)
	}

	return v.Visit(ctx.Classdef())
}

func (v *Visitor) VisitFunDef(ctx *parser.FunDefContext) interface{} {
	ident := ctx.ID().GetText()

	t, _ := v.TypeOf(ident)
	signature := t.Type.(TFun)

	v.Depth++
	for _, arg := range signature.Args {
		defer v.Visitor.ShadowLocal(arg.Ident, arg.Type)()
	}
	v.Depth--

	v.CurFun = &signature
	defer func() { v.CurFun = nil }()

	v.curBlock = BasicBlock{Label: ident + ":"}

	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitBaseClassDef(ctx *parser.BaseClassDefContext) interface{} {
	panic("Classes are not supported in this revision.")
}

func (v *Visitor) VisitDerivedClassDef(ctx *parser.DerivedClassDefContext) interface{} {
	panic("Classes are not supported in this revision.")
}
