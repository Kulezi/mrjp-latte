package typecheck

import (
	"fmt"
	. "latte/compiler/frontend/types"
	"latte/parser"
)

func (v *Visitor) EvalClass(ident string) TClass {
	t, ok := v.TypeOfGlobal(ident)
	if !ok {
		panic(fmt.Sprintf("typecheck: found undeclared class %s", ident))
	}

	signature, ok := t.Type.(TClass)
	if !ok {
		panic(fmt.Sprintf("typecheck: identifier %s is not a class", ident))
	}

	return signature
}

func (v *Visitor) ConflictLocal(ident string) (TypeInfo, bool) {
	return v.Signatures.ConflictLocal(ident, v.Depth)
}

func (v *Visitor) ShadowLocal(ident string, t Type) (drop func()) {
	return v.Signatures.ShadowLocal(ident, t, v.Depth)
}

func (v *Visitor) TypeOfLocal(ident string) (TypeInfo, bool) {
	res, ok := v.Signatures.Locals[ident]
	return res, ok
}

func (v *Visitor) TypeOfGlobal(ident string) (TypeInfo, bool) {
	res, ok := v.Signatures.Globals[ident]
	return res, ok
}

func (v *Visitor) TypeOf(ident string) (TypeInfo, bool) {
	if res, ok := v.TypeOfLocal(ident); ok {
		return res, ok
	}

	return v.TypeOfGlobal(ident)
}

func (v *Visitor) EnterType(t Type, lvalue bool) (exit func()) {
	oldLocals := v.Signatures.Locals
	oldGlobals := v.Signatures.Globals
	v.Signatures.Locals = make(Env)
	v.Signatures.Globals = make(Env)

	switch t := t.(type) {
	case TClass:
		for ident, t := range t.Fields {
			v.ShadowLocal(ident, t.Type)
		}
	case TArray:
		if lvalue {
			v.ShadowLocal("length", TReadOnly{Type: TInt{}})
		} else {
			v.ShadowLocal("length", TInt{})
		}
	}

	return func() {
		v.Signatures.Locals = oldLocals
		v.Signatures.Globals = oldGlobals
	}
}

func (v *Visitor) EnterClass(signature TClass) (exit func()) {
	v.CurClass = &signature
	v.Depth++
	// Put all fields into enviroment.
	oldLocals := v.Signatures.Locals
	v.Signatures.Locals = make(Env)
	for ident, t := range signature.Fields {
		v.ShadowLocal(ident, t.Type)
	}

	v.ShadowLocal("self", signature)

	return func() {
		v.Signatures.Locals = oldLocals
		v.Depth--
		v.CurClass = nil
	}
}

func (v *Visitor) isSubClass(expected, child Type) bool {
	class, ok := child.(TClass)
	if !ok {
		return false
	}

	for !SameType(expected, class) && class.Parent != nil {
		parentRef := *class.Parent
		parent, _ := v.TypeOfGlobal(parentRef.String())
		class = parent.Type.(TClass)
	}

	return SameType(expected, class)
}

func (v *Visitor) ExpectType(expected Type, ctx parser.IExprContext) error {
	if ctx == nil {
		return nil
	}

	got, err := v.EvalExpr(ctx)
	if err != nil {
		return err
	}
	if !SameType(expected, got) && !v.isSubClass(expected, got) {
		return UnexpectedTypeError{
			Expr:     ctx,
			Expected: expected,
			Got:      got,
		}
	}

	return nil
}
