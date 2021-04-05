package main

import "fmt"

func main() {
	//Copy (create new copies without refference)
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Create a slice x:", x)
	fmt.Println("x len=", len(x), ", cap=", cap(x))

	//Identical new copy
	fmt.Println("= Creating identical slice then original")
	idx := make([]int, len(x))
	copied := copy(idx, x)
	fmt.Println("Copied identical amount as length, length of x=", len(x), "/copied=", copied)

	fmt.Println("- Creating smaller slice then original")
	smallx := make([]int, 4)
	copied = copy(smallx, x)
	fmt.Println("smaller slice: ", smallx)
	fmt.Println("Copied: ", copied)

	fmt.Println("+ Creating bigger slice then original")
	bigx := make([]int, len(x)+10)
	copied = copy(bigx, x)
	fmt.Println("new bigger slice: ", bigx)
	fmt.Println("Copied: ", copied)

	fmt.Println("+ Creating bigger slice then original copy just first first 3 (subslice)")
	biggerx := make([]int, len(x)+10)
	copied = copy(biggerx, x[:3])
	fmt.Println("new bigger slice: ", biggerx)
	fmt.Println("Copied: ", copied)

	/*
		  asovak@ASOVAK-M-924Q golang-basics % go run slices2.go
		Create a slice x: [1 2 3 4 5 6 7 8 9]
		x len= 9 , cap= 9
		= Creating identical slice then original
		Copied identical amount as length, length of x= 9 /copied= 9
		- Creating smaller slice then original
		smaller slice:  [1 2 3 4]
		Copied:  4
		+ Creating bigger slice then original
		new bigger slice:  [1 2 3 4 5 6 7 8 9 0 0 0 0 0 0 0 0 0 0]
		Copied:  9
		+ Creating bigger slice then original copy just first first 3 (subslice)
		new bigger slice:  [1 2 3 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
		Copied:  3
	*/
}
