package compiler

import (
	"latte/parser"

	. "latte/compiler/frontend/types"
)

type compiler struct {
	tree    parser.IProgramContext
	globals map[string]Type
	parent  map[string]TClassRef
}

func CompileX64(filename string) error {
	// basename := strings.TrimSuffix(filename, path.Ext(filename))

	_, err := genX64(filename)
	if err != nil {
		return err
	}

	// if err := saveAssembly(asm, basename+".s"); err != nil {
	// 	return fmt.Errorf("failed to save assembly to .s file: %w", err)
	// }

	// if err := compileNASM(basename+".s", basename+".o"); err != nil {
	// 	return fmt.Errorf("nasm compilation error: %w", err)
	// }

	// if err := link(basename+".o", basename); err != nil {
	// 	return fmt.Errorf("link error: %w", err)
	// }

	return nil
}
