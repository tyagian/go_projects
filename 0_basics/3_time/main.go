package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Parsing time from time package")
	presentTime := time.Now()

	// Parse the time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	//fmt.Println(presentTime.Date())
	// createdDate := time.Date(2020, time.August, 10,23,23,0,0, time.UTC())
	// fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
