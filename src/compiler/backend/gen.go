package backend

import (
	"fmt"
	"latte/compiler/backend/ir"
	. "latte/compiler/config"
	"latte/compiler/frontend"
	"latte/compiler/frontend/types"
)

const (
	syscallExit = 60
	QWORD       = 8

	rax = `%rax`
	rbx = `%rbx`
	rcx = `%rcx`
	rdx = `%rdx`
	rsi = `%rsi`
	rdi = `%rdi`
	rsp = `%rsp`
	rbp = `%rbp`
	r8  = `%r8`
	r9  = `%r9`
	r10 = `%r10`
	r11 = `%r11`
	r12 = `%r12`
	r13 = `%r13`
	r14 = `%r14`
	r15 = `%r15`
)

var scratchRegisters = []string{rax, rdi, rsi, rdx, rcx, r8, r9, r10, r11}
var preservedRegisters = []string{rbx, rsp, rbp, r12, r13, r14, r15}

func GenX64(s frontend.State, config Config) string {
	cfg, finfo := ir.Generate(s, config)
	if config.PrintIR {
		fmt.Println(cfg)
	}

	x64 := X64Generator{FunInfo: finfo, stringAdressses: make(map[string]string)}

	x64.Emit(".data\n")

	for _, block := range cfg.Nodes {
		x64.GenDataFromBlock(block)
	}

	x64.Emit(".text\n.globl main\n")

	for _, block := range cfg.Nodes {
		x64.GenFromBlock(block)
	}

	return x64.res
}

type X64Generator struct {
	FunInfo         ir.FunInfo
	curFun          ir.VarInfo
	res             string
	alignLabelCnt   int
	stringAdressses map[string]string
}

func (x64 *X64Generator) GenDataFromBlock(block ir.BasicBlock) {
	for _, q := range block.Ops {
		switch q := q.(type) {
		case ir.QBinOp:
			x64.EmitDataForLoc(q.Lhs)
			x64.EmitDataForLoc(q.Rhs)
		case ir.QCall:
			for _, arg := range q.Args {
				x64.EmitDataForLoc(arg)
			}
		case ir.QMov:
			x64.EmitDataForLoc(q.Src)
		case ir.QPush:
			x64.EmitDataForLoc(q.Src)
		case ir.QRelOp:
			x64.EmitDataForLoc(q.Lhs)
			x64.EmitDataForLoc(q.Rhs)
		case ir.QRet:
			x64.EmitDataForLoc(q.Value)
		}
	}
}

func (x64 *X64Generator) EmitDataForLoc(loc ir.Location) {
	if loc, ok := loc.(ir.LConst); ok {
		if s, ok := loc.Value.(string); ok {
			if _, ok := x64.stringAdressses[s]; !ok {
				label := fmt.Sprintf("_str_%d", len(x64.stringAdressses))
				x64.stringAdressses[s] = label
				x64.EmitLabel(label)
				for _, v := range s {
					x64.EmitOp(".byte 0x%x", v)
				}
				x64.EmitOp(".byte %x", 0)
			}
		}
	}
}

func (x64 *X64Generator) EmitLabel(s string) {
	x64.res += s + ":\n"
}

func (x64 *X64Generator) Emit(format string, args ...interface{}) {
	x64.res += fmt.Sprintf(format, args...)
}

func (x64 *X64Generator) EmitOp(format string, args ...interface{}) {
	x64.Emit("\t"+format+"\n", args...)
}

func (x64 *X64Generator) EmitFunctionProlog() {
	x64.EmitOp("pushq %s", rbp)
	x64.EmitOp("movq %s, %s", rsp, rbp)
	varCnt := x64.curFun.VariableCount
	if varCnt > 0 {
		x64.EmitOp("subq $%#x, %s", varCnt*QWORD, rsp)
	}

	for _, reg := range preservedRegisters {
		if reg != rbp && reg != rsp {
			x64.EmitOp("pushq %s", reg)
		}
	}

	// The stack is now [args], return address, oldRbp.
	offset := (len(x64.curFun.Signature.Args) + 1)
	for i := range x64.curFun.Signature.Args {
		x64.EmitOp("movq %#x(%s), %s", QWORD*(offset-i), rbp, rax)
		x64.EmitOp("movq %s, %#x(%s)", rax, -QWORD*(i+1), rbp)
	}
}

