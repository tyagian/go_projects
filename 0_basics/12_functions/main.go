package main

import "fmt"

// main works because it act as entrypoint for program
func main() {
	fmt.Println("Welcome to functions in golang")
	greeting()

	greeterTwo()

	result := adder(3, 5)
	fmt.Println(result)

	fmt.Println(proAdder(3, 5, 6, 5))
}

// signatures define type of value to accept and return
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func greeting() {
	fmt.Println("Namastey from golang")
}

func greeterTwo() {
	fmt.Println("Another method")
}
