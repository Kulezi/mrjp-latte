package ir

import (
	"latte/compiler/frontend/types"
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
	label := v.GetFunctionLabel(ident)

	v.StartBlock(label)

	beforeVars := v.varCount
	for _, arg := range signature.Args {
		_, drop := v.ShadowLocal(arg.Ident, arg.Type)
		defer drop()
	}
	v.Depth--

	v.CurFun = &signature
	defer func() { v.CurFun = nil }()

	v.Visit(ctx.Block())
	afterVars := v.varCount
	v.FunInfo[label] = VarInfo{
		Signature:     signature,
		Function:      label,
		Offset:        beforeVars,
		VariableCount: afterVars - beforeVars,
	}

	// Emit a return, in case of non-void it will be unreachable and optimized away later.
	v.EmitQuad(QVRet{})
	return nil
}

// TODO: Methods below should generate:
// Class VTable, first field being the class constructor.

func (v *Visitor) VisitBaseClassDef(ctx *parser.BaseClassDefContext) interface{} {

	Unimplemented("Classes are not supported in this revision\n\tFound one at %s", types.PosFromToken(ctx.GetStart()))
	return nil
}

func (v *Visitor) VisitDerivedClassDef(ctx *parser.DerivedClassDefContext) interface{} {
	Unimplemented("Classes are not supported in this revision\n\tFound one at %s", types.PosFromToken(ctx.GetStart()))
	return nil
}
