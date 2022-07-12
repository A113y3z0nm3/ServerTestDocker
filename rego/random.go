package rego

import (
	"math/rand"
	"io/ioutil"
)

func Random(p string) string {
	files, _ := ioutil.ReadDir(p)
	return files[rand.Intn(len(files))].Name()
}
