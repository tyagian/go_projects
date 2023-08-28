package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")
	var fruitList [5]string
	fruitList[0] = "apple"
	fruitList[1] = "orange"
	fruitList[2] = "banana"
	fruitList[3] = "cherry"
	fmt.Println("value of array fruits", fruitList)

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("veggy list is:", vegList, len(vegList))
}
