package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// zmodifikujem si prompt, aby som vedel rozlíšiť
	// či som v hosťovskom systéme alebo v behovom protredí
	os.Setenv("PS1", "anton> ")

	// pripravim si podproces a presmerujem
	// štandardné vstupy a výstupy
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// nový podproces bude mať vlastné menný
	// priestor pre použivateľa
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUSER | syscall.CLONE_NEWPID,

		// //mapovanie použivateľa
		// UidMappings: []syscall.SysProcIDMap{
		// 	{
		// 		ContainerID: 0,
		// 		HostID:      os.Getuid(),
		// 		Size:        1,
		// 	},
		// },

		// // mapovanie skupiny
		// GidMappings: []syscall.SysProcIDMap{
		// 	{
		// 		ContainerID: 0,
		// 		HostID:      os.Getuid(),
		// 		Size:        1,
		// 	},
		// },
	}

	// sputenie podprocesu
	cmd.Run()
}
