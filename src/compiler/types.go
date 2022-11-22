package compiler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func posFromToken(t antlr.Token) string {
	return fmt.Sprintf("line %d, column %d", t.GetLine(), t.GetColumn())
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

type TBaseClass struct {
	Name   string
	Fields map[string]Type
}

func (t TBaseClass) String() string {
	return t.Name
}

func (TBaseClass) Position() string {
	return "unknown"
}

type TDerivedClass struct {
	Name   string
	Fields map[string]Type
	Parent TClassRef
}

func (t TDerivedClass) String() string {
	return t.Name
}

func (TDerivedClass) Position() string {
	return "unknown"
}

type TClassRef struct {
	ID antlr.TerminalNode
}

func (t TClassRef) String() string {
	return string(t.ID.GetText())
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
	var res = f.Result.String() + "("
	for _, v := range f.Args {
		res += v.String()
	}

	res += ")"
	return res
}

func (t TFun) Position() string {
	return posFromToken(t.ID.GetSymbol())
}
