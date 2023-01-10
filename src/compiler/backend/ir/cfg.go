package ir

import (
	"latte/compiler/config"
	"latte/compiler/frontend"
	"latte/compiler/frontend/typecheck"
)

type ControlFlowGraph struct {
	Nodes    []BasicBlock
	Succ     map[Label][]Label
	BlockIdx map[Label]int
}

type VarInfo struct {
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
		Succ:     make(map[Label][]Label),
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
				cfg.Succ[label] = append(cfg.Succ[label], jmp.Dst)
			case QJnz:
				cfg.Succ[label] = append(cfg.Succ[label], jmp.Dst)
				cfg.Succ[label] = append(cfg.Succ[label], cfg.Nodes[i+1].Label)
			case QJz:
				cfg.Succ[label] = append(cfg.Succ[label], jmp.Dst)
				cfg.Succ[label] = append(cfg.Succ[label], cfg.Nodes[i+1].Label)
			case QRet, QVRet:
			default:
				cfg.Succ[label] = append(cfg.Succ[label], cfg.Nodes[i+1].Label)
			}
		} else {
			// log.Println("empty label wtf")
			cfg.Succ[label] = append(cfg.Succ[label], cfg.Nodes[i+1].Label)
		}
	}

	return cfg, v.FunInfo
}

func (cfg *ControlFlowGraph) String() string {
	ir := ""
	for _, block := range cfg.Nodes {
		ir += block.Label.Name + "\n"
		for _, op := range block.Ops {
			ir += "\t" + op.String() + "\n"
		}
	}

	return ir
}
