// +build OMIT

package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func (v Vertex) ValueMethod() {
	v.X = v.X * 10
	v.Y = v.Y * 10
}

func ValueFunction(v Vertex) {
	v.X = v.X + 1
	v.Y = v.Y + 1
}

func main() {
	v1 := Vertex{3, 4}
	v1.ValueMethod()
	ValueFunction(v1)
	fmt.Println(v1)


	v2 := &Vertex{5, 6}
	v2.ValueMethod()
	ValueFunction(*v2)
	fmt.Println(v2)
	
}