package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount accepts a string and returns a map with counts of occurences in string
func WordCount(s string) map[string]int {
	arr := strings.Split(s, " ") 
	m := make(map[string]int)
	for _,val := range arr {
		if m[val] == 0 {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
