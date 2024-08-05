package main

import (
	. "example.com/gen"
)

type AdderImpl struct {
}

// Implement the ExportsDocsAdder0_1_0_Add interface to ensure the component satisfies the
// `adder` world
func (i AdderImpl) Add(x, y uint32) uint32 {
	return x + y
}

// To enable our component to be a library, implement the component in the
// `init` function which is always called first when a Go package is run.
func init() {
	example := AdderImpl{}
	SetExportsDocsAdder0_1_0_Add(example)
}

// main is required for the `WASI` target, even if it isn't used.
func main() {}
