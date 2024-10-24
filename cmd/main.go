package main

import (
	"bincooo/sdk-examples/cmd/annotation"
	"github.com/bincooo/sdk/gen"
)

func init() {
	gen.Alias[annotation.GetMapping]()
	gen.Alias[annotation.PutMapping]()
	gen.Alias[annotation.DelMapping]()
	gen.Alias[annotation.PostMapping]()
}

// @Gen(target="wire")
func main() {
	gen.Process()
}
