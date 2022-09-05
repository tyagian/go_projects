package main

import "log"

// func main() {

/** maps
myMap := make(map[string]string)
myMap["dog"] = "Samson"
myMap["other-dog"] = "Cassie"
myMap["dog"] = "fido"

log.Println(myMap["dog"])
log.Println(myMap["other-dog"]) */
/* myMap := make(map[string]int)
myMap["First"] = 1
myMap["Second"] = 2
log.Println(myMap["First"],myMap["Second"]) */

/*
	type User struct {
		FirstName string
		LastName string
	}

	func main() {
		myMap := make(map[string]User)
		me := User {
			FirstName: "Trevor",
			LastName: "Sawler",
		}
		myMap["me"] = me
		log.Println(myMap["me"].FirstName)
	}
*/
/* Slices
func main() {
	var mySlice []string
	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "Trevor")
	log.Println(mySlice)
} */
// slice of string
func main() {
	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names[1:3])
}
