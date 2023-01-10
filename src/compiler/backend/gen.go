package backend

import (
	"fmt"
	"latte/compiler/backend/ir"
	. "latte/compiler/config"
	"latte/compiler/frontend"
	"latte/compiler/frontend/types"
)

// const sampleASM = `; ----------------------------------------s------------------------------------------------
// ; Writes "Hello, World" to the console using only system calls. Runs on 64-bit Linux only.
// ; To assemble and run:
// ;
// ;     nasm -felf64 hello.asm && ld hello.o && ./a.out
// ; ----------------------------------------------------------------------------------------

//           global    _start

//           section   .text
// _start:   mov       rax, 1                  ; system call for write
//           mov       rdi, 1                  ; file handle 1 is stdout
//           mov       rsi, message            ; address of string to output
//           mov       rdx, 13                 ; number of bytes
//           syscall                           ; invoke operating system to do the write
//           mov       rax, 60                 ; system call for exit
//           xor       rdi, rdi                ; exit code 0
//           syscall                           ; invoke operating system to exit

//           section   .data
// message:  db        "Hello, World", 10      ; note the newline at the end`

// const sampleATT = `.data
// hello:
//     .string "Hello world!\n"

// .text
// .globl _start
// _start:
//     movl $4, %eax # write(1, hello, strlen(hello))
//     movl $1, %ebx
//     movl $hello, %ecx
//     movl $13, %edx
//     int  $0x80

//     movl $1, %eax # exit(0)
//     movl $0, %ebx
//     int  $0x80`

const prologATT = `.text
.globl _start
`

// FIXME: generate code.
func GenX64(s frontend.State, config Config) string {
	// cfg := MakeSSA(ir.Generate(s, config))

	cfg := ir.Generate(s, config)
	for _, block := range cfg.Nodes {
		fmt.Println(block.Label)
		for _, op := range block.Ops {
			fmt.Println("\t" + op.String())
		}
	}

	x64 := X64Generator{}
	x64.Emit(prologATT)
	for _, block := range cfg.Nodes {
		x64.GenFromBlock(block)
	}

	return x64.res
}

type X64Generator struct {
	curFun ir.Label
	res    string
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

func (x64 *X64Generator) GenFromBlock(block ir.BasicBlock) {
	if block.Label.IsFunction {
		x64.curFun = block.Label
	}

	if block.Label.Name == "main" {
		x64.EmitLabel("_start")
	} else {
		x64.EmitLabel(block.Label.Name)
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
	case ir.QUnOp:
		x64.EmitUnOp(q)
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
	case ir.QPop:
		x64.EmitPop(q)
	default:
		panic("unsupported quadruple")
	}
}

const (
	rax = "%%rax"
	rbx = "%%rbx"
	rcx = "%%rcx"
	rdx = "%%rdx"
	rdi = "%%rdi"
)

func (x64 *X64Generator) EmitLoad(register string, loc ir.Location) {
	switch loc := loc.(type) {
	case ir.LConst:
		x64.EmitOp(fmt.Sprintf("movq $%s, %s", loc, register))
	case ir.LReg:
		// Temporaries go on stack, so we need to pop
		if loc.Variable == "" {
			x64.EmitOp("popq %s", register)
		} else {
			x64.EmitOp("movq %%%s, %s", loc, register)
		}
	}
}

func (x64 *X64Generator) EmitQMov(q ir.QMov) {
	dst := q.Dst.(ir.LReg)
	x64.EmitLoad(rax, q.Src)
	x64.EmitOp("mov %s, %s", rax, dst)
}

func (x64 *X64Generator) EmitBinOp(q ir.QBinOp) {
	x64.EmitLoad(rax, q.Lhs)
	x64.EmitLoad(rbx, q.Rhs)
	switch q.Op {
	case "+":
		x64.EmitOp("addq %s, %s", rbx, rax)
	case "-":
		x64.EmitOp("subq %s, %s", rbx, rax)
	case "*":
		x64.EmitOp("imulq %s", rbx)
	case "/":
		x64.EmitOp("divq %s", rbx)
	case "%":
		x64.EmitOp("divq %s", rbx)
		x64.EmitOp("xchg %s, %s", rdx, rax)
	default:
		panic("unsupported binary operator")
	}

	x64.EmitOp("pushq %s", rax)
}

func (x64 *X64Generator) EmitRelOp(q ir.QBinOp) {
	x64.EmitLoad(rax, q.Lhs)
	x64.EmitLoad(rbx, q.Rhs)
	switch q.Op {
	case ">":
		x64.EmitOp("cmp %s, %s", rax, rbx)
		x64.EmitOp("setg %s", rax)
	// TODO: "<", "<=", "==", "!=", ">="
	default:
		panic("unsupported binary operator")
	}
}

func (x64 *X64Generator) EmitUnOp(q ir.QUnOp) {
	x64.EmitLoad(rax, q.Arg)
	switch q.Op {
	case "!":
		x64.EmitOp("neg %s", rax)
	case "-":
		x64.EmitOp("not %s", rax)
	default:
		panic("unsupported unary operator")
	}

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
	x64.EmitOp("jz %s", q.Dst)
}

func (x64 *X64Generator) EmitRet(q ir.QRet) {
	if x64.curFun.Name == "main" {
		x64.EmitLoad(rax, ir.LConst{Type_: types.TInt{}, Value: 60})
		x64.EmitLoad(rdi, q.Value)
		x64.EmitOp("syscall")
	} else {
		x64.EmitLoad(rax, q.Value)
		x64.EmitOp("ret")
	}
}

func (x64 *X64Generator) EmitVRet(q ir.QVRet) {
	x64.EmitOp("ret")
}

func (x64 *X64Generator) EmitPop(q ir.QPop) {
	dst := q.Dst.(ir.LReg)
	x64.EmitOp("pop %s", rax)
	x64.EmitOp("mov %s, %s", rax, dst)
}

func (x64 *X64Generator) EmitCall(q ir.QCall) {
	x64.EmitOp("call %s", q.Label.Name)
	// TODO: use result
}
