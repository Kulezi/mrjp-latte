package backend

import "latte/compiler/backend/ir"

type SSAControlFlowGraph struct {
	*ir.ControlFlowGraph
}

func MakeSSA(cfg ir.ControlFlowGraph) SSAControlFlowGraph {
	panic("unimplemented")
}
