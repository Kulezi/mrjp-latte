package compiler

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"latte/compiler/config"
	"latte/compiler/frontend/types"
	"math/rand"
	"os"
	"path"
	"strings"
	"testing"
)

var goodDirs = []string{
	// "../../lattests/good",
	// "../../mgtests/lattests/tests/margdoc",
	// "../../mgtests/lattests/tests/margdoc/lvalues",
	// "../../mrjp-tests/good/basic",
	// "../../mrjp-tests/good/arrays",
	// "../../mrjp-tests/good/hardcore",
	// "../../mrjp-tests/good/virtual",
	// "../../mrjp-tests/gr5",
	"../../defaultowe_testy/lattests/good",
}

var badDirs = []string{
	// "../../lattests/bad",
	// "../../mrjp-tests/bad/semantic",
	// "../../mgtests/lattests/tests/margdoc/bad",
	// "../../mgtests/lattests/tests/margdoc/casting",
}

func TestGood(t *testing.T) {
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
						basename := strings.TrimSuffix(filename, ".lat")

						intermediate := basename + ".s"
						err := CompileX64(config.Config{
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

func TestBad(t *testing.T) {
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
						err := CompileX64(config.Config{
							TypeCheck: true,
							Source:    filename,
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
			return "(" + randomConstStringExpr(maxDepth-1) + op + randomConstStringExpr(maxDepth-1) + ")"
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
