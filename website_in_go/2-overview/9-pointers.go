package main
import "log"

func main() {
	var myString string
	myString = "Green"
	//changeUsingPointer(myString)
	log.Println("myString is set to",myString)
	changeUsingPointer(&myString)
	log.Println("After func call", myString)
}

func changeUsingPointer(s *string) {
	newval := "Red"
	*s = newval
}