package main

import "fmt"

func changeValueByPointer(x *int) {
	newValue := 2
	fmt.Println("Changing value now to: ", newValue)

	*x = newValue
}

func main() {
	var myNum int = 5
	fmt.Println("Setting number to: ", myNum)
	changeValueByPointer(&myNum)
	fmt.Println("New value set to: ", myNum)
}
