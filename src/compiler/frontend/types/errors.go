package types

import (
	"fmt"
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// https://stackoverflow.com/questions/26524302/how-to-preserve-whitespace-when-we-use-text-attribute-in-antlr4
func showSource(ctx antlr.ParserRuleContext) string {
	if ctx.GetStart() == nil || ctx.GetStop() == nil ||
		ctx.GetStart().GetStart() < 0 || ctx.GetStop().GetStop() < 0 {
		return ctx.GetText()
	}

	return ctx.GetStart().GetInputStream().GetText(
		ctx.GetStart().GetStart(),
		ctx.GetStop().GetStop(),
	)
}

func showTypes(m map[string]struct{}) string {
	res := "["
	comma := false
	for v, _ := range m {
		if comma {
			res += ","
		} else {
			comma = true
		}
		res += v
	}

	return res + "]"
}

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
	Expr     antlr.ParserRuleContext
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
		e.Expected,
		showSource(e.Expr),
		e.Got,
		PosFromToken(e.Expr.GetStart()),
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
		PosFromToken(e.Ident.GetSymbol()),
	)
}

type ArraySizeTypeError struct {
	Expr parser.IExprContext
	Type Type
}

func (e ArraySizeTypeError) Error() string {
	return fmt.Sprintf(
		`expression
	%s
needs to evaluate to type
	int
to be a valid array size, but evalues to type 
	%s
	at %s`,
		showSource(e.Expr),
		e.Type,
		e.Type.Position(),
	)
}

type ExpectedArrayError struct {
	Expr antlr.ParserRuleContext
	Got  Type
}

func (e ExpectedArrayError) Error() string {
	return fmt.Sprintf(
		`expected the expression
	%s
to evaluate to an array, but got
	%s
	at %s`,
		showSource(e.Expr),

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
		showSource(e.Expr),
		e.Type1,
		e.Type2,
		PosFromToken(e.Expr.GetStart()),
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
		showTypes(e.ValidTypes),
		e.Type,
		showSource(e.Expr),
		PosFromToken(e.Expr.GetStart()),
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
		showSource(e.Expr),
		PosFromToken(e.Expr.GetStart()),
	)
}

type NotAFunctionError struct {
	Ident antlr.TerminalNode
	Type  Type
}

func (e NotAFunctionError) Error() string {
	return fmt.Sprintf(
		`identifier %s can't be called as it's not a function, but a value of type
	%s
	at %s`,
		e.Ident,
		e.Type,
		PosFromToken(e.Ident.GetSymbol()),
	)
}

type MissingReturnValueError struct {
	Ctx      antlr.ParserRuleContext
	Expected Type
}

func (e MissingReturnValueError) Error() string {
	return fmt.Sprintf(
		`missing return value
	at %s
expected a value of type
	%s`,
		PosFromToken(e.Ctx.GetStart()),
		e.Expected,
	)
}

type NotAnArrayError struct {
	Ctx, Expr antlr.ParserRuleContext
	Type      Type
}

func (e NotAnArrayError) Error() string {
	return fmt.Sprintf(
		`expected expression
	%s
	at %s
to evaluate to an array, but got value of type
	%s`,
		showSource(e.Expr),
		PosFromToken(e.Ctx.GetStart()),
		e.Type,
	)
}

type MissingReturnError struct {
	Fun TFun
}

func (e MissingReturnError) Error() string {
	return fmt.Sprintf(
		`missing return in function
	%s
	declared at %s`,
		e.Fun,
		e.Fun.Position(),
	)
}

type VoidReturnWithValueError struct {
	Ctx antlr.ParserRuleContext
	Fun TFun
}

func (e VoidReturnWithValueError) Error() string {
	return fmt.Sprintf(
		`can't return from function
	%s
	as its return type is void, but found
	%s
	at %s`,
		e.Fun,
		showSource(e.Ctx),
		PosFromToken(e.Ctx.GetStart()),
	)
}

type MainInvalidSignatureError struct {
	Type Type
}

func (e MainInvalidSignatureError) Error() string {
	return fmt.Sprintf(
		`main should have signature
	int main()
but has
	%s`,
		e.Type,
	)

}

type DeclarationWithoutBlockError struct {
	Ctx antlr.ParserRuleContext
}

func (e DeclarationWithoutBlockError) Error() string {
	return fmt.Sprintf(
		`found variable declaration that needs a new block in
	%s
	at %s`,
		showSource(e.Ctx),
		PosFromToken(e.Ctx.GetStart()))
}

type ConstOutOfRangeError struct {
	Ctx antlr.ParserRuleContext
}

func (e ConstOutOfRangeError) Error() string {
	return fmt.Sprintf(
		`constant 
	%s
	declared at %s
is too big to fit inside int type`,
		showSource(e.Ctx),
		PosFromToken(e.Ctx.GetStart()),
	)
}

type ZeroDivisionError struct {
	Ctx antlr.ParserRuleContext
}

func (e ZeroDivisionError) Error() string {
	return fmt.Sprintf(
		`division by zero in expression
	%s
	at %s`,
		showSource(e.Ctx),
		PosFromToken(e.Ctx.GetStart()),
	)
}
