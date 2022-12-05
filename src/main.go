package main

import (
	"fmt"
	"latte/compiler"
	"latte/compiler/config"
	"os"
)

func main() {
	config := config.ReadConfig()
	if err := compiler.CompileX64(config); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "OK")
}
