// +build OMIT

package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func (v Vertex) ValueRecevierMethod() {
	v.X = 30 
}

func (v *Vertex) PointerRecevierMethod() {
	v.Y = 40 
}

func main() {
	v := Vertex{3, 4}
	v.ValueRecevierMethod()
	v.PointerRecevierMethod()
	fmt.Println(v)
}
