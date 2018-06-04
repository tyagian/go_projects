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
		numSlice := make([]int, 5, 10)

	fmt.Println(numSlice)
	numSlice1 := append(numSlice, 0, -1)
	fmt.Println(numSlice1)

	Function: return value type with function defination
	func main() {
	//favNums := [5]float64{1, 2, 3, 5}
	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)
	}

	func next2Values(number int) (int, int) {
		//favNums := [5]float64{1, 2, 3, 5}
		return number + 1, number + 2
	}
	#when receiving n number of arguments in function:
	func subtractThem(args ...int) int {
	//favNums := [5]float64{1, 2, 3, 5}
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}
	return finalValue
	## function inside function: Closure
	## defer function: Calling a function with defer will execute is after main() after or at the end even if we call it in the begining in main() function
	## call by value
	## call by refernce:
	func main() {
	//favNums := [5]float64{1, 2, 3, 5}
	x := 0
	changeXValNow(&x)
	fmt.Println("x=", x)
	fmt.Println("Memory Address for x=", &x)
	}
	func changeXValNow(x *int) {
	*x = 2
	}
	## pointers:
	func main() {
	yPtr := new(int)
	changeYValNow(yPtr)
	fmt.Println("y=", *yPtr)
	}
	func changeYValNow(yPtr *int) {
	*yPtr = 2
	}
	## struct:

	func main() {
		//favNums := [5]float64{1, 2, 3, 5}
		rect1 := Rectangle{0, 50, 10, 10}
		//fmt.Println("Rectangle is", rect1.width, "wide")
		fmt.Println("Area of rectangle is", rect1.area())
	}

	type Rectangle struct {
		leftX  float64
		topY   float64
		height float64
		width  float64
	}

	func (rect *Rectangle) area() float64 {
		return rect.width * rect.height
	}

	## Interfaces: .... all tied to shape() interface here
	func main() {
		//favNums := [5]float64{1, 2, 3, 5}
		rect := Rectangle{20, 50}
		cir := Circle{4}
		fmt.Println("Rectangle Area=", getArea(rect))
		fmt.Println("Circle Area=", getArea(cir))
		//fmt.Println("Rectangle is", rect1.width, "wide")
		//fmt.Println("Area of rectangle is", rect1.area())
	}

	type Shape interface {
		area() float64
	}

	type Rectangle struct {
		height float64
		width  float64
	}

	type Circle struct {
		radius float64
	}

	func (r Rectangle) area() float64 {
		return r.height * r.width
	}

	func (c Circle) area() float64 {
		return math.Pi * math.Pow(c.radius, 2)
	}

	func getArea(shape Shape) float64 {
		return shape.area()
	}

	####### string func, files I/o, accepting user input, type casting, goroutines, channels

	## String func:

	func main() {

	file, err := os.Create("samp.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is some random text")
	file.Close()
	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)
	//csvString := "1,2,3,4,5,6"
	//fmt.Println(strings.Split(csvString, ","))
	//listOfLetters := []string{"c","a","b"}
	//sort.Strings(listOfLetters)
	//fmt.Println("Letters:", listOfLetters)
	//listOfNums := strings.Join([]string{"3","2","1"}, ", ");
	//fmt.Println(listOfNums)
	//sampString := "Hello World"
	//fmt.Println(strings.Contains(sampString, "lo"))
	//fmt.Println(strings.Index(sampString, "lo"))
	//fmt.Println(strings.Count(sampString, "l"))
	//fmt.Println(strings.Replace(sampString, "l","x",3))

}
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file, err := os.Create("samp.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is some random text")
	file.Close()
	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)
	//csvString := "1,2,3,4,5,6"
	//fmt.Println(strings.Split(csvString, ","))
	//listOfLetters := []string{"c","a","b"}
	//sort.Strings(listOfLetters)
	//fmt.Println("Letters:", listOfLetters)
	//listOfNums := strings.Join([]string{"3","2","1"}, ", ");
	//fmt.Println(listOfNums)
	//sampString := "Hello World"
	//fmt.Println(strings.Contains(sampString, "lo"))
	//fmt.Println(strings.Index(sampString, "lo"))
	//fmt.Println(strings.Count(sampString, "l"))
	//fmt.Println(strings.Replace(sampString, "l","x",3))

}
