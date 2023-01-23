package ir

import (
	"fmt"
	. "latte/compiler/frontend/types"
)

type Const interface{}
type Location interface {
	String() string
	Type() Type
}

type LSelfField struct {
	Type_  Type
	Name   string
	Offset int
}

func (v LSelfField) String() string { return fmt.Sprintf("self.%s", v.Name) }
func (v LSelfField) Type() Type     { return v.Type_ }

type LMem struct {
	Type_ Type
	Addr  Location
}

func (v LMem) String() string { return fmt.Sprintf("*%s", v.Addr) }
func (v LMem) Type() Type     { return v.Type_ }

type LDrop struct{ Type_ Type }

func (v LDrop) String() string { return "LDrop" }
func (v LDrop) Type() Type     { return v.Type_ }

type LConst struct {
	Type_ Type
	Value Const
}

func (v LConst) String() string {
	if v, ok := v.Value.(string); ok {
		return fmt.Sprintf("\"%s\"", v)
	}
	return fmt.Sprintf("%v", v.Value)
}

func (v LConst) Type() Type {
	return v.Type_
}

type LReg struct {
	Type_ Type
	Name  string

	// Values below are non-zero only for a non-temporary variable
	Variable string
	Index    int
}

func (v LReg) String() string {
	return v.Name
}

func (v LReg) Type() Type {
	return v.Type_
}

type LUnassigned struct {
	Type_ Type
}

func (v LUnassigned) String() string {
	return "unassigned"
}

func (v LUnassigned) Type() Type {
	return v.Type_
}

type BasicBlock struct {
	Label Label
	Ops   []Quadruple
}

type Quadruple interface {
	IsJump() bool
	String() string
}

type QBase struct{}

func (QBase) IsJump() bool   { return false }
func (QBase) String() string { return "placeholder operation" }

type QMov struct {
	QBase
	Src, Dst Location
}

func (q QMov) String() string { return fmt.Sprintf("%s = %s", q.Dst, q.Src) }

type QVRet struct{}

func (q QVRet) String() string { return "return" }
func (q QVRet) IsJump() bool   { return true }

type QRet struct {
	Value Location
}

func (q QRet) String() string { return fmt.Sprintf("return %s", q.Value) }
func (q QRet) IsJump() bool   { return true }

type QNeg struct {
	QBase
	Dst Location
	Arg Location
}

func (q QNeg) String() string { return fmt.Sprintf("%s = -%s", q.Dst, q.Arg) }

type QBinOp struct {
	QBase
	Op       string
	Dst      Location
	Lhs, Rhs Location
}

func (q QBinOp) String() string { return fmt.Sprintf("%s = %s %s %s", q.Dst, q.Lhs, q.Op, q.Rhs) }

type QRelOp struct {
	QBase
	Op                   string
	LFalse, LTrue, LNext Label
	Lhs, Rhs             Location
}

func (q QRelOp) String() string {
	return fmt.Sprintf("if %s %s %s goto %s else goto %s", q.Lhs, q.Op, q.Rhs, q.LTrue, q.LFalse)
}

type QJmp struct {
	Dst Label
}

func (q QJmp) String() string { return "goto " + q.Dst.Name }
func (q QJmp) IsJump() bool   { return true }

type QJz struct {
	Value Location
	Dst   Label
}

func (q QJz) String() string { return fmt.Sprintf("if !%s goto %s", q.Value, q.Dst) }
func (q QJz) IsJump() bool   { return true }

type QJnz struct {
	Value Location
	Dst   Label
}

func (q QJnz) String() string { return fmt.Sprintf("if %s goto %s", q.Value, q.Dst) }
func (q QJnz) IsJump() bool   { return true }

type QCall struct {
	QBase
	Signature TFun
	Label     Label
	Dst       Location
	Args      []Location
}

func (q QCall) String() string { return fmt.Sprintf("%s = call %s(%s)", q.Dst, q.Label, q.Args) }

type QCallMethod struct {
	QBase
	Signature TFun
	Label     Location
	Dst       Location
	Args      []Location
}

func (q QCallMethod) String() string { return fmt.Sprintf("%s = call %s(%s)", q.Dst, q.Label, q.Args) }

type QPush struct {
	QBase
	Src Location
}

func (q QPush) String() string { return fmt.Sprintf("push %s", q.Src) }

type QArrayAccess struct {
	QBase
	Array Location
	Index Location
	Dst   Location
}

func (q QArrayAccess) String() string { return fmt.Sprintf("%s = %s[%s]", q.Dst, q.Array, q.Index) }

type QArrayDeref struct {
	QBase
	Array Location
	Index Location
	Dst   Location
}

func (q QArrayDeref) String() string { return fmt.Sprintf("%s = *%s[%s]", q.Dst, q.Array, q.Index) }

type QDeref struct {
	QBase
	Src Location
	Dst Location
}

func (q QDeref) String() string { return fmt.Sprintf("%s = *%s", q.Dst, q.Src) }

type QNewArray struct {
	QBase
	Type Type
	Size Location
	Dst  Location
}

func (q QNewArray) String() string { return fmt.Sprintf("%s = new %s[%s]", q.Dst, q.Type, q.Size) }

type QNewClass struct {
	QBase
	Class TClass
	Dst   Location
}

func (q QNewClass) String() string { return fmt.Sprintf("%s = new %s", q.Dst, q.Class) }

type QDup struct{ QBase }

func (q QDup) String() string { return "dup" }
