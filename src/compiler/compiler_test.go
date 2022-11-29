package compiler

import (
	"io/ioutil"
	"path"
	"testing"
)

var goodDirs = []string{
	"../../lattests/good",
}

var badDirs = []string{
	"../../lattests/bad",
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
