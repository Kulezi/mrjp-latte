package main

import (
	"fmt"
	"latte/compiler"
	"os"
)

func checkArgs() bool {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <input.lat>\n", os.Args[0])
		return false
	}

	return true
}

func main() {
	if !checkArgs() {
		return
	}

	if err := compiler.CompileX64(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "OK")
}
