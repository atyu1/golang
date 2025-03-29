package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

func printName(index int, name string) {
	for i := 0; i < index; i++ {
		mutex.Lock()
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("#%d - Name: %s\n", i, name)
		mutex.Unlock()
	}
}

func main() {
	names := []string{"Peter", "Mark", "John", "Etelle", "Ferenc"}

	go printName(4, names[0])
	go printName(3, names[1])
	fmt.Scanln()
}
