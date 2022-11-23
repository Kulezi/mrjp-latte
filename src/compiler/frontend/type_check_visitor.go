package frontend

import (
	"fmt"
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type varDropper struct {
	drop  func()
	depth int
}

type typeCheckVisitor struct {
	parser.BaseLatteVisitor
	state        *state
	depth        int
	dropperStack []varDropper
	curType      TypeInfo
}

func makeTypeCheckVisitor(s *state) *typeCheckVisitor {
	return &typeCheckVisitor{state: s}
}

func (v *typeCheckVisitor) Run() error {
	return v.Visit(v.state.tree).(error)
}

func (v *typeCheckVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *typeCheckVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	res := make([]interface{}, 0)
	for _, child := range ctx.AllTopDef() {
		if err, ok := v.Visit(child).(error); ok {
			return err
		}
	}

	return res
}

func (v *typeCheckVisitor) VisitTopDef(ctx *parser.TopDefContext) interface{} {
	if fun := ctx.Fundef(); fun != nil {
		return v.Visit(fun)
	}

	return v.Visit(ctx.Classdef())
}

func (v *typeCheckVisitor) ShadowLocal(ident string, typ Type) (drop func()) {
	return v.state.signatures.ShadowLocal(ident, typ, v.depth)
}

func (v *typeCheckVisitor) TypeOfLocal(ident string) (TypeInfo, bool) {
	res, ok := v.state.signatures.Locals[ident]
	return res, ok
}

func (v *typeCheckVisitor) TypeOfGlobal(ident string) (TypeInfo, bool) {
	res, ok := v.state.signatures.Globals[ident]
	return res, ok
}

func (v *typeCheckVisitor) VisitFunDef(ctx *parser.FunDefContext) interface{} {
	ident := ctx.ID().GetText()
	// First try searching for a method in local identifiers, then for a function.
	t, ok := v.TypeOfLocal(ident)
	if !ok {
		t, ok = v.TypeOfGlobal(ident)
		if !ok {
			panic(fmt.Sprintf("typecheck: found undeclared function/method %s at %s", ctx.ID().GetText(), posFromToken(ctx.GetStart())))
		}
	}

	signature, ok := t.Type.(TFun)
	if !ok {
		panic(fmt.Sprintf("typecheck: identifier %s is not a function/method, at %s", ctx.ID().GetText(), posFromToken(ctx.GetStart())))
	}

	v.depth++
	for ident, typ := range signature.Args {
		defer v.ShadowLocal(ident, typ)()
	}
	v.depth--

	return v.Visit(ctx.Block())
}

func (v *typeCheckVisitor) VisitBaseClassDef(ctx *parser.BaseClassDefContext) interface{} {
	t, ok := v.TypeOfGlobal(ctx.ID().GetText())
	if !ok {
		panic(fmt.Sprintf("typecheck: found undeclared class %s at %s", ctx.ID().GetText(), posFromToken(ctx.GetStart())))
	}

	signature, ok := t.Type.(TClass)
	if !ok {
		panic(fmt.Sprintf("typecheck: identifier %s is not a class, at %s", ctx.ID().GetText(), posFromToken(ctx.GetStart())))
	}

	// Put all fields into enviroment.
	for ident, typ := range signature.Fields {
		defer v.ShadowLocal(ident, typ)()
	}

	// Evaluate methods
	for _, field := range ctx.AllField() {
		if err, ok := v.Visit(field).(error); ok {
			return err
		}
	}

	return nil
}

func (v *typeCheckVisitor) VisitDerivedClassDef(ctx *parser.DerivedClassDefContext) interface{} {
	t, ok := v.TypeOfGlobal(ctx.ID(0).GetText())
	if !ok {
		panic(fmt.Sprintf("typecheck: found undeclared class %s at %s", ctx.ID(0).GetText(), posFromToken(ctx.GetStart())))
	}

	signature, ok := t.Type.(TClass)
	if !ok {
		panic(fmt.Sprintf("typecheck: identifier %s is not a class, at %s", ctx.ID(0).GetText(), posFromToken(ctx.GetStart())))
	}

	// Put all fields into enviroment.
	for ident, typ := range signature.Fields {
		defer v.ShadowLocal(ident, typ)()
	}

	// Evaluate methods
	for _, field := range ctx.AllField() {
		if err, ok := v.Visit(field).(error); ok {
			return err
		}
	}

	return nil
}

func (v *typeCheckVisitor) VisitClassFieldDef(ctx *parser.ClassFieldDefContext) interface{} {
	return nil
}

func (v *typeCheckVisitor) VisitClassMethodDef(ctx *parser.ClassMethodDefContext) interface{} {
	return v.Visit(ctx.Fundef())
}

func (v *typeCheckVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.depth++
	defer func() { v.depth-- }()
	fmt.Printf("block %d\n", v.depth)
	for _, stmt := range ctx.AllStmt() {
		if err, ok := v.Visit(stmt).(error); ok {
			return err
		}
	}

	for len(v.dropperStack) > 0 {
		dropper := v.dropperStack[len(v.dropperStack)-1]
		if dropper.depth == v.depth {
			dropper.drop()
			v.dropperStack = v.dropperStack[:len(v.dropperStack)-1]
		}
	}
	return nil
}

func (v *typeCheckVisitor) VisitSBlockStmt(ctx *parser.SBlockStmtContext) interface{} {
	return v.Visit(ctx.Block())
}

func (v *typeCheckVisitor) VisitSDecl(ctx *parser.SDeclContext) interface{} {
	typ := v.Visit(ctx.Type_()).(Type)
	if classRef, ok := typ.BaseType().(TClassRef); ok {
		if _, ok := v.TypeOfGlobal(classRef.String()); !ok {
			return UnknownVariableTypeError{
				typ,
			}
		}
	}

	for _, item := range ctx.AllItem() {
		v.Visit(item)
	}
	return nil
}

func (v *typeCheckVisitor) VisitTSingular(ctx *parser.TSingularContext) interface{} {
	return v.Visit(ctx.Singular_type_())
}

func (v *typeCheckVisitor) VisitTInt(ctx *parser.TIntContext) interface{} {
	return TInt{StartToken: ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitTStr(ctx *parser.TStrContext) interface{} {
	return TString{StartToken: ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitTBool(ctx *parser.TBoolContext) interface{} {
	return TBool{StartToken: ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitTVoid(ctx *parser.TVoidContext) interface{} {
	return TVoid{StartToken: ctx.GetStart()}
}

func (v *typeCheckVisitor) VisitTClass(ctx *parser.TClassContext) interface{} {
	return TClassRef{
		ID: ctx.ID(),
	}
}

func (v *typeCheckVisitor) VisitTArray(ctx *parser.TArrayContext) interface{} {
	typ := v.Visit(ctx.Type_()).(Type)
	return TArray{
		StartToken: ctx.GetStart(),
		Elem:       typ,
	}
}
