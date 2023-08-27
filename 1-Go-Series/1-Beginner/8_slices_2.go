// remove slice based on index
package main

import (
	"fmt"
)

func main() {
	var courses = []string{"react", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
