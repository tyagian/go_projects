package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	greeter()

	result := adder(3, 4)
	fmt.Println("Result is ", result)
	proresult, ProMsg := proAdder(3, 4, 5, 6, 6)
	fmt.Println("Pro Result is ", proresult, ProMsg)

}

// take 2 values and return int type
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// n number of input values
// ...int means all values will be of type integer
func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, " value"
}

func greeter() {
	fmt.Println("Namastey from golang")
}
