package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - github.com"

	file, err := os.Create("./mylocalfile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is:", length)
	// if you don't to do more after adding content to the file
	// so we defer closing file
	defer file.Close()
	readFile("./mylocalfile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)

	checkNilError(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

// To avoid repeatition of writing if err != nil condition
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
