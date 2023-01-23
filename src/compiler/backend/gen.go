package backend

import (
	"fmt"
	"latte/compiler/backend/ir"
	. "latte/compiler/config"
	"latte/compiler/frontend"
	"latte/compiler/frontend/types"
	"runtime"
)

const (
	syscallExit = 60
	DWORD       = 4

	eax = "%eax"
	ebx = "%ebx"
	ecx = "%ecx"
	edx = "%edx"
	esi = "%esi"
	edi = "%edi"
	esp = "%esp"
	ebp = "%ebp"
)

var scratchRegisters = []string{eax, ecx, edx, esp}
var preservedRegisters = []string{ebp, esi, edi, ebx}

func GenX64(s frontend.State, config Config) string {
	cfg, finfo, vtables, functionLabels := ir.Generate(s, config)
	if config.PrintIR {
		fmt.Println(cfg)
	}

	x86 := X86Generator{FunInfo: finfo, stringAdresses: make(map[string]string), vtables: vtables}

	x86.Emit(".data\n")
	x86.EmitVTables(vtables, functionLabels)

	for _, block := range cfg.Nodes {
		x86.GenDataFromBlock(block)
	}

	x86.Emit(".text\n.globl main\n")

	for _, block := range cfg.Nodes {
		x86.GenFromBlock(block)
	}

	return x86.res
}

func (x86 *X86Generator) EmitVTables(vtables map[string]ir.VTableInfo, functionLabels map[ir.Fname]ir.Label) {
	for _, vtable := range vtables {
		x86.EmitLabel(vtable.Label.Name)
		class := vtable.Class
		fields := make([]ir.Label, class.TotalMethods)
		for _, field := range vtable.Class.Fields {
			if fun, ok := field.Type.(types.TFun); ok {
				fields[field.Offset] = functionLabels[ir.Fname{
					Class: field.Origin,
					Name:  fun.Ident,
				}]
			}
		}

		for _, label := range fields {
			x86.EmitOp(".long %s", label.Name)
		}
	}
}

type X86Generator struct {
	FunInfo        ir.FunInfo
	curFun         ir.VarInfo
	res            string
	stringAdresses map[string]string
	vtables        map[string]ir.VTableInfo
}

func (x86 *X86Generator) GenDataFromBlock(block ir.BasicBlock) {
	for _, q := range block.Ops {
		switch q := q.(type) {
		case ir.QBinOp:
			x86.EmitDataForLoc(q.Lhs)
			x86.EmitDataForLoc(q.Rhs)
		case ir.QCall:
			for _, arg := range q.Args {
				x86.EmitDataForLoc(arg)
			}
		case ir.QMov:
			x86.EmitDataForLoc(q.Src)
		case ir.QPush:
			x86.EmitDataForLoc(q.Src)
		case ir.QRelOp:
			x86.EmitDataForLoc(q.Lhs)
			x86.EmitDataForLoc(q.Rhs)
		case ir.QRet:
			x86.EmitDataForLoc(q.Value)
		}
	}
}

func (x86 *X86Generator) EmitDataForLoc(loc ir.Location) {
	if loc, ok := loc.(ir.LConst); ok {
		if s, ok := loc.Value.(string); ok {
			if _, ok := x86.stringAdresses[s]; !ok {
				label := fmt.Sprintf("_str_%d", len(x86.stringAdresses))
				x86.stringAdresses[s] = label
				x86.EmitLabel(label)
				for _, v := range s {
					x86.EmitOp(".byte 0x%x", v)
				}
				x86.EmitOp(".byte %x", 0)
			}
		}
	}
}

func (x86 *X86Generator) EmitLabel(s string) {
	x86.res += s + ":\n"
}

func (x86 *X86Generator) Emit(format string, args ...interface{}) {
	x86.res += fmt.Sprintf(format, args...)
}

func (x86 *X86Generator) EmitOp(format string, args ...interface{}) {
	x86.Emit("\t"+format+"\n", args...)
}

