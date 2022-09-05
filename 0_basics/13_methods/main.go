package main

import "fmt"

func main() {

	fmt.Println("Structs in Golang")

	// no inheritance in golang, no Super or parent

	anuj := User{"anuj", "anuj@go.dev", true, 16}
	fmt.Println(anuj)

	fmt.Printf("anuj details are: %+v\n", anuj)
	fmt.Printf("Name is %v and email is %v. \n", anuj.Name, anuj.Email)
	anuj.GetStatus()
	anuj.NewMail()
}

// User starting with upper case since it's like a Class and needs to be imported
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
