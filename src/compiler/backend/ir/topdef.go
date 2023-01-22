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
	label := v.GetFunctionLabel(ident)

	v.StartBlock(label)

	beforeVars := v.varCount

	if v.CurClass != nil {
		_, drop := v.ShadowLocal("self", *v.CurClass)
		defer drop()
	}

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

func (v *Visitor) VisitBaseClassDef(ctx *parser.BaseClassDefContext) interface{} {
	class := v.EvalClass(ctx.ID().GetText())
	defer v.EnterClass(class)()

	// Evaluate methods
	for _, field := range ctx.AllField() {
		if err, ok := v.Visit(field).(error); ok {
			return err
		}
	}

	return nil
}

func (v *Visitor) VisitDerivedClassDef(ctx *parser.DerivedClassDefContext) interface{} {
	class := v.EvalClass(ctx.ID(0).GetText())
	defer v.EnterClass(class)()

	// Evaluate methods
	for _, field := range ctx.AllField() {
		if err, ok := v.Visit(field).(error); ok {
			return err
		}
	}

	return nil
}

func (v *Visitor) VisitClassFieldDef(ctx *parser.ClassFieldDefContext) interface{} {
	// We still need to check if returned type exists.
	return nil
}

func (v *Visitor) VisitClassMethodDef(ctx *parser.ClassMethodDefContext) interface{} {
	// Check method validity.
	return v.Visit(ctx.Fundef())
}
