package main

import "fmt"

func main() {

	fmt.Println("Structs in Golang")

	// no inheritance in golang, no Super or parent

	anuj := User{"anuj", "anuj@go.dev", true, 16}
	fmt.Println(anuj)

	fmt.Printf("anuj details are: %+v\n", anuj)
	fmt.Printf("Name is %v and email is %v. \n", anuj.Name, anuj.Email)
}

// User starting with upper case since it's like a Class and needs to be imported
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
