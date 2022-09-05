package main
import "log"

type mystruct struct {
	FirstName string
}

func (m *mystruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar mystruct
	myVar.FirstName = "John"

	myVar2 := mystruct {
		FirstName: "Mary",
	}
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar is set to",myVar2.printFirstName())
}