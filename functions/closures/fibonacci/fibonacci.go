package main

import (
	"fmt"
	"time"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	return func (x int) int {
		if x == 0 || x == 1 {
			return x
		}
		fib := fibonacci()
		return fib(x - 1) + fib(x - 2)
	}
}

// because its recursive with no memoization
// this function takes a REALLY long time to run......
// i stopped it in the middle of processing up to 1000
// go memoized the values of the previous called methods for me! woah!
// so it sped the time up, but this is how long it took: over 10 minutes and that's going to grow at 2n yikes!!!!!!
// so i started over to see how long 51 would be
// 6m9.227698433s
func main() {
	start := time.Now()
	f := fibonacci()
	for i := 0; i < 51; i++ {
		fmt.Println(f(i), "i at:", i)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
