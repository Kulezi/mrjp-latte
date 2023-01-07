package backend

import (
	"latte/compiler/backend/ir"
	. "latte/compiler/config"
	"latte/compiler/frontend"
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

// FIXME: generate code.
func GenX64(s frontend.State, config Config) string {
	cfg := ir.Generate(s, config)
	// v := ir.MakeVisitor(typecheck.MakeVisitor(s.Signatures), cfg)
	// v.Visit(s.Tree)
	// ir := ""
	// for _, block := range v.Blocks {
	// 	ir += block.Label + "\n"
	// 	for _, op := range block.Ops {
	// 		ir += "\t" + op.String() + "\n"
	// 	}
	// }
	return cfg.String()
}
