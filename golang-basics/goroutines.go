package main

import (
	"fmt"
	"runtime"
	"sync"
)

func print(x int) {
	fmt.Printf("#%v: %v\n", runtime.NumGoroutine(), x)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 100; i++ {
		go print(i)
		wg.Done()
	}
	wg.Wait()
}
