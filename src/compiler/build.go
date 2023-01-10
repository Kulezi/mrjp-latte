package compiler

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func saveAssembly(asm, filepath string) error {
	log.Println(filepath)
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
		"-O0",
		"-m64",
		"-fno-pie",
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
		"gcc",
		sourcePath,
		"/home/pawelputra/studia/mrjp/latte/lib/runtime.o",
		"-o",
		targetPath).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s", err, out)
	}

	_ = os.Remove(sourcePath)

	return nil
}
