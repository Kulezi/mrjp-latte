package ir

import (
	"latte/compiler/config"
	"latte/compiler/frontend"
	"latte/compiler/frontend/typecheck"
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
