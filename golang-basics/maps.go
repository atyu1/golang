package main

import "fmt"

func main() {
	points := map[string]int{
		"John":  4,
		"Jack":  3,
		"Peter": 7,
	}

	peterPoint, ok := points["Peter"]
	fmt.Println(peterPoint, ok)

	larryPoint, ok := points["Larry"]
	fmt.Println(larryPoint, ok)

	//Delete
	delete(points, "John")
	fmt.Println(points)

	/*
		  go run maps.go
		7 true
		0 false
	*/

}
