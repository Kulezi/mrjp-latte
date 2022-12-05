package compiler

import (
	"fmt"
	"os"
	"os/exec"
)

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

func compileGCC(sourcePath, targetPath string) error {
	out, err := exec.Command(
		"gcc",
		"-c",
		"-m64",
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
