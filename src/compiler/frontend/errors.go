package frontend

import "fmt"

type DuplicateIdentifierError struct {
	Ident string
	Pos1  string
	Pos2  string
}

func (e DuplicateIdentifierError) Error() string {
	return fmt.Sprintf(
		`found two identical identifiers
	%s
in the same scope
	at %v
	and %v`,
		e.Ident,
		e.Pos1,
		e.Pos2,
	)
}

type MethodOverrideError struct {
	ParentClass, ChildClass, ParentMethod, ChildMethod Type
	MethodName                                         string
}

func (e MethodOverrideError) Error() string {
	return fmt.Sprintf(
		`can't override method %s declared in parent class %s in child class %s, signature mismatch:
	%s declared at %s 
is not the same as
	%s declared at %s`,
		e.MethodName,
		e.ParentClass,
		e.ChildClass,
		e.ParentMethod,
		e.ParentMethod.Position(),
		e.ChildMethod,
		e.ChildMethod.Position(),
	)
}

type UnknownVariableTypeError struct {
	Type Type
}

func (e UnknownVariableTypeError) Error() string {
	return fmt.Sprintf(
		`can't declare variable(s) of type
	%s
	at %s
as there is no class named %s`,
		e.Type.String(),
		e.Type.Position(),
		e.Type.BaseType().String(),
	)
}
