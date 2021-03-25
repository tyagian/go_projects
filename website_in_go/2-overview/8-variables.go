package main
import "log"
import "fmt"
func main() {
	var whatToSay string
	var saySomethingElse string
	var i int
	//whatToSay,_ := saySomething("Hello")
	//log.Println(whatToSay)
	whatToSay, _ = saySomething("Hello")
	fmt.Println(saySomething("Hellozzz"))
	
	fmt.Println(whatToSay)
	saySomethingElse, _ = saySomething("GoodBye")
	log.Println(saySomethingElse)
	log.Println(saySomething("Finally"))
	i = 7
	i = 8
	log.Println(i)	
}
func saySomething(s string) (string, string) {
	return s, "World"
}
