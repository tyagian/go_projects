package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("file example")
	content := "This needs to go in a file - example.com"
	file, err := os.Create("./myfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)

	checkNilErr(err)

	// fmt.Println("Text data inside the file is \n", databyte)
	// databyte by default will throw binary
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
