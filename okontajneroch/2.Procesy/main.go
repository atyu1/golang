package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {

	pid, _ := syscall.ForkExec("/usr/bin/ping", []string{"ping", "-c", "10", "localhost"}, nil)

	// waiting for child process
	for {
		var s syscall.WaitStatus
		wpid, err := syscall.Wait4(pid, &s, syscall.WNOHANG, nil)
		if err != nil || wpid != 0 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	// child process is dead, we can continue
	fmt.Println("The end!")
}
