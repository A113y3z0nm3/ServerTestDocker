package rego

import (
	"math/rand"
	"time"
	"io/ioutil"
)

func random(p string) string {
	files, _ := ioutil.ReadDir(p)
	rand.Seed(time.Now().UTC().UnixNano())
	return files[rand.Intn(len(files))].Name()
}
