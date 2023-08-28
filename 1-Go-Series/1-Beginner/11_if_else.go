package main

import (
	"fmt"
)

func main() {

	fmt.Println("if else example")
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "result user"
	} else if loginCount > 10 {
		result = "count is 10"
	} else {
		result = "count is not enough"
	}
	fmt.Println(result)
	// if err != nil {

	// }
}
