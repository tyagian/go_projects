package main

import (
	"fmt"
)

func myfunction(firstName string, lastName string) (string, error) {
	return firstName + " " + lastName, nil
}

func main() {
	fmt.Println("Hello World")

	fullName, err := myfunction("Elliot", "Forbes")
	if err != nil {
		fmt.Println("Handle Error Case")
	}
	fmt.Println(fullName)
}
