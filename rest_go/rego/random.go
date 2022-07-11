package rego

import (
	"math/rand"
	"io/ioutil"
	"time"
)

func Random(p string) string {
	files, _ := ioutil.ReadDir(p)
	rand.Seed(time.Now().UTC().UnixNano())
	return files[rand.Intn(len(files))].Name()
}
