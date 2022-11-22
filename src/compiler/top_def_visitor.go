package compiler

import (
	"fmt"
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// GlobalDeclVisitor is responsible for checking:
//
// - duplicate global identifiers and function argument identifiers
//
// - cyclic inheritance
type GlobalDeclVisitor struct {
	parser.BaseLatteVisitor
	Globals  map[string]Type
	Parent   map[string]TClassRef
	inMethod bool
}

func MakeGlobalDeclVisitor() *GlobalDeclVisitor {
	return &GlobalDeclVisitor{
		Globals: make(map[string]Type),
		Parent:  make(map[string]TClassRef),
	}
}

func (v *GlobalDeclVisitor) createGlobal(id string, t Type) error {
	if v, ok := v.Globals[id]; ok {
		return DuplicateIdentifierError{
			Ident: id,
			Pos1:  v.Position(),
			Pos2:  t.Position(),
		}
	}
	v.Globals[id] = t
	return nil
}

type visitState int

const (
	unvisited visitState = iota
	onStack
	popped
)

// Finds the cycle that class is on, returns nil if it doesn't exist.
func (v *GlobalDeclVisitor) findCycle(class TClassRef, stack []TClassRef, vis map[string]visitState) []TClassRef {
	if vis[class.String()] == popped {
		return nil
	}

	// Inheritance graph has only closed simple cycles, so non-inheriting class can't be on it.
	if _, ok := v.Parent[class.String()]; !ok {
		return nil
	}

	switch vis[class.String()] {
	case unvisited:
		vis[class.String()] = onStack
		stack = append(stack, class)
		if cycle := v.findCycle(v.Parent[class.String()], stack, vis); cycle != nil {
			return cycle
		}
		vis[class.String()] = popped
	case onStack:
		for i, v := range stack {
			if v.String() == class.String() {
				return stack[i:]
			}
		}
	}

	return nil
}

func (v *GlobalDeclVisitor) CheckCyclicInheritance() error {
	vis := make(map[string]visitState)
	for _, ref := range v.Globals {
		var ok bool
		class, ok := ref.(TClass)
		if !ok {
			continue
		}

		if cycle := v.findCycle(class.AsRef(), nil, vis); cycle != nil {
			return fmt.Errorf("found cyclic inheritance: %v", cycle)
		}
	}

	return nil
}

func (v *GlobalDeclVisitor) Run(tree antlr.ParseTree) error {
	if err, ok := v.Visit(tree).(error); ok && err != nil {
		return err
	}

	return v.CheckCyclicInheritance()
}

func (v *GlobalDeclVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *GlobalDeclVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	res := make([]interface{}, 0)
	for _, child := range ctx.AllTopDef() {
		if err, ok := v.Visit(child).(error); ok {
			return err
		}
	}

	return res
}

func (v *GlobalDeclVisitor) VisitTopDef(ctx *parser.TopDefContext) interface{} {
	if fun := ctx.Fundef(); fun != nil {
		return v.Visit(fun)
	}

	return v.Visit(ctx.Classdef())
}

// Returns the function signature as a Type.
func (v *GlobalDeclVisitor) VisitFunDef(ctx *parser.FunDefContext) interface{} {
	fun := TFun{
		ID:     ctx.ID(),
		Args:   make(map[string]Type),
		Result: v.Visit(ctx.Type_()).(Type),
	}

	if ctx.Arg() != nil {
		args := v.Visit(ctx.Arg())
		if err, ok := args.(error); ok {
			return err
		}
		fun.Args = args.(map[string]Type)
	}

	if !v.inMethod {
		if err := v.createGlobal(fun.ID.GetText(), fun); err != nil {
			return err
		}
	}

	return fun
}

func (v *GlobalDeclVisitor) VisitArg(ctx *parser.ArgContext) interface{} {
	args := make(map[string]Type)
	for i, id := range ctx.AllID() {
		typ := ctx.Type_(i)
		if v, ok := args[id.GetText()]; ok {
			return DuplicateIdentifierError{
				Ident: id.GetText(),
				Pos1:  v.Position(),
				Pos2:  posFromToken(typ.GetStart()),
			}
		}
		// FIXME: eval type
		args[id.GetText()] = v.Visit(typ).(Type)
	}

	return args
}

func (v *GlobalDeclVisitor) VisitTSingular(ctx *parser.TSingularContext) interface{} {
	return v.Visit(ctx.Singular_type_())
}

func (v *GlobalDeclVisitor) VisitTInt(ctx *parser.TIntContext) interface{} {
	return TInt{StartToken: ctx.GetStart()}
}

func (v *GlobalDeclVisitor) VisitTStr(ctx *parser.TStrContext) interface{} {
	return TString{StartToken: ctx.GetStart()}
}

func (v *GlobalDeclVisitor) VisitTBool(ctx *parser.TBoolContext) interface{} {
	return TBool{StartToken: ctx.GetStart()}
}

func (v *GlobalDeclVisitor) VisitTVoid(ctx *parser.TVoidContext) interface{} {
	return TVoid{StartToken: ctx.GetStart()}
}

func (v *GlobalDeclVisitor) VisitTClass(ctx *parser.TClassContext) interface{} {
	return TClassRef{
		ID: ctx.ID(),
	}
}

func (v *GlobalDeclVisitor) VisitTArray(ctx *parser.TArrayContext) interface{} {
	typ := v.Visit(ctx.Type_()).(Type)
	return TArray{
		StartToken: ctx.GetStart(),
		Elem:       typ,
	}
}

type field struct {
	ID   string
	Type Type
}

func (v *GlobalDeclVisitor) collectFields(fields []parser.IFieldContext) (map[string]Type, error) {
	ret := make(map[string]Type)
	for _, fieldCtx := range fields {
		res := v.Visit(fieldCtx)
		if err, ok := res.(error); ok {
			return nil, err
		}

		f := res.(field)
		ret[f.ID] = f.Type
	}

	return ret, nil
}

func (v *GlobalDeclVisitor) VisitBaseClassDef(ctx *parser.BaseClassDefContext) interface{} {
	fields, err := v.collectFields(ctx.AllField())
	if err != nil {
		return err
	}

	class := TClass{
		ID:     ctx.ID(),
		Fields: fields,
	}

	return v.createGlobal(class.String(), class)
}

// Returns the function signature as a Type.
func (v *GlobalDeclVisitor) VisitDerivedClassDef(ctx *parser.DerivedClassDefContext) interface{} {
	parent := TClassRef{ID: ctx.ID(1)}
	fields, err := v.collectFields(ctx.AllField())
	if err != nil {
		return err
	}

	class := TClass{
		ID:     ctx.ID(0),
		Fields: fields,
		Parent: &parent,
	}

	v.Parent[class.String()] = parent
	return v.createGlobal(class.String(), class)
}

func (v *GlobalDeclVisitor) VisitClassFieldDef(ctx *parser.ClassFieldDefContext) interface{} {
	return field{
		ID:   ctx.ID().GetText(),
		Type: v.Visit(ctx.Type_()).(Type),
	}
}

func (v *GlobalDeclVisitor) VisitClassMethodDef(ctx *parser.ClassMethodDefContext) interface{} {
	v.inMethod = true
	res := v.Visit(ctx.Fundef())
	v.inMethod = false
	if err, ok := res.(error); ok {
		return err
	}

	fun := res.(TFun)
	return field{
		ID:   fun.ID.GetText(),
		Type: fun,
	}
}

var _ parser.LatteVisitor = &GlobalDeclVisitor{}