func (x86 *X86Generator) EmitFunctionProlog() {
	x86.EmitOp("pushl %s", ebp)
	x86.EmitOp("movl %s, %s", esp, ebp)
	varCnt := x86.curFun.VariableCount
	if varCnt > 0 {
		x86.EmitOp("subl $%#x, %s", varCnt*DWORD, esp)
	}

	for _, reg := range preservedRegisters {
		if reg != ebp && reg != esp {
			x86.EmitOp("pushl %s", reg)
		}
	}

	argCnt := len(x86.curFun.Signature.Args)
	if x86.curFun.Signature.IsMethod {
		argCnt++
	}
	// The stack is now [args], return address, oldRbp.
	offset := (argCnt + 1)

	for i := 0; i < argCnt; i++ {
		x86.EmitOp("movl %#x(%s), %s", DWORD*(offset-i), ebp, eax)
		x86.EmitOp("movl %s, %#x(%s)", eax, -DWORD*(i+1), ebp)
	}
}

func (x86 *X86Generator) EmitFunctionEpilog() {
	for i := len(preservedRegisters) - 1; i >= 0; i-- {
		reg := preservedRegisters[i]
		if reg != ebp && reg != esp {
			x86.EmitOp("popl %s", reg)
		}
	}

	varCnt := x86.curFun.VariableCount
	if varCnt > 0 {
		x86.EmitOp("addl $%#x, %s", varCnt*DWORD, esp)
	}

	x86.EmitOp("movl %s, %s", ebp, esp)
	x86.EmitOp("popl %s", ebp)
}

func (x86 *X86Generator) GenFromBlock(block ir.BasicBlock) {
	label := block.Label
	x86.EmitLabel(label.Name)

	// Prepare stack frame.
	if label.IsFunction {
		x86.curFun = x86.FunInfo[label]
		x86.EmitFunctionProlog()
	}

	for _, op := range block.Ops {
		x86.GenFromQuad(op)
	}
}

func (x86 *X86Generator) GenFromQuad(q ir.Quadruple) {
	switch q := q.(type) {
	case ir.QMov:
		x86.EmitQMov(q)
	case ir.QBinOp:
		x86.EmitBinOp(q)
	case ir.QRelOp:
		x86.EmitRelOp(q)
	case ir.QNeg:
		x86.EmitNeg(q)
	case ir.QJmp:
		x86.EmitJmp(q)
	case ir.QJz:
		x86.EmitJz(q)
	case ir.QJnz:
		x86.EmitJnz(q)
	case ir.QRet:
		x86.EmitRet(q)
	case ir.QVRet:
		x86.EmitVRet(q)
	case ir.QCall:
		x86.EmitCall(q)
	case ir.QCallMethod:
		x86.EmitCallMethod(q)
	case ir.QPush:
		x86.EmitPush(q)
	case ir.QArrayAccess:
		x86.EmitArrayAccess(q)
	case ir.QArrayDeref:
		x86.EmitArrayDeref(q)
	case ir.QDeref:
		x86.EmitDeref(q)
	case ir.QNewArray:
		x86.EmitNewArray(q)
	case ir.QNewClass:
		x86.EmitNewClass(q)
	default:
		panic("unsupported quadruple")
	}
}

// Returns the address of a variable relative to rbp, or "" if it's a temporary.
func (x86 *X86Generator) getLoc(loc ir.Location) string {
	switch loc := loc.(type) {
	case ir.LReg:
		if loc.Variable == "" {
			return ""
		}
		varOffset := loc.Index
		funOffset := x86.curFun.Offset
		offset := (varOffset - funOffset + 1) * DWORD
		return fmt.Sprintf("%#x(%s)", -offset, ebp)
	case ir.LMem:
		x86.EmitLoad(ecx, loc.Addr)
		return fmt.Sprintf("(%s)", ecx)
	case ir.LSelfField:
		x86.EmitOp("movl %#x(%s), %s #emitload-sf1", -DWORD, ebp, ecx)
		x86.EmitOp("leal %#x(%s), %s #emitload-sf2", DWORD*loc.Offset, ecx, ecx)
		return fmt.Sprintf("(%s)", ecx)
	default:
		panic("unsupported location")
	}
}

