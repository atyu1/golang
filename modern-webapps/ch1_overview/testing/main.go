package main

import (
	"errors"
	"log"
)

func divide(x, y float32) (float32, error) {
	var res float32
	if y == 0 {
		return res, errors.New("Cannot divide by 0")
	}
	res = x / y
	return res, nil
}

func main() {
	value, err := divide(100.0, 10.0)
	if err != nil {
		log.Println("Error with division:", err)
		return
	}
	log.Println(value)
}
