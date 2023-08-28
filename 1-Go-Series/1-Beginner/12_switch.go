package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("switch example")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of fice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1. Move 1 spot")

	case 2:
		// fallthorugh used then we want to
		// https://stackoverflow.com/questions/188461/switch-statement-fall-through-should-it-be-allowed
		fmt.Println("Move 2 spot")
		fallthorugh
	case 3:
		fmt.Println("Move 3 spot")

	case 4:
		fmt.Println("Move 4 spot")
	case 5:
		fmt.Println("Move 5 spot")
	case 6:
		fmt.Println("Move 6 spot")
	default:
		fmt.Println("what was the number?")
	}
}
