package main

import (
	"fmt"
	"math"
)

// MyFloat is a float
// to use a data type with a method (function  receiver)
// you must declare the type within the same package
type MyFloat float64

// Abs is a method (function receiver)
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}