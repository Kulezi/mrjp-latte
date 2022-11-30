package typecheck

import (
	"fmt"
	. "latte/compiler/frontend/types"
	"latte/parser"
)

func (v *visitor) evalClass(ident string) TClass {
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

func (v *visitor) ConflictLocal(ident string) (TypeInfo, bool) {
	return v.signatures.ConflictLocal(ident, v.depth)
}

func (v *visitor) ShadowLocal(ident string, t Type) (drop func()) {
	return v.signatures.ShadowLocal(ident, t, v.depth)
}

func (v *visitor) TypeOfLocal(ident string) (TypeInfo, bool) {
	res, ok := v.signatures.Locals[ident]
	return res, ok
}

func (v *visitor) TypeOfGlobal(ident string) (TypeInfo, bool) {
	res, ok := v.signatures.Globals[ident]
	return res, ok
}

func (v *visitor) TypeOf(ident string) (TypeInfo, bool) {
	if res, ok := v.TypeOfLocal(ident); ok {
		return res, ok
	}

	return v.TypeOfGlobal(ident)
}

func (v *visitor) EnterClass(signature TClass) (exit func()) {
	v.depth++
	// Put all fields into enviroment.
	var exits []func()
	for ident, t := range signature.Fields {
		exits = append(exits, v.ShadowLocal(ident, t))
	}

	return func() {
		for _, exit := range exits {
			exit()
		}
		v.depth--
	}
}

func (v *visitor) isSubClass(expected, child Type) bool {
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

func (v *visitor) ExpectType(expected Type, ctx parser.IExprContext) error {
	if ctx == nil {
		return nil
	}

	got, err := v.evalExpr(ctx)
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
