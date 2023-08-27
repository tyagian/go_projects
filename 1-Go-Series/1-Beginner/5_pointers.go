package main

import "fmt"

/*
Whenever we declare any var, they are reference of memory.
When we store their value, they get store in a memory.
Problem is that when sometimes you pass on these variables on variety of functions and classes,
these variables don't get passed along but copy get passed along.
This create irrregularity in the program, to avoid this we pointers.

Pointers give surety, we are passing memory address of the variable,
which guarantee actual values passed on.

Also https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac

*/
func main() {

	fmt.Println("Welcome to a class on pointers")
	// var one int
	//var ptr *int
	//fmt.Println("Value of pointer is", ptr)
	// Output: Value of pointer is <nil>
	// so initial value of pointer is nil

	mynumber := 23
	var ptr = &mynumber
	fmt.Println("Value of pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)
	*ptr = *ptr + 2
	fmt.Println("Value of pointer is", mynumber)
}
