package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Main Start")
	cmd := exec.Command("ldd", "--version")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Printf("GLIBC version: %s", output)
}
