package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in Golang")
	// no inheritance in golang, no super or parent
	anuj := User{"anuj", "anuj@asda.casad", true, 6}
	//fmt.Println(anuj)
	fmt.Printf("anuj details are: %+v\n", anuj)
	fmt.Printf("Name is %v and email is %v.", anuj.Name, anuj.Email)

}

// User as Capital U because it's like Class name in other language so use it
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
