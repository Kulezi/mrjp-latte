package compiler

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"latte/compiler/config"
	"latte/compiler/frontend/types"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
)

var goodDirs = []string{
	"../../tests/kulezi_tests/good",
	"../../tests/kulezi_tests/extensions/arrays1",
	"../../tests/kulezi_tests/extensions/objects1",
	"../../tests/kulezi_tests/extensions/objects2",
	"../../tests/kulezi_tests/extensions/struct",

	// "../../mgtests/lattests/tests/margdoc",
	// "../../mgtests/lattests/tests/margdoc/lvalues",
	"../../tests/mrjp-tests/good/basic",
	"../../tests/mrjp-tests/good/arrays",
	// "../../mrjp-tests/good/hardcore",
	// "../../mrjp-tests/good/virtual",
	// "../../mrjp-tests/gr5",
	"../../tests/official_tests/good",
}

var badDirs = []string{
	"../../tests/official_tests/bad",
	"../../tests/kulezi_tests/bad",
	"../../tests/mrjp-tests/bad/semantic",
	"../../tests/mgtests/lattests/tests/margdoc/bad",
	"../../tests/mgtests/lattests/tests/margdoc/casting",
}

func TestTypecheckGood(t *testing.T) {
	for _, dir := range goodDirs {
		items, err := ioutil.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}

		for _, item := range items {
			if !item.IsDir() {
				filename := dir + "/" + item.Name()
				if path.Ext(filename) == ".lat" {
					t.Run(filename, func(t *testing.T) {
						// t.Parallel()
						basename := strings.TrimSuffix(filename, ".lat")

						intermediate := basename + ".s"
						err := CompileX64(config.Config{
							TypeCheck:    true,
							Source:       filename,
							Intermediate: intermediate,
						})

						if err != nil {
							t.Fatal(err)
						}
					})
				}
			}
		}
	}
}

func TestTypecheckBad(t *testing.T) {
	for _, dir := range badDirs {
		items, err := ioutil.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}

		for _, item := range items {
			if !item.IsDir() {
				filename := dir + "/" + item.Name()
				if path.Ext(filename) == ".lat" {
					t.Run(filename, func(t *testing.T) {
						// t.Parallel()
						err := CompileX64(config.Config{
							Source: filename,
						})
						if err == nil {
							t.Fatal("test didn't fail!")
						} else {
							t.Log(err)
						}
					})
				}
			}
		}
	}
}

func getRuntimePath(t *testing.T) string {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	return path.Join(wd, "../../lib/runtime.o")
}

func TestGoodCompile(t *testing.T) {
	for _, dir := range goodDirs {
		items, err := ioutil.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}

		for _, item := range items {
			if !item.IsDir() {
				filename := dir + "/" + item.Name()
				if path.Ext(filename) == ".lat" {
					t.Run(filename, func(t *testing.T) {
						// t.Parallel()
						basename := strings.TrimSuffix(filename, ".lat")

						intermediate := basename + ".s"

						err := CompileX64(config.Config{
							Source:       filename,
							Intermediate: intermediate,
							Target:       basename,
							Runtime:      getRuntimePath(t),
						})

						if err != nil {
							t.Fatal(err)
						}
					})
				}
			}
		}
	}
}

func TestGoodAnswers(t *testing.T) {
	for _, dir := range goodDirs {
		items, err := ioutil.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}

		for _, item := range items {
			if !item.IsDir() {
				filename := dir + "/" + item.Name()
				if path.Ext(filename) == ".lat" {
					t.Run(filename, func(t *testing.T) {
						// t.Parallel()
						basename := strings.TrimSuffix(filename, ".lat")

						intermediate := basename + ".s"
						err = CompileX64(config.Config{
							Source:       filename,
							Intermediate: intermediate,
							Target:       basename,
							Runtime:      getRuntimePath(t),
						})

						if err != nil {
							t.Fatal(err)
						}

						cmd := exec.Command(fmt.Sprintf("./%s", basename))
						inputFile, err := os.Open(basename + ".input")
						if err == nil {
							cmd.Stdin = inputFile
						}

						outputFile, err := os.Create(basename + ".output_test")
						if err != nil {
							t.Fatal(err)
						}
						cmd.Stdout = outputFile

						if err := cmd.Start(); err != nil {
							t.Fatal(err)
						}

						if err := cmd.Wait(); err != nil {
							if exiterr, ok := err.(*exec.ExitError); ok {
								t.Logf("Exit Status: %d", exiterr.ExitCode())
							} else {
								t.Fatalf("cmd.Wait: %v", err)
							}
						}

						defer func() {
							_ = os.Remove(intermediate)
							_ = os.Remove(basename)
							outputFile.Close()
							_ = os.Remove(basename + ".output_test")
						}()

						diff, err := exec.Command("cmp", basename+".output_test", basename+".output").CombinedOutput()
						t.Log(string(diff))
						if err != nil {
							t.Fatal(err)
						} else {
							if len(diff) > 0 {
								t.Fatal(diff)
							}
						}
					})
				}
			}
		}
	}
}

