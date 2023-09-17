package main

import (
	"fmt"
	"os"
)

func main() {
	//var args []st ring
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: ./hello world argument \n")
		os.Exit(1)
	}
	fmt.Printf("hello world\nArguments:  %v\n", args[1:])
}
