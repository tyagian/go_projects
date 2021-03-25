package main
import (
	"time"
	"log"
)

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age			int
	BirthDate	time.Time
}

func main() {
	user := User {
		FirstName: "Trevor",
		LastName: "Sawler",
	}
	log.Println(user.FirstName, user.LastName,"BirthDate:", user.BirthDate)
}