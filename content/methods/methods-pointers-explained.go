// +build OMIT

package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func ValueFunction(v Vertex)  {
	 v.X = 30
}

func PointerFunction(v *Vertex) {
	v.Y = 40
}

func main() {
	v := Vertex{3, 4}
	ValueFunction(v)
	PointerFunction(&v)
	fmt.Println(v)
}
