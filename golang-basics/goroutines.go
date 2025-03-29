package main

import (
	"fmt"
	"runtime"
	"sync"
)

func print(x int) {
	fmt.Printf("#%v: %v\n", runtime.NumGoroutine(), x)
	wg.Done()
}

var wg sync.WaitGroup

func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go print(i)
	}
	wg.Wait()
}
