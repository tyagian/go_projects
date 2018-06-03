// basic tutorial by Derek Banas
/*  multi-line comment
1. go run: go run
2. godoc fmt println
3. var age int = 40
// loop
	favNums3 := [5]float64{1, 2, 3, 4, 5}
	for i, value := range favNums3 {
		fmt.Println(value, i)
	}
	or
	favNums[5] float64
	favNums[0] = 163
	favNums[1] = 982
	favNums[2] = 159
	favNums[3] = 112
	favNums[4] = 15  Note if you don't define arroay[n] and print it will give 0

	Slices: are like array but when you define them, you leave out the size
*/
package main

import (
	"fmt"
)

func main() {
	//favNums := [5]float64{1, 2, 3, 5}
	var favNums [5]float64
	favNums[0] = 163
	favNums[1] = 982
	favNums[2] = 159
	//favNums[3] = 112
	favNums[4] = 154
	for i, value := range favNums {
		fmt.Println(value, i)
	}
}