func (x64 *X64Generator) EmitFunctionEpilog() {
	for i := len(preservedRegisters) - 1; i >= 0; i-- {
		reg := preservedRegisters[i]
		if reg != rbp && reg != rsp {
			x64.EmitOp("popq %s", reg)
		}
	}

	varCnt := x64.curFun.VariableCount
	if varCnt > 0 {
		x64.EmitOp("addq $%#x, %s", varCnt*QWORD, rsp)
	}

	x64.EmitOp("movq %s, %s", rbp, rsp)
	x64.EmitOp("popq %s", rbp)
}

func (x64 *X64Generator) GenFromBlock(block ir.BasicBlock) {
	label := block.Label
	x64.EmitLabel(label.Name)

	// Prepare stack frame.
	if label.IsFunction {
		x64.curFun = x64.FunInfo[label]
		x64.EmitFunctionProlog()
	}

	for _, op := range block.Ops {
		x64.GenFromQuad(op)
	}
}

func (x64 *X64Generator) GenFromQuad(q ir.Quadruple) {
	switch q := q.(type) {
	case ir.QMov:
		x64.EmitQMov(q)
	case ir.QBinOp:
		x64.EmitBinOp(q)
	case ir.QRelOp:
		x64.EmitRelOp(q)
	case ir.QNeg:
		x64.EmitNeg(q)
	case ir.QJmp:
		x64.EmitJmp(q)
	case ir.QJz:
		x64.EmitJz(q)
	case ir.QJnz:
		x64.EmitJnz(q)
	case ir.QRet:
		x64.EmitRet(q)
	case ir.QVRet:
		x64.EmitVRet(q)
	case ir.QCall:
		x64.EmitCall(q)
	case ir.QPush:
		x64.EmitPush(q)
	case ir.QArrayAccess:
		x64.EmitArrayAccess(q)
	case ir.QDeref:
		x64.EmitDeref(q)
	case ir.QNewArray:
		x64.EmitNewArray(q)
	default:
		panic("unsupported quadruple")
	}
}

func (x64 *X64Generator) getLoc(loc ir.Location) string {
	reg := loc.(ir.LReg)
	if reg.Variable == "" {
		return ""
	}
	varOffset := reg.Index
	funOffset := x64.curFun.Offset
	offset := (varOffset - funOffset + 1) * QWORD
	return fmt.Sprintf("%#x(%s)", -offset, rbp)
}

func (x64 *X64Generator) EmitLoad(register string, loc ir.Location) {
	// Void doesn't need to be stored anywhere, and won't be popped from stack.
	if _, ok := loc.Type().(types.TVoid); ok {
		return
	}

	switch loc := loc.(type) {
	case ir.LConst:
		switch v := loc.Value.(type) {
		case string:
			x64.EmitOp("movq $%s, %s", x64.stringAdressses[v], register)
		case bool:
			// Value 0 represents false, 1 represents true.
			x := 0
			if v {
				x = 1
			}
			x64.EmitOp("movq $%#x, %s", x, register)
		case int:
			x64.EmitOp("movq $%d, %s", v, register)
		}
	case ir.LReg:
		// Temporaries go on stack, so we need to pop
		if loc.Variable == "" {
			x64.EmitOp("popq %s", register)
		} else {
			x64.EmitOp("movq %s, %s", x64.getLoc(loc), register)
		}
	}
}

func (x64 *X64Generator) EmitQMov(q ir.QMov) {
	x64.EmitLoad(rax, q.Src)
	// In case of a standalone expression we can forget the result.
	if _, ok := q.Dst.(ir.LDrop); ok {
		return
	}
	dst := x64.getLoc(q.Dst)
	if dst == "" {
		x64.EmitOp("pushq %s", rax)
	} else {
		x64.EmitOp("movq %s, %s", rax, x64.getLoc(q.Dst))
	}
}

func (x64 *X64Generator) EmitPush(q ir.QPush) {
	x64.EmitLoad(rax, q.Src)
	x64.EmitOp("pushq %s", rax)
}

