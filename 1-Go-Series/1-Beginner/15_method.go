package main

import (
	"fmt"
)

// Note: Go methods are similar to Go function with one difference,
// i.e, the method contains a receiver argument in it.
// With the help of the receiver argument, the method can access the
// properties of the receiver. Here, the receiver can be of struct
// type or non-struct type.
// https://www.geeksforgeeks.org/methods-in-golang/

func main() {
	fmt.Println("Structs in Golang")
	// no inheritance in golang, no super or parent
	anuj := User{"anuj", "anuj@asda.casad", true, 6}
	//fmt.Println(anuj)
	fmt.Printf("anuj details are: %+v\n", anuj)
	fmt.Printf("Name is %v and email is %v.\n", anuj.Name, anuj.Email)
	anuj.GetStatus()
	anuj.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", anuj.Name, anuj.Email)

}

// User as Capital U because it's like Class name in other language so use it
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int
}

// example of method
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

// another method
func (u User) NewMail() {
	u.Email = "test.go.dev"
	fmt.Println("Email of user is : ", u.Email)
}
