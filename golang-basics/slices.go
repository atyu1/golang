package main

import (
	"fmt"
)

// This will create [1,0,3,0,5]
var globalX = []int{1, 3: 3, 5: 5}

func main() {
	// Init var
	var x []string

	fmt.Println("Print globalX: ", globalX)
	fmt.Println("Print 3rd item: ", globalX[2])

	fmt.Println("Is the x empty? ", x == nil) //reference by nil

	fmt.Println("Lenght of slice globalX: ", len(globalX))
	fmt.Println("Lenght of slice x: ", len(x))

	//Add elemenets
	fmt.Println("Adding values to the slice x...")
	x = append(x, "First string")
	x = append(x, "Second text", "Third item")

	fmt.Println("All elements in x now: ", globalX)

	//Capacity
	z := []int{1, 2, 3}
	fmt.Println("Getting capacity and lenght of z")
	fmt.Println("values, capacity, length")
	fmt.Println(z, len(z), cap(z))
	z = append(z, 4)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 5)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 6)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 7)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 8)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 9)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 10)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 11, 12, 13, 14, 15, 16, 17)
	fmt.Println(z, len(z), cap(z))

	//make
	numbers := make([]int, 5)
	fmt.Println("Getting capacity and lenght of numbers (maked for 5 elem)")
	fmt.Println("values, capacity, length")
	fmt.Println(numbers, len(numbers), cap(numbers))

	fmt.Println("Behaviour after append")
	numbers = append(numbers, 13)
	fmt.Println(numbers, len(numbers), cap(numbers))

	/// Append issues on sub-slices
	fmt.Println("Assigning slices are keeping capacity and append behaves strange")
	myx := make([]int, 0, 5)
	myx = append(myx, 1, 2, 3, 4)
	myy := myx[:2]
	myz := myx[2:]
	fmt.Println("cap myx, cap myy, cap myz")
	fmt.Println(cap(myx), cap(myy), cap(myz))
	myy = append(myy, 30, 40, 50)
	myx = append(myx, 60)
	myz = append(myz, 70)
	fmt.Println("cap myx, cap myy, cap myz")
	fmt.Println(cap(myx), cap(myy), cap(myz))
	fmt.Println("myx:", myx)
	fmt.Println("myy:", myy)
	fmt.Println("myz:", myz)

	/*
			golang-basics % go run slices.go
		Print globalX:  [1 0 0 3 0 5]
		Print 3rd item:  0
		Is the x empty?  true
		Lenght of slice globalX:  6
		Lenght of slice x:  0
		Adding values to the slice x...
		All elements in x now:  [1 0 0 3 0 5]
		Getting capacity and lenght of z
		values, capacity, length
		[1 2 3] 3 3
		[1 2 3 4] 4 6
		[1 2 3 4 5] 5 6
		[1 2 3 4 5 6] 6 6
		[1 2 3 4 5 6 7] 7 12
		[1 2 3 4 5 6 7 8] 8 12
		[1 2 3 4 5 6 7 8 9] 9 12
		[1 2 3 4 5 6 7 8 9 10] 10 12
		[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17] 17 24
		Getting capacity and lenght of numbers (maked for 5 elem)
		values, capacity, length
		[0 0 0 0 0] 5 5
		Behaviour after append
		[0 0 0 0 0 13] 6 10
		Assigning slices are keeping capacity and append behaves strange
		cap myx, cap myy, cap myz
		5 5 3
		cap myx, cap myy, cap myz
		5 5 3
		myx: [1 2 30 40 70]
		myy: [1 2 30 40 70]
		myz: [30 40 70]
	*/
}
