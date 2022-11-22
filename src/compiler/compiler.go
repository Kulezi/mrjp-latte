package compiler

import (
	"fmt"
	"latte/parser"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

const sampleASM = `; ----------------------------------------------------------------------------------------
; Writes "Hello, World" to the console using only system calls. Runs on 64-bit Linux only.
; To assemble and run:
;
;     nasm -felf64 hello.asm && ld hello.o && ./a.out
; ----------------------------------------------------------------------------------------

          global    _start

          section   .text
_start:   mov       rax, 1                  ; system call for write
          mov       rdi, 1                  ; file handle 1 is stdout
          mov       rsi, message            ; address of string to output
          mov       rdx, 13                 ; number of bytes
          syscall                           ; invoke operating system to do the write
          mov       rax, 60                 ; system call for exit
          xor       rdi, rdi                ; exit code 0
          syscall                           ; invoke operating system to exit

          section   .data
message:  db        "Hello, World", 10      ; note the newline at the end`

type compiler struct {
	tree    parser.IProgramContext
	globals map[string]Type
	parent  map[string]TClassRef
}

func (c *compiler) SemanticCheck() error {
	// Evaluate method/function signatures and inheritance tree.
	visitor := MakeGlobalDeclVisitor()
	if err := visitor.Run(c.tree); err != nil {
		return err
	}

	c.globals = visitor.Globals
	c.parent = visitor.Parent

	// if err := c.TypeClassMethods(); err != nil {
	// 	return err
	// }

	fmt.Println(c.globals, c.parent)
	// fmt.Fprintln(os.Stderr, (&TypeCheckVisitor{}).Visit(c.tree))
	return nil
}

func (c *compiler) Parse(filename string) error {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return err
	}

	lexErrors := &CustomErrorListener{}
	lexer := parser.NewLatteLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexErrors)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	parseErrors := &CustomErrorListener{}
	p := parser.NewLatteParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(parseErrors)
	p.BuildParseTrees = true
	c.tree = p.Program()

	if err := lexErrors.Check("lexer error:"); err != nil {
		return err
	}

	if err := parseErrors.Check("parse error:"); err != nil {
		return err
	}

	return c.SemanticCheck()
}

// FIXME: generate code.
func GenX64(filename string) (string, error) {
	c := compiler{}
	if err := c.Parse(filename); err != nil {
		return "", err
	}

	return sampleASM, nil
}

func saveAssembly(asm, filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(asm); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func compileNASM(sourcePath, targetPath string) error {
	out, err := exec.Command(
		"nasm",
		"-f",
		"elf64",
		"-o",
		targetPath,
		sourcePath,
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s", err, out)
	}

	return nil
}

func link(sourcePath, targetPath string) error {
	out, err := exec.Command(
		"ld",
		sourcePath,
		"-o",
		targetPath).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s", err, out)
	}

	_ = os.Remove(sourcePath)

	return nil
}

func CompileX64(filename string) error {
	basename := strings.TrimSuffix(filename, path.Ext(filename))

	asm, err := GenX64(filename)
	if err != nil {
		return fmt.Errorf("code generation failed: %w", err)
	}

	if err := saveAssembly(asm, basename+".s"); err != nil {
		return fmt.Errorf("failed to save assembly to .s file: %w", err)
	}

	if err := compileNASM(basename+".s", basename+".o"); err != nil {
		return fmt.Errorf("nasm compilation error: %w", err)
	}

	if err := link(basename+".o", basename); err != nil {
		return fmt.Errorf("link error: %w", err)
	}

	return nil
}
