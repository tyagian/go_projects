package main

import "fmt"

type Demo struct {
	Counter int
}

func main() {
	d := Demo{Counter: 0}
	fmt.Println("Before ptr inc", d.Counter)
	fmt.Println("IncPtr return:", d.IncPtr())
	fmt.Println("After ptr inc", d.Counter)
	fmt.Println("Before val inc", d.Counter)
	fmt.Println("IncVal return:", d.IncVal())
	fmt.Println("After val inc:", d.Counter)
}

func (d *Demo) IncPtr() int {
	d.Counter++
	return d.Counter
}

func (d Demo) IncVal() int {
	d.Counter++
	return d.Counter
}
