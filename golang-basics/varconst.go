package main

import (
	"fmt"
)

const x = "Spring is here!"
const y int = 24

var (
	glob1 = "xxx"
	glob2 = x
)

func main() {

	local := true
	var localy int = y

	fmt.Println("Untyped const: ", x)
	fmt.Println("Typed const (int): ", y)
	fmt.Println("Global var: ", glob1)
	fmt.Println("Global var assigned from x: ", glob2)
	fmt.Println("Local by direct assignemnt: ", local)
	fmt.Println("Local var by assignment: ", localy)

	/*
	  pc golang-basics % go run varconst.go
	  Untyped const:  Spring is here!
	  Typed const (int):  24
	  Global var:  xxx
	  Global var assigned from x:  Spring is here!
	  Local by direct assignemnt:  true
	  Local var by assignment:  24
	*/

}
