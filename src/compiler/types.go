package compiler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func posFromToken(t antlr.Token) string {
	return fmt.Sprintf("line %d, column %d", t.GetLine(), t.GetColumn())
}

func SameType(a, b Type) bool {
	return a.String() == b.String()
}

type Type interface {
	String() string
	Position() string
}

type TVoid struct {
	StartToken antlr.Token
}

func (TVoid) String() string {
	return "void"
}

func (t TVoid) Position() string {
	return posFromToken(t.StartToken)
}

type TInt struct {
	StartToken antlr.Token
}

func (TInt) String() string {
	return "int"
}

func (t TInt) Position() string {
	return posFromToken(t.StartToken)
}

type TString struct {
	StartToken antlr.Token
}

func (TString) String() string {
	return "string"
}

func (t TString) Position() string {
	return posFromToken(t.StartToken)
}

type TBool struct {
	StartToken antlr.Token
}

func (TBool) String() string {
	return "boolean"
}

func (t TBool) Position() string {
	return posFromToken(t.StartToken)
}

type TClass struct {
	ID     antlr.TerminalNode
	Fields map[string]Type
	Parent *TClassRef
}

func (t TClass) String() string {
	return t.ID.GetText()
}

func (t TClass) Position() string {
	return posFromToken(t.ID.GetSymbol())
}

func (t TClass) AsRef() TClassRef {
	return TClassRef{
		ID: t.ID,
	}
}

type TClassRef struct {
	ID antlr.TerminalNode
}

func (t TClassRef) String() string {
	return t.ID.GetText()
}

func (t TClassRef) Position() string {
	return posFromToken(t.ID.GetSymbol())
}

type TArray struct {
	StartToken antlr.Token
	Elem       Type
}

func (t TArray) String() string {
	return t.Elem.String() + "[]"
}

func (t TArray) Position() string {
	return posFromToken(t.StartToken)
}

type TFun struct {
	ID     antlr.TerminalNode
	Args   map[string]Type
	Result Type
}

func (f TFun) String() string {
	var res = f.Result.String() + " " + f.ID.GetText() + "("
	comma := false
	for _, v := range f.Args {
		if !comma {
			comma = true
		} else {
			res += ","
		}
		res += v.String()
	}

	res += ")"
	return res
}

func (t TFun) Position() string {
	return posFromToken(t.ID.GetSymbol())
}
