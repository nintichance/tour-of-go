package main

import "fmt"

// Vertex is a struct, which can store multiple related items into a contiguous block of memory
// therefore only a single pointer is needed to access all values within a particular struct
// a struct is similar to a record in a database as it stores multiple data types related to an entity
type Vertex struct {
	Lat,Long float64
	Butt string
}

// m is a map (map declaration?)
// a map in go is a hash map 
// collection of key-value pairs
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	// map literals are like struct literals but keys are required
	m["Bells Labs"] = Vertex{
		40.68433, 
		-74.39967,
		"mutt",
	}
	fmt.Println(m["Bells Labs"].Butt)
}