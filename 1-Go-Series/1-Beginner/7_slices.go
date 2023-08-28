//
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slice")
	var fruitList = []string{"apple", "avacado", "peach"}
	// slice when length is not defined
	fmt.Println("type of fruitlist is %T\n", fruitList)

	// we don't add values in slice like array. We use append() here to add
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	fmt.Println(fruitList[1:3])

	// make() is used for slices, maps, or channels.
	// make() allocates memory on the heap and initializes and
	// puts zero or empty strings into values. Unlike new() ,
	// make() returns the same type as its argument.
	// Slice: The size determines the length.

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 454
	highScores[2] = 123
	highScores[3] = 541
	//highScores[4] = 675
	// if you add 5th element in size of [4] slice, we will get error:
	// panic: runtime error: index out of range [4] with length 4
	fmt.Println(highScores)

	// adding more elements
	highScores = append(highScores, 455, 555, 122)
	fmt.Println(highScores)
	// sort method() is available in slices but not in array
	sort.Ints(highScores)
	fmt.Println(highScores)

	// true if integer values are sorted
	fmt.Println(sort.IntsAreSorted(highScores))
}
