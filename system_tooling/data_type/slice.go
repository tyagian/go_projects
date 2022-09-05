package main

import "fmt"

func main() {

	//names := [3]string{"keith", "kevin"}
	names := []string{}
	fmt.Println(names)

	names[0] = "Parker"
	names[1] = "Django"
	fmt.Println(names)
	fmt.Println("names[2] is nil:", names[2] == "")
	names = append(names, "ABC")
	fmt.Println(names)
}