func (x86 *X86Generator) EmitLoad(register string, loc ir.Location) {
	// Void doesn't need to be stored anywhere, and won't be popped from stack.
	if _, ok := loc.Type().(types.TVoid); ok {
		return
	}

	switch loc := loc.(type) {
	case ir.LConst:
		switch v := loc.Value.(type) {
		case string:
			x86.EmitOp("movl $%s, %s", x86.stringAdresses[v], register)
		case bool:
			// Value 0 represents false, 1 represents true.
			x := 0
			if v {
				x = 1
			}
			x86.EmitOp("movl $%#x, %s", x, register)
		case int:
			x86.EmitOp("movl $%d, %s", v, register)
		}
	case ir.LReg:
		// Temporaries go on stack, so we need to pop
		if loc.Variable == "" {
			x86.EmitOp("popl %s", register)
		} else {
			x86.EmitOp("movl %s, %s\t# emitload-reg-%s", x86.getLoc(loc), register, loc)
		}
	case ir.LMem:
		x86.EmitLoad(register, loc.Addr)
		x86.EmitOp("movl (%s), %s", register, register)
	case ir.LSelfField:
		x86.EmitOp("movl %#x(%s), %s #emitload-sf1", -DWORD, ebp, ecx)
		x86.EmitOp("movl %#x(%s), %s #emitload-sf2", DWORD*loc.Offset, ecx, register)
	default:
		panic(":(((")
	}

}

func (x86 *X86Generator) EmitQMov(q ir.QMov) {
	x86.dbg(q)

	switch dst := q.Dst.(type) {
	// In case of a standalone expression we can forget the result.
	case ir.LDrop:
		x86.EmitLoad(eax, q.Src)
		return
	case ir.LReg, ir.LMem:
		addr := x86.getLoc(dst)
		x86.EmitLoad(eax, q.Src)
		if addr == "" {
			x86.EmitOp("pushl %s\t# %s", eax, q)
		} else {
			x86.EmitOp("movl %s, %s \t# %s", eax, addr, q)
		}
	case ir.LSelfField:
		x86.EmitLoad(eax, q.Src)
		x86.EmitOp("movl %s, %s", eax, x86.getLoc(dst))
	default:
		panic(":(((")
	}
}

func (x86 *X86Generator) EmitPush(q ir.QPush) {
	x86.dbg(q)

	x86.EmitLoad(eax, q.Src)
	for i := 0; i <= q.Additional; i++ {
		x86.EmitOp("pushl %s", eax)
	}
}

func (x86 *X86Generator) EmitStringAdd(q ir.QBinOp) {
	x86.dbg(q)

	x86.EmitLoad(eax, q.Lhs)
	x86.EmitLoad(ebx, q.Rhs)
	x86.EmitOp("pushl %s", ebx)
	x86.EmitOp("pushl %s", eax)

	x86.EmitOp("call concat")
	x86.EmitOp("addl $%#x, %s", 2*DWORD, esp)

	if dst, ok := q.Dst.(ir.LReg); ok && dst.Variable != "" {
		x86.EmitOp("movl %s, %s", eax, x86.getLoc(dst))
	} else {
		x86.EmitOp("pushl %s", eax)
	}
}

func (x86 *X86Generator) EmitBinOp(q ir.QBinOp) {
	x86.dbg(q)

	if _, ok := q.Lhs.Type().(types.TString); ok {
		x86.EmitStringAdd(q)
		return
	}

	x86.EmitLoad(ebx, q.Rhs)
	x86.EmitLoad(eax, q.Lhs)
	switch q.Op {
	case "+":
		x86.EmitOp("addl %s, %s", ebx, eax)
	case "-":
		x86.EmitOp("subl %s, %s", ebx, eax)
	case "*":
		x86.EmitOp("imull %s", ebx)
	case "/":
		x86.EmitOp("cdq")
		x86.EmitOp("idivl %s", ebx)
	case "%":
		x86.EmitOp("cdq")
		x86.EmitOp("idivl %s", ebx)
		x86.EmitOp("xchg %s, %s", edx, eax)
	default:
		panic("unsupported binary operator")
	}

	switch dst := q.Dst.(type) {
	case ir.LReg:
		if dst.Variable != "" {
			x86.EmitOp("movl %s, %s", eax, x86.getLoc(dst))
		} else {
			x86.EmitOp("pushl %s #%#v", eax, q.Dst)
		}
	default:
		x86.EmitOp("movl %s, %s #xd", eax, x86.getLoc(dst))
	}

}

