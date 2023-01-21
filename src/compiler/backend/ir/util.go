package ir

import (
	"fmt"
	. "latte/compiler/frontend/types"
	"os"
)

func Unimplemented(format string, args ...interface{}) {
	format = "ERROR\n" + format
	fmt.Fprintf(os.Stderr, format, args...)
	panic("unimplemented")
	os.Exit(1)
}

var reservedFunctions = map[string]struct{}{
	"main":        {},
	"printInt":    {},
	"printString": {},
	"readInt":     {},
	"readString":  {},
	"error":       {},
}

func reservedFunction(ident string) bool {
	_, ok := reservedFunctions[ident]
	return ok
}

func (v *Visitor) GetFunctionLabel(ident string) Label {
	var fname Fname
	fname.Name = ident
	if v.CurClass != nil {
		fname.Class = (*v.CurClass).ID.GetText()
	}

	if label, ok := v.functionLabels[fname]; ok {
		return label
	}

	if v.CurClass == nil && reservedFunction(ident) {
		v.functionLabels[fname] = Label{IsFunction: true, Name: ident}
		return v.functionLabels[fname]
	}

	label := v.FreshLabel(fmt.Sprintf("_%s_%s", fname.Class, fname.Name))
	label.IsFunction = true
	v.functionLabels[fname] = label
	return label
}

func (v *Visitor) EnterClass(signature TClass) (exit func()) {
	v.CurClass = &signature
	v.Depth++
	// Put all fields into enviroment.
	oldLocals := v.Signatures.Locals
	v.Signatures.Locals = make(Env)
	for ident, t := range signature.Fields {
		v.ShadowLocal(ident, t.Type)
	}

	v.ShadowLocal("self", signature)

	return func() {
		v.Signatures.Locals = oldLocals
		v.Depth--
		v.CurClass = nil
	}
}

func (v *Visitor) FreshTemp(prefix string, t Type) LReg {
	tries := 0
	for {
		ident := prefix
		if tries > 0 {
			ident += fmt.Sprintf("_%d", tries)
		}

		if _, ok := v.allAddresses[ident]; !ok {
			v.allAddresses[ident] = struct{}{}
			return LReg{
				Type_: t,
				Name:  ident,
			}
		}
		tries++
	}
}

func (v *Visitor) PushLabels(lTrue, lFalse, lNext Label) (pop func()) {
	oldTrue, oldFalse, oldNext := v.lTrue, v.lFalse, v.lNext
	v.lTrue, v.lFalse, v.lNext = lTrue, lFalse, lNext
	return func() {
		v.lTrue, v.lFalse, v.lNext = oldTrue, oldFalse, oldNext
	}
}

func (v *Visitor) GetLabels() (lTrue, lFalse, lNext Label) {
	return v.lTrue, v.lFalse, v.lNext
}

func (v *Visitor) FreshLabel(prefix string) Label {
	res := fmt.Sprintf("_%s_%d", prefix, v.totalLabels)
	v.totalLabels++
	return Label{Name: res}
}

func (v *Visitor) StartBlock(l Label) {
	if v.curBlock.Label.Name != "" {
		v.CFG[v.curBlock.Label] = v.curBlock
		v.Blocks = append(v.Blocks, v.curBlock)
	}

	v.curBlock = BasicBlock{Label: l}
}

func (v *Visitor) PushField(ident string, loc LMem) (drop func()) {
	oldSignature, ok := v.Signatures.Locals[ident]
	oldLoc, ok := v.variableLocations[ident]

	v.variableLocations[ident] = loc
	v.Signatures.Locals[ident] = TypeInfo{
		Type: loc.Type_,
	}

	return func() {
		if ok {
			v.Signatures.Locals[ident] = oldSignature
			v.variableLocations[ident] = oldLoc
		}
	}
}

func (v *Visitor) ShadowLocal(ident string, t Type) (location Location, drop func()) {
	loc := v.FreshTemp(ident, t)
	loc.Variable = ident
	loc.Index = v.varCount
	v.varCount++

	oldSignature, ok := v.Signatures.Locals[ident]
	oldLoc, ok := v.variableLocations[ident]

	v.variableLocations[ident] = loc
	v.Signatures.Locals[ident] = TypeInfo{
		Type: t,
	}

	return loc, func() {
		if ok {
			v.Signatures.Locals[ident] = oldSignature
			v.variableLocations[ident] = oldLoc
		}
	}
}

func (v *Visitor) GetLocal(ident string) Location {
	return v.variableLocations[ident]
}

func (v *Visitor) EmitQuad(q Quadruple) {
	v.lastQuad = q
	v.curBlock.Ops = append(v.curBlock.Ops, q)
	if q.IsJump() {
		if v.curBlock.Label.Name == "" {
			v.curBlock.Label = v.FreshLabel("after_jump")
		}
		v.CFG[v.curBlock.Label] = v.curBlock
		v.Blocks = append(v.Blocks, v.curBlock)
		v.curBlock = BasicBlock{}
	}
}

func (v *Visitor) TypeOfLocal(ident string) (TypeInfo, bool) {
	res, ok := v.Signatures.Locals[ident]
	return res, ok
}

func (v *Visitor) TypeOfGlobal(ident string) (TypeInfo, bool) {
	res, ok := v.Signatures.Globals[ident]
	return res, ok
}

func (v *Visitor) TypeOf(ident string) (TypeInfo, bool) {
	if res, ok := v.TypeOfLocal(ident); ok {
		return res, ok
	}

	return v.TypeOfGlobal(ident)
}
