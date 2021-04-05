package main

import "fmt"

// https://play.golang.org/p/Ol1STUckhrm

func main() {
	var s string = "Hello, "
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs) //Utf8 bytes
	fmt.Println(rs) // Runes
}