func (x86 *X86Generator) EmitStringRelOp(q ir.QRelOp) {
	x86.dbg(q)

	x86.EmitLoad(eax, q.Lhs)
	x86.EmitLoad(ebx, q.Rhs)
	x86.EmitOp("pushl %s", ebx)
	x86.EmitOp("pushl %s", eax)

	x86.EmitOp("call compare")
	x86.EmitOp("addl $%#x, %s", 2*DWORD, esp)

	x86.EmitOp("test %s, %s", eax, eax)
	var op string
	switch q.Op {
	case "==":
		op = "jz"
	case "!=":
		op = "jnz"
	default:
		panic("unsupported string operation")
	}

	if q.LNext == q.LTrue {
		x86.EmitOp("%s %s", inverseJmp[op], q.LFalse)
	} else if q.LNext == q.LFalse {
		x86.EmitOp("%s %s", op, q.LTrue)
	} else {
		x86.EmitOp("%s %s", op, q.LTrue)
		x86.EmitOp("jmp %s", op, q.LFalse)
	}
}

var inverseJmp = map[string]string{
	"jz":  "jnz",
	"jnz": "jz",
	"jge": "jl",
	"jg":  "jle",
	"je":  "jne",
	"jne": "je",
	"jl":  "jge",
	"jle": "jg",
}

func (x86 *X86Generator) EmitRelOp(q ir.QRelOp) {
	x86.dbg(q)

	if _, ok := q.Lhs.Type().(types.TString); ok {
		x86.EmitStringRelOp(q)
		return
	}

	x86.EmitLoad(ebx, q.Rhs)
	x86.EmitLoad(eax, q.Lhs)
	x86.EmitOp("cmp %s, %s", ebx, eax)

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
		x86.EmitOp("%s %s", inverseJmp[op], q.LFalse)
	} else if q.LNext == q.LFalse {
		x86.EmitOp("%s %s", op, q.LTrue)
	} else {
		x86.EmitOp("%s %s", op, q.LTrue)
		x86.EmitOp("jmp %s", op, q.LFalse)
	}
}

func (x86 *X86Generator) EmitNeg(q ir.QNeg) {
	x86.dbg(q)

	x86.EmitLoad(eax, q.Arg)
	x86.EmitOp("neg %s", eax)
	x86.EmitOp("pushl %s", eax)
}

func (x86 *X86Generator) EmitJmp(q ir.QJmp) {
	x86.dbg(q)

	x86.EmitOp("jmp %s", q.Dst)
}

func (x86 *X86Generator) EmitJz(q ir.QJz) {
	x86.dbg(q)

	x86.EmitLoad(eax, q.Value)
	x86.EmitOp("test %s, %s", eax, eax)
	x86.EmitOp("jz %s", q.Dst)
}

func (x86 *X86Generator) EmitJnz(q ir.QJnz) {
	x86.dbg(q)

	x86.EmitLoad(eax, q.Value)
	x86.EmitOp("test %s, %s", eax, eax)
	x86.EmitOp("jnz %s", q.Dst)
}

func (x86 *X86Generator) EmitRet(q ir.QRet) {
	x86.dbg(q)

	// if x86.curFun.Function.Name == "main" {
	// 	// Put return value to rdi, it's a scratch register so epilog won't change it.
	// 	x86.EmitLoad(rdi, q.Value)
	// 	x86.EmitFunctionEpilog()
	// 	x86.EmitLoad(rax, ir.LConst{Type_: types.TInt{}, Value: syscallExit})
	// 	x86.EmitOp("syscall")
	// } else {
	x86.EmitLoad(eax, q.Value)
	x86.EmitFunctionEpilog()
	x86.EmitOp("ret")
	// }
}

