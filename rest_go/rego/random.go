package rego

import (
	"math/rand"
	"io/ioutil"
)

func Random(p string, t int64) string {
	files, _ := ioutil.ReadDir(p)
	rand.Seed(t)
	return files[rand.Intn(len(files))].Name()
}
