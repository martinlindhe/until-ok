package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	for {
		err := runInteractiveCommand(os.Args[1], os.Args[2:]...)
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(3 * time.Second)
	}
}

func runInteractiveCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
