package main
import "log"
//var s string
var s = "seven"
func main () {
	var s2 = "six"
	log.Println("s is",s)
	log.Println("s2 is",s2)
	saySomething("xxx")
}

func saySomething(s string) (string, string) {
	log.Println("s from SaySomething",s)
	return s,"World"
}