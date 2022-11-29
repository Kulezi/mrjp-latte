package compiler

import (
	"io/ioutil"
	"path"
	"testing"
)

var goodDirs = []string{
	"../../lattests/good",
	"../../mgtests/lattests/tests/margdoc",
	"../../mgtests/lattests/tests/margdoc/lvalues",
	"../../mrjp-tests/good/basic",
	"../../mrjp-tests/good/arrays",
	"../../mrjp-tests/good/hardcore",
	"../../mrjp-tests/good/virtual",
	"../../mrjp-tests/gr5",
}

var badDirs = []string{
	"../../lattests/bad",
	"../../mrjp-tests/bad/semantic",
	"../../mgtests/lattests/tests/margdoc/bad",
	"../../mgtests/lattests/tests/margdoc/casting",
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
						err := CompileX64(filename)
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
						err := CompileX64(filename)
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
