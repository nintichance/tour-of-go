package main

import (
	"fmt"
	"time"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int, map[int]int) int {
	return func (x int, m map[int]int) int {
		fib := fibonacci()
		if m[x] == 0 {
			if x == 0 || x == 1 {
				m[x] = x
				return x
			} 
			result := fib(x - 1, m) + fib(x - 2, m)
			m[x] = result
			return m[x]
		} 
		return m[x]
	}
}

func main() {
	start := time.Now()
	f := fibonacci()
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		result := f(i, m)
		fmt.Println(result, "i at:", i)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
