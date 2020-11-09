package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value is:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value is:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value is:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value is:", v, "Present?", ok)

}