package config

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

const optLevels = 3

type Config struct {
	TypeCheck bool
	PrintIR   bool

	// Level one optimizations.
	AllocRegisters                bool
	EliminateCommonSubexpressions bool

	// Level two optimizations.
	ReduceStrength             bool
	OptimizeInductionVariables bool

	// Level three optimizations.
	InlineFunctions bool

	Source       string
	Intermediate string
	Target       string
	Runtime      string
}

func ReadConfig() Config {
	config := Config{}

	flag.BoolVar(
		&config.PrintIR,
		"print-ir",
		false,
		"Print intermediate representation (doesn't work with typecheck flag)",
	)

	flag.BoolVar(
		&config.TypeCheck,
		"typecheck",
		false,
		"Check program correctness without compiling",
	)

	flag.StringVar(
		&config.Target,
		"o",
		"",
		"Target binary name",
	)

	useOptLevel := make([]bool, optLevels+1)
	flag.BoolVar(
		&useOptLevel[0],
		"O0",
		false,
		"Don't use any optimizations",
	)

	for i := 1; i <= 3; i++ {
		flag.BoolVar(
			&useOptLevel[i],
			fmt.Sprintf("O%d", i),
			false,
			fmt.Sprintf("Use all optimizations with level <= %d", i),
		)
	}

	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "ERROR\nusage: %s [flags] <input.lat>\n", os.Args[0])
		os.Exit(1)
	}

	provided := 0
	for i := 0; i <= optLevels; i++ {
		if useOptLevel[i] {
			provided++
			if provided > 1 {
				fmt.Fprintf(os.Stderr, "ERROR\nOnly one -O<n> flag can be used")
				os.Exit(1)
			}

			switch i {
			case 1:
				config.AllocRegisters = true
				config.EliminateCommonSubexpressions = true
			case 2:
				config.OptimizeInductionVariables = true
				config.ReduceStrength = true
			case 3:
				config.InlineFunctions = true
			}
		}
	}

	config.Source = flag.Arg(0)

	basename := strings.TrimSuffix(config.Source, path.Ext(config.Source))
	if config.Target == "" {
		config.Target = basename
	}

	config.Intermediate = basename + ".s"

	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR\n", err)
		os.Exit(1)
	}

	config.Runtime = path.Join(wd, "lib/runtime.o")
	return config
}
