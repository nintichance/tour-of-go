package main

import "golang.org/x/tour/pic"
import "math/rand"

// Pic is a demonstration of 2D arrays, looping, range, and coercion in go
func Pic(dx, dy int) [][]uint8 {
 y := make([][]uint8, dy)
 
 for i := range y {
 	x := make([]uint8, dx)
	for j := range x {
		x[j] = uint8(rand.Intn(255))
	}
 	y[i] = x
 }
 
 return y
}

func main() {
	pic.Show(Pic)
}