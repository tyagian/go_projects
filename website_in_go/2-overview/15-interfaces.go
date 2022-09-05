package main

import "log"

type Animal interface {
	Says() string
	NumberofLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberofTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepard",
	}
	Println(dog)
}
func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberofLegs() int {
	return 4
}
func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberofLegs(), "legs)")
}
