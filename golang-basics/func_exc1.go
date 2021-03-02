package main

import "fmt"

func foo() int {
	return 23
}

func bar() (int, string) {
	return 23, "Nice"
}

func main() {
	fmt.Println("Foo: ", foo())

	num, txt := bar()
	fmt.Printf("Bar: %v and %v\n", num, txt)
}
