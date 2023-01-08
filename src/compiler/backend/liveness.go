package backend

import (
	"latte/compiler/backend/ir"

	set "github.com/deckarep/golang-set/v2"
)

type AliveSet set.Set[string]

func AliveSetFromLocations(locs ...ir.Location) AliveSet {
	res := set.NewSet[string]()
	for _, loc := range locs {
		if name, ok := registerName(loc); ok {
			res.Add(name)
		}
	}

	return res
}

func registerName(loc ir.Location) (name string, ok bool) {
	if loc, ok := loc.(ir.LReg); ok {
		return loc.String(), true
	}

	return "", false
}

// Quadruple extended with info about liveness.
type QuadInfo struct {
	ir.Quadruple

	// // Variables alive before executing this quadruple.
	// in AliveSet
	// // Variables alive after executing this quadruple.
	// out AliveSet
}

// Returns a set of variables used by this quadruple.
func (q *QuadInfo) GetUsed() AliveSet {
	switch q := q.Quadruple.(type) {
	case ir.QUnOp:
		return AliveSetFromLocations(q.Arg)
	case ir.QBinOp:
		return AliveSetFromLocations(q.Lhs, q.Rhs)
	case ir.QCall:
		return AliveSetFromLocations(q.Args...)
	case ir.QJnz:
		return AliveSetFromLocations(q.Value)
	case ir.QJz:
		return AliveSetFromLocations(q.Value)
	case ir.QMov:
		return AliveSetFromLocations(q.Src)
	case ir.QRet:
		return AliveSetFromLocations(q.Value)
	default:
		return AliveSetFromLocations()
	}
}

// Returns a set of variables killed by this quadruple.
func (q *QuadInfo) GetKill() AliveSet {
	switch q := q.Quadruple.(type) {
	case ir.QUnOp:
		return AliveSetFromLocations(q.Dst)
	case ir.QBinOp:
		return AliveSetFromLocations(q.Dst)
	case ir.QCall:
		return AliveSetFromLocations(q.Dst)
	case ir.QMov:
		return AliveSetFromLocations(q.Dst)
	case ir.QPop:
		return AliveSetFromLocations(q.Dst)
	default:
		return AliveSetFromLocations()
	}
}
