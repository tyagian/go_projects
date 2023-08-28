package main

import (
	"fmt"
)

func main() {
	// defer move function with defer at the end
	defer fmt.Println("Hello")
	fmt.Println("World")
	// output:
	// World
	// Hello

	// but if there are multiple defer functions,
	// they work with last in first out
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("World")
	// output:
	// World
	// Three
	// Two
	// One
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
