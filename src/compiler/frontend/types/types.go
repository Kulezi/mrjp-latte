package types

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Env map[string]TypeInfo

type TypeInfo struct {
	Type
	Depth int
}

type Signatures struct {
	Globals Env
	Locals  Env
	Parent  map[string]TClassRef
}

func MakeSignatures() Signatures {
	return Signatures{
		Globals: Env{
			"printInt": TypeInfo{
				Type: TFun{
					Ident: "printInt",
					Args: []FArg{{
						Ident: "x",
						Type:  TInt{},
					}},
					Result: TVoid{},
				},
			},
			"printString": TypeInfo{
				Type: TFun{
					Ident: "printString",
					Args: []FArg{{
						Ident: "x",
						Type:  TString{},
					}},
					Result: TVoid{},
				},
			},
			"error": TypeInfo{
				Type: TFun{
					Ident:  "error",
					Result: TVoid{IsReturn: true},
				},
			},
			"readInt": TypeInfo{
				Type: TFun{
					Ident:  "readInt",
					Result: TInt{},
				},
			},
			"readString": TypeInfo{
				Type: TFun{
					Ident:  "readString",
					Result: TString{},
				},
			},
		},
		Locals: make(Env),
		Parent: make(map[string]TClassRef),
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

func PosFromToken(t antlr.Token) string {
	if t == nil {
		return "unknown position"
	}
	return fmt.Sprintf("line %d, column %d", t.GetLine(), t.GetColumn())
}

func SameType(a, b Type) bool {
	return a.String() == b.String()
}

type Type interface {
	String() string
	Position() string
	BaseType() Type
	Const() (value interface{}, ok bool)
	DefaultValue() interface{}
}

type TVoid struct {
	StartToken antlr.Token
	IsReturn   bool
}

func (TVoid) String() string {
	return "void"
}

func (t TVoid) Position() string {
	return PosFromToken(t.StartToken)
}

func (t TVoid) BaseType() Type { return t }

func (t TVoid) Const() (value interface{}, ok bool) {
	return nil, false
}

func (t TVoid) DefaultValue() interface{} {
	panic("can't take default value of void type")
}

type TInt struct {
	StartToken antlr.Token
	Value      *int
}

func (TInt) String() string {
	return "int"
}

func (t TInt) Position() string {
	return PosFromToken(t.StartToken)
}

func (t TInt) BaseType() Type { return t }

func (t TInt) Const() (value interface{}, ok bool) {
	if t.Value == nil {
		return nil, false
	}

	return *t.Value, true
}

func (t TInt) DefaultValue() interface{} {
	return 0
}

type TString struct {
	StartToken antlr.Token
	Value      *string
}

func (TString) String() string {
	return "string"
}

func (t TString) Position() string {
	return PosFromToken(t.StartToken)
}

func (t TString) BaseType() Type { return t }

func (t TString) Const() (value interface{}, ok bool) {
	if t.Value == nil {
		return nil, false
	}

	return *t.Value, true
}

func (t TString) DefaultValue() interface{} {
	return ""
}

type TBool struct {
	StartToken antlr.Token
	Value      *bool
}

func (TBool) String() string {
	return "boolean"
}

func (t TBool) Position() string {
	return PosFromToken(t.StartToken)
}

func (t TBool) BaseType() Type { return t }

func (t TBool) Const() (value interface{}, ok bool) {
	if t.Value == nil {
		return nil, false
	}

	return *t.Value, true
}

func (t TBool) DefaultValue() interface{} {
	return false
}

type FieldInfo struct {
	Type Type

	// For methods: index in vtable.
	// For fields: offset from struct beginning.
	Offset int
}

type TClass struct {
	ID              antlr.TerminalNode
	Fields          map[string]FieldInfo
	TotalNonMethods int
	TotalMethods    int
	Parent          *TClassRef
}

func (t TClass) String() string {
	return t.ID.GetText()
}

func (t TClass) Position() string {
	return PosFromToken(t.ID.GetSymbol())
}

func (t TClass) Print() {
	fmt.Println("class", t)
	for name, value := range t.Fields {
		fmt.Println("\t", name, value)
	}
}

func (t TClass) AsRef() TClassRef {
	return TClassRef{
		ID: t.ID,
	}
}

func (t TClass) BaseType() Type { return t }

func (t TClass) Const() (value interface{}, ok bool) {
	return nil, false
}

func (t TClass) DefaultValue() interface{} {
	return 0
}

type TClassRef struct {
	ID antlr.TerminalNode
}

func (t TClassRef) String() string {
	return t.ID.GetText()
}

func (t TClassRef) Position() string {
	return PosFromToken(t.ID.GetSymbol())
}

func (t TClassRef) BaseType() Type { return t }

func (t TClassRef) Const() (value interface{}, ok bool) {
	return nil, false
}

func (t TClassRef) DefaultValue() interface{} {
	panic("can't take default value of a classref")
}

type TArray struct {
	StartToken antlr.Token
	Elem       Type
}

func (t TArray) String() string {
	return t.Elem.String() + "[]"
}

func (t TArray) Position() string {
	return PosFromToken(t.StartToken)
}

func (t TArray) BaseType() Type { return t.Elem.BaseType() }

func (t TArray) Const() (value interface{}, ok bool) {
	return nil, false
}

func (t TArray) DefaultValue() interface{} {
	return 0
}

type FArg struct {
	Ident string
	Type  Type
}

type TFun struct {
	Ident    string
	Terminal antlr.TerminalNode
	Args     []FArg
	Result   Type
}

func (t TFun) String() string {
	var res = t.Result.String() + " " + t.Ident + "("
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
	if t.Terminal == nil {
		return "unknown position"
	}

	return PosFromToken(t.Terminal.GetSymbol())
}

func (t TFun) BaseType() Type { return t }

func (t TFun) Const() (value interface{}, ok bool) {
	return nil, false
}

func (t TFun) DefaultValue() interface{} {
	panic("can't take default value of a function")
}

type TReadOnly struct {
	Type Type
}

func (t TReadOnly) String() string                      { return fmt.Sprintf("readonly(%s)", t.Type) }
func (t TReadOnly) Position() string                    { return t.Type.Position() }
func (t TReadOnly) BaseType() Type                      { return t.Type }
func (t TReadOnly) Const() (value interface{}, ok bool) { return nil, false }
func (t TReadOnly) DefaultValue() interface{}           { panic("can't take default value of a readonly") }
