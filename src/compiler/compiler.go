package compiler

import (
	"fmt"
	"io/ioutil"
	"latte/compiler/backend"
	. "latte/compiler/config"
	"latte/compiler/frontend"
	"os"
)

func CompileX64(cfg Config) error {
	s, err := frontend.Run(cfg.Source)
	if err != nil {
		return err
	}

	if cfg.TypeCheck {
		return nil
	}

	asm := backend.GenX64(s, cfg)

	if err := saveAssembly(asm, cfg.Intermediate); err != nil {
		return fmt.Errorf("failed to save assembly to .s file: %w", err)
	}

	linkFile, err := ioutil.TempFile(os.TempDir(), "latte_link")
	if err != nil {
		return err
	}
	defer os.Remove(linkFile.Name())

	if err := compileGCC(cfg.Intermediate, linkFile.Name()); err != nil {
		return fmt.Errorf("gcc compilation error: %w", err)
	}

	if err := link(cfg.Runtime, linkFile.Name(), cfg.Target); err != nil {
		return fmt.Errorf("link error: %w", err)
	}

	return nil
}
