package ir

import "fmt"

type Const interface{}
type Location interface {
	String() string
}

type LConst struct {
	Value Const
}

func (v LConst) String() string {
	return fmt.Sprintf("%d", v.Value)
}

type LAddr uint

func (addr LAddr) String() string {
	return fmt.Sprintf("x_%d", addr)
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

func (q QMov) String() string { return fmt.Sprintf("%s := %s", q.Dst, q.Src) }

type QRet struct {
	Value Location
}

func (q QRet) String() string { return fmt.Sprintf("return %s", q.Value) }

func (q QRet) IsJump() bool { return true }
