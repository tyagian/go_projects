package main

import (
	"fmt"
)

func main() {

	fmt.Println("switch example")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Saturday"}
	//fmt.Println(days)

	// one approach. Not used much
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// more used approach
	for i := range days {
		fmt.Println(days[i])
	}
	// same output in both above loops

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	// if we don't care about index
	for _, day := range days {
		fmt.Printf("value is %v\n", day)
	}

	// break to come out of loop
	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 5 {
			break
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	// continue to skip values in a loop
	rougueVal := 1
	for rougueVal < 10 {
		if rougueVal == 5 {
			rougueVal++
			continue
		}
		fmt.Println("Value is: ", rougueVal)
		rougueVal++
	}

	// goto example
	// if condition match then jump to "goto"
	rougueVl := 1
	for rougueVl < 10 {
		if rougueVl == 2 {
			goto lco

		}
		fmt.Println("Value is: ", rougueVl)
		rougueVl++
	}
lco:
	fmt.Println("value is 2")
}
