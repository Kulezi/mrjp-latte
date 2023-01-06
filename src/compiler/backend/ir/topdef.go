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
		loc, drop := v.ShadowLocal(arg.Ident, arg.Type)
		v.EmitLoadArg(loc, arg.Type)
		defer drop()
	}
	v.Depth--

	v.CurFun = &signature
	defer func() { v.CurFun = nil }()

	label := v.GetFunctionLabel(ident)
	v.StartBlock(label)

	v.Visit(ctx.Block())

	// Emit a return, in case of non-void it will be unreachable and optimized away later.
	v.EmitQuad(QVRet{})

	return nil
}

func (v *Visitor) VisitBaseClassDef(ctx *parser.BaseClassDefContext) interface{} {
	panic("Classes are not supported in this revision.")
}

func (v *Visitor) VisitDerivedClassDef(ctx *parser.DerivedClassDefContext) interface{} {
	panic("Classes are not supported in this revision.")
}