func (x86 *X86Generator) EmitVRet(q ir.QVRet) {
	x86.dbg(q)

	x86.EmitFunctionEpilog()
	x86.EmitOp("ret")
}

func (x86 *X86Generator) EmitCall(q ir.QCall) {
	x86.dbg(q)

	x86.EmitOp("call %s", q.Label.Name)
	// Pop arguments from the stack.
	x86.EmitOp("addl $%#x, %s", len(q.Args)*DWORD, esp)

	if _, ok := q.Signature.Result.(types.TVoid); !ok {
		x86.EmitOp("pushl %s", eax)
	}
}

func (x86 *X86Generator) EmitCallMethod(q ir.QCallMethod) {
	x86.dbg(q)

	x86.EmitLoad(eax, q.Label)
	x86.EmitOp("call *%s", eax)
	// Pop arguments from the stack.
	x86.EmitOp("addl $%#x, %s", len(q.Args)*DWORD, esp)

	if _, ok := q.Signature.Result.(types.TVoid); !ok {
		x86.EmitOp("pushl %s", eax)
	}
}

func (x86 *X86Generator) EmitArrayAccess(q ir.QArrayAccess) {
	x86.dbg(q)

	x86.EmitLoad(eax, q.Array)
	x86.EmitLoad(ebx, q.Index)

	// First field of an array stores its length.
	displacement := DWORD
	x86.EmitOp("leal %d(%s, %s, %d), %s	#%s", displacement, eax, ebx, DWORD, eax, q)
	x86.EmitOp("pushl %s", eax)
}

func (x86 *X86Generator) EmitArrayDeref(q ir.QArrayDeref) {
	x86.dbg(q)

	x86.EmitLoad(eax, q.Array)
	x86.EmitLoad(ebx, q.Index)

	// First field of an array stores its length.
	displacement := DWORD
	x86.EmitOp("pushl %d(%s, %s, %d)", displacement, eax, ebx, DWORD)
}

func (x86 *X86Generator) EmitDeref(q ir.QDeref) {
	x86.dbg(q)
	x86.EmitLoad(eax, q.Src)
	x86.EmitOp("movl (%s), %s", eax, eax)
	if dst, ok := q.Dst.(ir.LReg); ok {
		if dst.Variable != "" {
			x86.EmitOp("movl %s, %s", eax, x86.getLoc(dst))
		} else {
			x86.EmitOp("pushl %s", eax)
		}
	}
}

func (x86 *X86Generator) EmitNewArray(q ir.QNewArray) {
	x86.dbg(q)
	x86.EmitLoad(eax, q.Size)
	x86.EmitOp("pushl %s", eax)
	x86.EmitCall(ir.QCall{
		Signature: newArraySignature,
		Label:     ir.Label{IsFunction: true, Name: "newArray"},
		Dst:       q.Dst,
		Args:      []ir.Location{q.Size},
	})
}

func (x86 *X86Generator) EmitNewClass(q ir.QNewClass) {
	x86.dbg(q)
	size := ir.LConst{Type_: types.TInt{}, Value: q.Class.TotalNonMethods}
	x86.EmitOp("pushl $%s", x86.vtables[q.Class.ID.GetText()].Label)
	x86.EmitLoad(eax, size)
	x86.EmitOp("pushl %s", eax)
	x86.EmitCall(ir.QCall{
		Signature: newClassSignature,
		Label:     ir.Label{IsFunction: true, Name: "newClass"},
		Dst:       q.Dst,
		Args:      []ir.Location{ir.LConst{}, size},
	})
}

var newArraySignature = types.TFun{
	Ident:  "newArray",
	Args:   []types.FArg{{Ident: "size", Type: types.TInt{}}},
	Result: types.TInt{},
}

var newClassSignature = types.TFun{
	Ident: "newClass",
	Args: []types.FArg{
		{Ident: "vtable", Type: types.TInt{}},
		{Ident: "size", Type: types.TInt{}},
	},
	Result: types.TInt{},
}

func (x86 *X86Generator) dbg(q ir.Quadruple) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	trace := f.Name()
	x86.EmitOp("# %s %s", q, trace)
}
