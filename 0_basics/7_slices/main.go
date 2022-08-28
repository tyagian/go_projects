package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Testing slices in this tutorial")
	var FruitList = []string{}
	fmt.Printf("Type of data is %T\n", FruitList)

	FruitList = append(FruitList, "Mango", "Banana")
	fmt.Println(FruitList)

	FruitList = append(FruitList[1:])
	fmt.Println(FruitList)

	//fmt.Println(FruitList[1:3])
	/* panic: runtime error: slice bounds out of range [:3] with capacity 2

	goroutine 1 [running]:
	main.main()
			/Users/swatityagi/go/src/go_projects/0_basics/7_slices/main.go:14 +0x200
	exit status 2
	*/

	// initialize the data by defining type with it
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 432
	highScores[2] = 345
	highScores[3] = 113

	fmt.Println(highScores)
	// we can add more data to that
	highScores = append(highScores, 555, 345, 321)
	fmt.Println(highScores)

	// check if integers are sorted or not
	fmt.Println(sort.IntsAreSorted(highScores))

	// sort integers
	sort.Ints(highScores)
	fmt.Println(highScores)

	// check if integers are sorted or not
	fmt.Println(sort.IntsAreSorted(highScores))

	// reverse sort
	sort.Sort(sort.Reverse(sort.IntSlice(highScores)))
	fmt.Println(highScores)

	// How to remove values from slices
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