func (x64 *X64Generator) EmitStringAdd(q ir.QBinOp) {
	x64.EmitLoad(rsi, q.Rhs)
	x64.EmitLoad(rdi, q.Lhs)
	x64.EmitOp("call concat")
	if dst, ok := q.Dst.(ir.LReg); ok && dst.Variable != "" {
		x64.EmitOp("movq %s, %s", rax, x64.getLoc(dst))
	} else {
		x64.EmitOp("pushq %s", rax)
	}
}

func (x64 *X64Generator) EmitBinOp(q ir.QBinOp) {
	if _, ok := q.Lhs.Type().(types.TString); ok {
		x64.EmitStringAdd(q)
		return
	}

	x64.EmitLoad(rbx, q.Rhs)
	x64.EmitLoad(rax, q.Lhs)
	switch q.Op {
	case "+":
		x64.EmitOp("addq %s, %s", rbx, rax)
	case "-":
		x64.EmitOp("subq %s, %s", rbx, rax)
	case "*":
		x64.EmitOp("imulq %s", rbx)
	case "/":
		x64.EmitOp("cqto")
		x64.EmitOp("idivq %s", rbx)
	case "%":
		x64.EmitOp("cqto")
		x64.EmitOp("idivq %s", rbx)
		x64.EmitOp("xchg %s, %s", rdx, rax)
	default:
		panic("unsupported binary operator")
	}

	if dst, ok := q.Dst.(ir.LReg); ok && dst.Variable != "" {
		x64.EmitOp("movq %s, %s", rax, x64.getLoc(dst))
	} else {
		x64.EmitOp("pushq %s", rax)
	}
}

func (x64 *X64Generator) EmitStringRelOp(q ir.QRelOp) {
	x64.EmitLoad(rsi, q.Rhs)
	x64.EmitLoad(rdi, q.Lhs)
	x64.EmitOp("call compare")
	x64.EmitOp("cmp %s, %s", rax, rax)
	var op string
	switch q.Op {
	case "==":
		op = "je"
	case "!=":
		op = "jne"
	default:
		panic("unsupported string operation")
	}

	if q.LNext == q.LTrue {
		x64.EmitOp("%s %s", inverseJmp[op], q.LFalse)
	} else if q.LNext == q.LFalse {
		x64.EmitOp("%s %s", op, q.LTrue)
	} else {
		x64.EmitOp("%s %s", op, q.LTrue)
		x64.EmitOp("jmp %s", op, q.LFalse)
	}
}

var inverseJmp = map[string]string{
	"jge": "jl",
	"jg":  "jle",
	"je":  "jne",
	"jne": "je",
	"jl":  "jge",
	"jle": "jg",
}

func (x64 *X64Generator) EmitRelOp(q ir.QRelOp) {
	if _, ok := q.Lhs.Type().(types.TString); ok {
		x64.EmitStringRelOp(q)
		return
	}

	x64.EmitLoad(rbx, q.Rhs)
	x64.EmitLoad(rax, q.Lhs)
	x64.EmitOp("cmp %s, %s", rbx, rax)

	var op string
	switch q.Op {
	case ">":
		op = "jg"
	case ">=":
		op = "jge"
	case "<=":
		op = "jle"
	case "<":
		op = "jl"
	case "==":
		op = "je"
	case "!=":
		op = "jne"
	default:
		panic("unsupported binary operator")
	}

	if q.LNext == q.LTrue {
		x64.EmitOp("%s %s", inverseJmp[op], q.LFalse)
	} else if q.LNext == q.LFalse {
		x64.EmitOp("%s %s", op, q.LTrue)
	} else {
		x64.EmitOp("%s %s", op, q.LTrue)
		x64.EmitOp("jmp %s", op, q.LFalse)
	}
}

func (x64 *X64Generator) EmitNeg(q ir.QNeg) {
	x64.EmitLoad(rax, q.Arg)
	x64.EmitOp("neg %s", rax)
	x64.EmitOp("pushq %s", rax)
}

func (x64 *X64Generator) EmitJmp(q ir.QJmp) {
	x64.EmitOp("jmp %s", q.Dst)
}

