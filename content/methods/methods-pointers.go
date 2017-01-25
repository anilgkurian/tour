// +build OMIT

package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func (v Vertex) ValueMethod() {
	v.X = 30 
}

func (v *Vertex) PointerMethod() {
	v.Y = 40 
}

func main() {
	v := Vertex{3, 4}
	v.ValueMethod()
	v.PointerMethod()
	fmt.Println(v)
}
