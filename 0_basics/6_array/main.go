package main

import "fmt"

func main() {

	fmt.Println("Welcome to array in Golang")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Tomato"

	fmt.Println("List of fruit is:", fruitList)
	fmt.Println("List of fruit is:", len(fruitList))

	/*
		go run main.go
		Welcome to array in Golang
		List of fruit is: [Apple Orange  Tomato]
	*/

	var vegList = [5]string{"potato", "beans", "mushroom"}

	fmt.Println("List of veggie is:", vegList)
	fmt.Println("List of veggie is:", len(vegList))

}
