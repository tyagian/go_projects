package main

import "fmt"

/*
func main() {

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
} */

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)

	}
}

// output: 43210
