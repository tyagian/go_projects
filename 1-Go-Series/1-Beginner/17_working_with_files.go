package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("file example")
	content := "This needs to go in a file - example.com"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is", length)
	defer file.Close()
}
