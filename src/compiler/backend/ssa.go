package backend

// type SSABasicBlock struct {
// 	Label string
// 	Ops   []QuadInfo

// 	in  AliveSet
// 	out AliveSet
// }

// // Enriches a basic block with liveness info.
// func EnrichBlock(block ir.BasicBlock) SSABasicBlock {
// 	res := SSABasicBlock{
// 		Label: block.Label,
// 		Ops:   make([]QuadInfo, len(block.Ops)),
// 	}

// 	for i := len(res.Ops) - 1; i >= 0; i-- {
// 		if i+1 < len(res.Ops) {
// 			res.Ops[i].out = res.Ops[i+1].in.Clone()
// 		} else {
// 			res.Ops[i].out = set.NewSet[string]()
// 		}

// 		res.Ops[i].in = res.Ops[i].out.Difference(res.Ops[i].GetKill()).Union(res.Ops[i].GetUsed())
// 	}

// 	return res
// }

// type SSAControlFlowGraph struct {
// 	Nodes    []SSABasicBlock
// 	Succ     map[ir.Label][]ir.Label
// 	BlockIdx map[ir.Label]int
// }

// func MakeSSA(cfg ir.ControlFlowGraph) SSAControlFlowGraph {
// 	res := SSAControlFlowGraph{
// 		Succ:     cfg.Succ,
// 		BlockIdx: cfg.BlockIdx,
// 	}

// 	// Enrich blocks with local liveness info.
// 	for _, block := range cfg.Nodes {
// 		res.Nodes = append(res.Nodes, EnrichBlock(block))
// 		res.in
// 	}
// 	// Enrich blocks with global liveness info.
// 	for {
// 		for _, block := range cfg.Nodes {

// 		}
// 	}

// 	panic("unimplemented")
// }
