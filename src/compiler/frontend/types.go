package frontend

import (
	"fmt"
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Env map[string]TypeInfo

type TypeInfo struct {
	Type
	Depth int
}

type state struct {
	tree       parser.IProgramContext
	signatures Signatures
}

type Signatures struct {
	Globals Env
	Locals  Env
	Parent  map[string]TClassRef
}

func MakeSignatures() Signatures {
	return Signatures{
		Globals: make(Env),
		Locals:  make(Env),
		Parent:  make(map[string]TClassRef),
	}
}

func (s *Signatures) ConflictGlobal(ident string) (TypeInfo, bool) {
	if v, ok := s.Globals[ident]; ok {
		return v, true
	}

	return TypeInfo{}, false
}

func (s *Signatures) ConflictLocal(ident string, depth int) (TypeInfo, bool) {
	if v, ok := s.Locals[ident]; ok {
		if v.Depth == depth {
			return v, true
		}
	}

	return TypeInfo{}, false
}

func (s *Signatures) ReplaceGlobal(ident string, t Type) {
	s.Globals[ident] = TypeInfo{
		Type:  t,
		Depth: 0,
	}
}

// Declare identifier with type t under assumption that it's not a redeclaration
// in the same scope.
func (s *Signatures) ShadowGlobal(ident string, t Type, depth int) (drop func()) {
	old, ok := s.Globals[ident]
	s.Globals[ident] = TypeInfo{
		Type:  t,
		Depth: depth,
	}

	if ok {
		if old.Depth == depth {
			panic("declare: variable redeclared in the same scope")
		}

		return func() {
			s.Globals[ident] = old
		}
	}

	return func() {
		delete(s.Globals, ident)
	}
}

// Declare identifier with type t under assumption that it's not a redeclaration
// in the same scope.
func (s *Signatures) ShadowLocal(ident string, t Type, depth int) (drop func()) {
	old, ok := s.Locals[ident]
	s.Locals[ident] = TypeInfo{
		Type:  t,
		Depth: depth,
	}

	if ok {
		if old.Depth == depth {
			panic("declare: variable redeclared in the same scope")
		}

		return func() {
			s.Locals[ident] = old
		}
	}

	return func() {
		delete(s.Locals, ident)
	}
}

func posFromToken(t antlr.Token) string {
	return fmt.Sprintf("line %d, column %d", t.GetLine(), t.GetColumn())
}

func sameType(a, b Type) bool {
	return a.String() == b.String()
}

var validRelOpArg = map[string]struct{}{
	"int":    {},
	"string": {},
}

var validAddOpArg = validRelOpArg
var validSubOpArg = map[string]struct{}{
	"int": {},
}

var validMulOpArg = map[string]struct{}{
	"int": {},
}

type Type interface {
	String() string
	Position() string
	BaseType() Type
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

func (t TVoid) BaseType() Type { return t }

type TInt struct {
	StartToken antlr.Token
}

func (TInt) String() string {
	return "int"
}

func (t TInt) Position() string {
	return posFromToken(t.StartToken)
}

func (t TInt) BaseType() Type { return t }

type TString struct {
	StartToken antlr.Token
}

func (TString) String() string {
	return "string"
}

func (t TString) Position() string {
	return posFromToken(t.StartToken)
}

func (t TString) BaseType() Type { return t }

type TBool struct {
	StartToken antlr.Token
}

func (TBool) String() string {
	return "boolean"
}

func (t TBool) Position() string {
	return posFromToken(t.StartToken)
}

func (t TBool) BaseType() Type { return t }

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

func (t TClass) BaseType() Type { return t }

type TClassRef struct {
	ID antlr.TerminalNode
}

func (t TClassRef) String() string {
	return t.ID.GetText()
}

func (t TClassRef) Position() string {
	return posFromToken(t.ID.GetSymbol())
}

func (t TClassRef) BaseType() Type { return t }

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

func (t TArray) BaseType() Type { return t.Elem.BaseType() }

type FArg struct {
	Ident string
	Type  Type
}

type TFun struct {
	ID     antlr.TerminalNode
	Args   []FArg
	Result Type
}

func (t TFun) String() string {
	var res = t.Result.String() + " " + t.ID.GetText() + "("
	comma := false
	for _, v := range t.Args {
		if !comma {
			comma = true
		} else {
			res += ","
		}
		res += v.Type.String()
	}

	res += ")"
	return res
}

func (t TFun) Position() string {
	return posFromToken(t.ID.GetSymbol())
}

func (t TFun) BaseType() Type { return t }
