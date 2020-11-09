package main

import "fmt"

// a map literal is the value within the map
// var m = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,	
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,
// 	},
// }

// you can omit the typing of the individual keys if type is declared on highest level

func main() {
	m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967, "glut"},
		"Google":    {37.42202, -122.08408, "wut"},
	}
	fmt.Println(m)
}