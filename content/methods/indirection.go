// +build OMIT

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) PointerMethod() {
	v.X = v.X * 10
	v.Y = v.Y * 10
}

func PointerFunction(v *Vertex) {
	v.X = v.X + 1
	v.Y = v.Y + 1
}

func main() {
	v := Vertex{3, 4}
	v.PointerMethod()
	PointerFunction(&v)

	p := &Vertex{5, 6}
	p.PointerMethod()
	PointerFunction(p)

	fmt.Println(v, p)
}