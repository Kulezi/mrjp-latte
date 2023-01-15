package ir

import (
	"fmt"
	"latte/compiler/config"
	"latte/compiler/frontend"
	"latte/compiler/frontend/typecheck"
	"latte/compiler/frontend/types"
)

type Edges struct {
	Succ map[Label][]Label
	Pred map[Label][]Label
}

func (e *Edges) Add(l1, l2 Label) {
	e.Succ[l1] = append(e.Succ[l1], l2)
	e.Pred[l2] = append(e.Pred[l2], l1)
}

type ControlFlowGraph struct {
	Nodes    []BasicBlock
	Edges    Edges
	BlockIdx map[Label]int
}

func (cfg ControlFlowGraph) String() string {
	var res string
	for _, block := range cfg.Nodes {
		res += fmt.Sprintf("%s:\n", block.Label)
		for _, op := range block.Ops {
			res += fmt.Sprintln("\t" + op.String())
		}
	}

	return res
}

type VarInfo struct {
	Signature     types.TFun
	Function      Label
	VariableCount int
	Offset        int
}

type FunInfo map[Label]VarInfo

// Generates a control flow graph that is in non-SSA form,
// along with information about variables and string literals.
func Generate(s frontend.State, config config.Config) (ControlFlowGraph, FunInfo) {
	v := MakeVisitor(typecheck.MakeVisitor(s.Signatures), config)
	v.Visit(s.Tree)

	cfg := ControlFlowGraph{
		Nodes:    v.Blocks,
		BlockIdx: make(map[Label]int),
		Edges: Edges{
			Succ: make(map[Label][]Label),
			Pred: make(map[Label][]Label),
		},
	}

	// for i, block := range cfg.Nodes {
	// 	label := block.Label
	// 	cfg.BlockIdx[label] = i
	// 	ops := block.Ops
	// 	if len(ops) > 0 {
	// 		// Getting next label won't go out of bounds,
	// 		// as the last block always ends in a return.
	// 		lastOp := ops[len(ops)-1]
	// 		switch jmp := lastOp.(type) {
	// 		case QJmp:
	// 			cfg.Edges.Add(label, jmp.Dst)
	// 		case QJnz:
	// 			cfg.Edges.Add(label, jmp.Dst)
	// 			cfg.Edges.Add(label, cfg.Nodes[i+1].Label)
	// 		case QJz:
	// 			cfg.Edges.Add(label, jmp.Dst)
	// 			cfg.Edges.Add(label, cfg.Nodes[i+1].Label)
	// 		case QRet, QVRet:
	// 		default:
	// 			cfg.Edges.Add(label, cfg.Nodes[i+1].Label)
	// 		}
	// 	} else {
	// 		cfg.Edges.Add(label, cfg.Nodes[i+1].Label)
	// 	}

	// }
	// var nonDead []BasicBlock
	// for _, block := range cfg.Nodes {
	// 	if _, ok := cfg.Edges.Pred[block.Label]; ok {
	// 		nonDead = append(nonDead, block)
	// 	}
	// }

	// cfg.Nodes = nonDead

	return cfg, v.FunInfo
}