func (x64 *X64Generator) EmitJz(q ir.QJz) {
	x64.EmitLoad(rax, q.Value)
	x64.EmitOp("test %s, %s", rax, rax)
	x64.EmitOp("jz %s", q.Dst)
}

func (x64 *X64Generator) EmitJnz(q ir.QJnz) {
	x64.EmitLoad(rax, q.Value)
	x64.EmitOp("test %s, %s", rax, rax)
	x64.EmitOp("jnz %s", q.Dst)
}

func (x64 *X64Generator) EmitRet(q ir.QRet) {
	if x64.curFun.Function.Name == "main" {
		// Put return value to rdi, it's a scratch register so epilog won't change it.
		x64.EmitLoad(rdi, q.Value)
		x64.EmitFunctionEpilog()
		x64.EmitLoad(rax, ir.LConst{Type_: types.TInt{}, Value: syscallExit})
		x64.EmitOp("syscall")
	} else {
		x64.EmitLoad(rax, q.Value)
		x64.EmitFunctionEpilog()
		x64.EmitOp("ret")
	}
}

func (x64 *X64Generator) EmitVRet(q ir.QVRet) {
	x64.EmitFunctionEpilog()
	x64.EmitOp("ret")
}

func (x64 *X64Generator) EmitCall(q ir.QCall) {
	// Calling runtime functions with arguments needs passing them through rdi register.
	if q.Label.Name == "printInt" || q.Label.Name == "printString" || q.Label.Name == "newArray" {
		x64.EmitOp("pop %s", rdi)
	}

	// If we are calling a foreign function we need to before calling.
	if _, ok := foreignFunctions[q.Label.Name]; ok {
		// Check stack alignment.
		x64.EmitOp("testq %s, %s", "$0x000000000000000F", rsp)
		lAlign := fmt.Sprintf("_align_%d", x64.alignLabelCnt)
		x64.alignLabelCnt++
		lEnd := fmt.Sprintf("_alignEnd_%d", x64.alignLabelCnt)
		x64.alignLabelCnt++

		x64.EmitOp("jnz %s", lAlign)

		// If alignment is not needed.
		x64.EmitOp("call %s", q.Label.Name)
		x64.EmitOp("jmp %s", lEnd)

		// If alignment is needed.
		x64.EmitLabel(lAlign)
		x64.EmitOp("subq $%#x, %s", QWORD, rsp)
		x64.EmitOp("call %s", q.Label.Name)

		x64.EmitOp("addq $%#x, %s", QWORD, rsp)
		x64.EmitLabel(lEnd)
	} else {
		x64.EmitOp("call %s", q.Label.Name)
		// Pop arguments from the stack.
		x64.EmitOp("addq $%#x, %s", len(q.Args)*QWORD, rsp)
	}

	if _, ok := q.Signature.Result.(types.TVoid); !ok {
		x64.EmitOp("pushq %s", rax)
	}
}

var foreignFunctions = map[string]struct{}{
	"newArray":    {},
	"printInt":    {},
	"readInt":     {},
	"error":       {},
	"printString": {},
	"readString":  {},
}

func (x64 *X64Generator) EmitArrayAccess(q ir.QArrayAccess) {
	x64.EmitLoad(rax, q.Array)
	x64.EmitLoad(rbx, q.Index)

	// First field of an array stores its length.
	displacement := QWORD
	x64.EmitOp("leaq %d(%s, %s, %d), %s", displacement, rax, rbx, QWORD, rax)
	x64.EmitOp("push %s", rax)
}

func (x64 *X64Generator) EmitDeref(q ir.QDeref) {
	x64.EmitLoad(rax, q.Src)
	x64.EmitOp("push (%s)", rax)
}

func (x64 *X64Generator) EmitNewArray(q ir.QNewArray) {
	x64.EmitCall(ir.QCall{
		Signature: newArraySignature,
		Label:     ir.Label{IsFunction: true, Name: "newArray"},
		Dst:       q.Dst,
		Args:      []ir.Location{q.Size},
	})
}

var newArraySignature = types.TFun{
	Ident:  "newArray",
	Args:   []types.FArg{{Ident: "size", Type: types.TInt{}}},
	Result: types.TInt{},
}
