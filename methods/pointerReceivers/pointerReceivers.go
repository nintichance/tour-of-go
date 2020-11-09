package main

import (
	"fmt"
	"math"
)

// Vertex is a struct
type Vertex struct {
	X, Y float64
}

// Abs is a method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale is a method (pointer receiver)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}