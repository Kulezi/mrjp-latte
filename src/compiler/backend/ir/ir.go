package ir

import (
	"fmt"
	"latte/compiler/config"
	"latte/compiler/frontend"
	"latte/compiler/frontend/typecheck"
	. "latte/compiler/frontend/types"
	"log"
)

func Generate(s frontend.State, config config.Config) ControlFlowGraph {
	v := MakeVisitor(typecheck.MakeVisitor(s.Signatures), config)
	v.Visit(s.Tree)

	cfg := ControlFlowGraph{
		Nodes:    v.Blocks,
		BlockIdx: make(map[Label]int),
		Edges:    make(map[string][]string),
	}

	for i, block := range cfg.Nodes {
		label := block.Label
		cfg.BlockIdx[label] = i
		ops := block.Ops
		if len(ops) > 0 {
			// Getting next label won't go out of bounds,
			// as the last block always ends in a return.
			lastOp := ops[len(ops)-1]
			switch jmp := lastOp.(type) {
			case QJmp:
				cfg.Edges[label] = append(cfg.Edges[label], jmp.Dst)
			case QJnz:
				cfg.Edges[label] = append(cfg.Edges[label], jmp.Dst)
				cfg.Edges[label] = append(cfg.Edges[label], cfg.Nodes[i+1].Label)
			case QJz:
				cfg.Edges[label] = append(cfg.Edges[label], jmp.Dst)
				cfg.Edges[label] = append(cfg.Edges[label], cfg.Nodes[i+1].Label)
			case QRet, QVRet:
			default:
				cfg.Edges[label] = append(cfg.Edges[label], cfg.Nodes[i+1].Label)
			}
		} else {
			log.Println("empty label wtf")
			cfg.Edges[label] = append(cfg.Edges[label], cfg.Nodes[i+1].Label)
		}
	}

	return cfg
}

type ControlFlowGraph struct {
	Nodes    []BasicBlock
	Edges    map[Label][]Label
	BlockIdx map[Label]int
}

func (cfg *ControlFlowGraph) String() string {
	ir := ""
	for _, block := range cfg.Nodes {
		ir += block.Label + "\n"
		for _, op := range block.Ops {
			ir += "\t" + op.String() + "\n"
		}
	}

	return ir
}

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
	Type_    Type
	Variable string
	Addr     uint
}

func (v LReg) String() string {
	varName := v.Variable
	if varName == "" {
		varName = "__tmp__"
	}

	return fmt.Sprintf("r_%s_%d", v.Variable, v.Addr)
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

func (q QBinOp) String() string { return fmt.Sprintf("%s = %s %s %s", q.Dst, q.Lhs, q.Op, q.Rhs) }

type QJmp struct {
	Dst Label
}

func (q QJmp) String() string { return "goto " + q.Dst }
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
	Label Label
	Dst   Location
	Args  []Location
}

func (q QCall) String() string { return fmt.Sprintf("%s = call %s(%s)", q.Dst, q.Label, q.Args) }

type QPop struct {
	QBase
	Dst Location
}

func (q QPop) String() string { return fmt.Sprintf("%s = pop", q.Dst) }
