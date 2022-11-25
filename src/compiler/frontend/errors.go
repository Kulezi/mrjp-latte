package frontend

import (
	"fmt"
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type DuplicateIdentifierError struct {
	Ident      string
	Pos1, Pos2 string
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

type UnknownClassError struct {
	Type Type
}

func (e UnknownClassError) Error() string {
	return fmt.Sprintf(
		`unknown class %s referenced
	at %s`,
		e.Type.String(),
		e.Type.Position(),
	)
}

type UnexpectedTypeError struct {
	Expr     parser.IExprContext
	Expected Type
	Got      Type
}

func (e UnexpectedTypeError) Error() string {
	return fmt.Sprintf(
		`expected type
	%s 
but expression
	%s
evaluates to type 
	%s
	at %s`,
		e.Expr.GetText(),
		e.Expected,
		e.Got,
		posFromToken(e.Expr.GetStart()),
	)
}

type UndeclaredIdentifierError struct {
	Ident antlr.TerminalNode
}

func (e UndeclaredIdentifierError) Error() string {
	return fmt.Sprintf(
		`undeclared identifier %s found
	at %s`,
		e.Ident,
		posFromToken(e.Ident.GetSymbol()),
	)
}

func countDimensions(t Type) int {
	if arr, ok := t.(TArray); ok {
		return countDimensions(arr.Elem) + 1
	}

	return 0
}

type ArrayDimensionsMismatchError struct {
	Expr *parser.EArrayRefContext
	Type Type
}

func (e ArrayDimensionsMismatchError) Error() string {
	return fmt.Sprintf(
		`invalid array access in expression
	%s
	at %s,
array of type
	%s
has only %d dimensions while the expression provides %d dimensions`,
		e.Expr.GetText(),
		posFromToken(e.Expr.GetStart()),
		e.Type,
		countDimensions(e.Type),
		len(e.Expr.AllExpr())-1,
	)
}

type ArrayIndexTypeError struct {
	Expr parser.IExprContext
	Type Type
}

func (e ArrayIndexTypeError) Error() string {
	return fmt.Sprintf(
		`expression
	%s
needs to evaluate to type
	int
to be a valid array index, but evalues to type 
	%s
	at %s`,
		e.Expr.GetText(),
		e.Type,
		e.Type.Position(),
	)
}

type ExpectedClassError struct {
	Expr parser.IExprContext
	Got  Type
}

func (e ExpectedClassError) Error() string {
	return fmt.Sprintf(
		`expected the expression
	%s
to evaluate to a class, but got
	%s
	at %s`,
		e.Expr.GetText(),
		e.Got,
		e.Got.Position(),
	)
}

type ArgTypeMismatchError struct {
	Expr         parser.IExprContext
	Type1, Type2 Type
}

func (e ArgTypeMismatchError) Error() string {
	return fmt.Sprintf(
		`expected expression
	%s
arguments to have same type, but got
	%s
	!=
	%s
	at %s`,
		e.Expr.GetText(),
		e.Type1,
		e.Type2,
		posFromToken(e.Expr.GetStart()),
	)
}

type InvalidOpArgsError struct {
	Expr       parser.IExprContext
	Type       Type
	ValidTypes map[string]struct{}
}

func (e InvalidOpArgsError) Error() string {
	return fmt.Sprintf(
		`operator arguments need to have a type from
	%s
	but are of type
	%s
	in expression
	%s
	at %s`,
		e.ValidTypes,
		e.Type,
		e.Expr.GetText(),
		posFromToken(e.Expr.GetStart()),
	)
}

type ExpectedFunctionError struct {
	Expr parser.IExprContext
	Type Type
}

func (e ExpectedFunctionError) Error() string {
	return fmt.Sprintf(
		`expected a function but expression
	%s
evaluates to type
	%s
	at %s`,
		e.Expr.GetText(),
		e.Type,
		posFromToken(e.Expr.GetStart()),
	)
}

type InvalidFunctionArgumentCountError struct {
	Expr *parser.EFunCallContext
	Fun  TFun
}

func (e InvalidFunctionArgumentCountError) Error() string {
	return fmt.Sprintf(
		`function
	%s
expects %d arguments, but %d were provided in call
	%s
	at %s`,
		e.Fun,
		len(e.Fun.Args),
		len(e.Expr.AllExpr()),
		e.Expr.GetText(),
		posFromToken(e.Expr.GetStart()),
	)
}

type NotAFunctionError struct {
	Ident antlr.TerminalNode
	Type  Type
}

func (e NotAFunctionError) Error() string {
	return fmt.Sprintf(
		`identifier %s can't be called as it's not a function, but a value of type
	%s`,
		posFromToken(e.Ident.GetSymbol()),
		e.Type,
	)
}
