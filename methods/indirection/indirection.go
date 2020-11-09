package main

import "fmt"

// Vertex is a struct
// eventually I'll refactor to only use one copy of Vertex
type Vertex struct {
	X, Y float64
}

// Scale scales the size of the collections
// it is a method
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleFunc is a function
// does the same
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}