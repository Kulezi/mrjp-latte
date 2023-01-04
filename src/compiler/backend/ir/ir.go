package ir

type Location interface{}

type LReg struct{}

type LMem struct{}

type LIntLit struct{}

type LStrLit struct{}

type BasicBlock struct {
	label string
	ops   []Quadruple
}

type Quadruple interface{}

type AddInt struct {
	destination, src, dst Location
}
