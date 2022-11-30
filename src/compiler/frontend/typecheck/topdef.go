package typecheck

import (
	"fmt"
	. "latte/compiler/frontend/types"
	"latte/parser"
)

func (v *visitor) VisitTopDef(ctx *parser.TopDefContext) interface{} {
	if fun := ctx.Fundef(); fun != nil {
		return v.Visit(fun)
	}

	return v.Visit(ctx.Classdef())
}

func (v *visitor) VisitFunDef(ctx *parser.FunDefContext) interface{} {
	ident := ctx.ID().GetText()

	t, ok := v.TypeOf(ident)
	if !ok {
		panic(fmt.Sprintf("undeclared identifier %s found at %s", ident, PosFromToken(ctx.GetStart())))
	}

	signature, ok := t.Type.(TFun)
	if !ok {
		panic(fmt.Sprintf("typecheck: identifier %s is not a function/method, at %s", ident, PosFromToken(ctx.GetStart())))
	}

	if _, ok := signature.Result.(TClassRef); ok {
		if _, ok := v.TypeOfGlobal(signature.Result.String()); !ok {
			return UnknownClassError{
				Type: signature.Result,
			}
		}
	}

	v.depth++
	for _, arg := range signature.Args {
		defer v.ShadowLocal(arg.Ident, arg.Type)()
	}
	v.depth--

	v.curFun = &signature
	defer func() { v.curFun = nil }()
	res := v.Visit(ctx.Block())
	if err, ok := res.(error); ok {
		return err
	}

	returns := res.(doesReturn)
	if !returns.always && !SameType(TVoid{}, signature.Result) {
		return MissingReturnError{
			Fun: signature,
		}
	}

	return nil
}

func (v *visitor) VisitBaseClassDef(ctx *parser.BaseClassDefContext) interface{} {
	signature := v.evalClass(ctx.ID().GetText())

	defer v.EnterClass(signature)()
	v.curClass = &signature
	// Evaluate methods
	for _, field := range ctx.AllField() {
		if err, ok := v.Visit(field).(error); ok {
			return err
		}
	}

	v.curClass = nil
	return nil
}

func (v *visitor) VisitDerivedClassDef(ctx *parser.DerivedClassDefContext) interface{} {
	signature := v.evalClass(ctx.ID(0).GetText())

	defer v.EnterClass(signature)()
	v.curClass = &signature
	// Evaluate methods
	for _, field := range ctx.AllField() {
		if err, ok := v.Visit(field).(error); ok {
			return err
		}
	}

	v.curClass = nil
	return nil
}

func (v *visitor) VisitClassFieldDef(ctx *parser.ClassFieldDefContext) interface{} {
	// We still need to check if returned type exists.
	return v.Visit(ctx.Nvtype_())
}
