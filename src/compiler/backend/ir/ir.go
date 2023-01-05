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
	Addr  uint
}

func (v LReg) String() string {
	return fmt.Sprintf("r_%d", v.Addr)
}

func (v LReg) Type() Type {
	return v.Type_
}

type BasicBlock struct {
	Label string
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

type QUnOp struct {
	QBase
	Op  string
	Dst Location
	Arg Location
}

func (q QUnOp) String() string { return fmt.Sprintf("%s = %s%s", q.Dst, q.Op, q.Arg) }

type QBinOp struct {
	QBase
	Op       string
	Dst      Location
	Lhs, Rhs Location
}

func (q QBinOp) String() string { return fmt.Sprintf("%s = %s%s%s", q.Dst, q.Lhs, q.Op, q.Rhs) }
