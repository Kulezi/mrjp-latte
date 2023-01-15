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

func compileGCC(sourcePath, targetPath string) error {
	out, err := exec.Command(
		"gcc",
		"-c",
		"-O0",
		"-m64",
		"-fno-pie",
		"-no-pie",
		"-o",
		targetPath,
		sourcePath,
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s", err, out)
	}

	return nil
}

func link(runtimePath, sourcePath, targetPath string) error {
	out, err := exec.Command(
		"gcc",
		sourcePath,
		runtimePath,
		"-O0",
		"-m64",
		"-no-pie",
		"-fno-pie",
		"-o",
		targetPath).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %s", err, out)
	}

	_ = os.Remove(sourcePath)

	return nil
}