var binBoolOps = []string{"<", "<=", "==", "!=", ">=", ">", "&&", "||"}
var binIntOps = []string{"+", "-", "/", "*", "%"}
var binStringOps = []string{"+"}
var unBoolOps = []string{"!"}
var unIntOps = []string{"-"}

func randInt(l, r int) int {
	return int(rand.Int63())%(r-l+1) + l
}

func randElem[T any](slice []T) T {
	return slice[int(rand.Int31())%len(slice)]
}

func randomConstIntExpr(maxDepth int) string {
	if randInt(0, 10) == 0 || maxDepth < 0 {
		return fmt.Sprintf("%d", randInt(0, 10))
	}

	var op string
	if randInt(0, 4) > 0 {
		op = randElem(binIntOps)
		return "(" + randomConstIntExpr(maxDepth-1) + op + randomConstIntExpr(maxDepth-1) + ")"
	} else {
		op = randElem(unIntOps)
		return "(" + op + randomConstIntExpr(maxDepth-1) + ")"
	}

}

func randomConstStringExpr(maxDepth int) string {
	if randInt(0, 10) == 0 || maxDepth < 0 {
		return fmt.Sprintf("\"%d\"", randInt(-10, 10))
	}

	op := randElem(binStringOps)
	return "(" + randomConstStringExpr(maxDepth-1) + op + randomConstStringExpr(maxDepth-1) + ")"
}

func randomConstBoolExpr(maxDepth int) string {
	if randInt(0, 10) == 0 || maxDepth < 0 {
		if randInt(0, 1) == 0 {
			return "false"
		} else {
			return "true"
		}
	}

	var op string
	if randInt(0, 4) > 0 {
		op = randElem(binBoolOps)
		if op == "&&" || op == "||" {
			return "(" + randomConstBoolExpr(maxDepth-1) + op + randomConstBoolExpr(maxDepth-1) + ")"
		}

		switch randInt(0, 2) {
		case 0:
			if !(op == "<" || op == ">" || op == ">=" || op == "<=") {
				return "(" + randomConstBoolExpr(maxDepth-1) + op + randomConstBoolExpr(maxDepth-1) + ")"
			}
			fallthrough
		case 1:
			if !(op == "<" || op == ">" || op == ">=" || op == "<=") {
				return "(" + randomConstStringExpr(maxDepth-1) + op + randomConstStringExpr(maxDepth-1) + ")"
			}
			fallthrough
		default:
			return "(" + randomConstIntExpr(maxDepth-1) + op + randomConstIntExpr(maxDepth-1) + ")"
		}
	} else {
		op = randElem(unBoolOps)
		return "(" + op + randomConstBoolExpr(maxDepth-1) + ")"
	}

}
func randomValidConstExpr(maxDepth int) string {
	rt := randInt(1, 3)
	switch rt {
	case 1:
		return randomConstIntExpr(maxDepth)
	case 2:
		return randomConstStringExpr(maxDepth)
	default:
		return randomConstBoolExpr(maxDepth)
	}
}

func TestRandomConstExpr(t *testing.T) {
	dir := os.TempDir()
	for i := 0; i < 1000; i++ {
		s := "int main() {\n\t" + randomValidConstExpr(5) + ";\treturn 0;\n}\n"
		f, err := os.CreateTemp(dir, "expr_test")
		if err != nil {
			t.Fatal(err)
		}

		os.WriteFile(f.Name(), []byte(s), fs.FileMode(os.O_RDONLY))
		err = CompileX64(config.Config{
			Source:    f.Name(),
			TypeCheck: true,
		})
		if err != nil {
			if _, ok := err.(types.ZeroDivisionError); !ok {
				t.Fatal(err)
			}
		}
	}
}
