package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Create("new_file.txt")
	if err != nil {
		log.Fatal("Cannot access files", err)
	}
	defer file.Close()
	fmt.Fprintf(file, "Thanks for reading this!")
}
