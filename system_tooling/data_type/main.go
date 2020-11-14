package main

import (
	"fmt"
	"system_tooling/data_type/message"
)

func main() {

	message := message.greeting("Keith", "Hello")
	fmt.Println(message)
}
