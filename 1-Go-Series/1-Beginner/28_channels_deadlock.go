package main

import (
	"fmt"
	"sync"
)

/*
go run 28_channels_deadlock.go
Channels and deadlock
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/28_channels_deadlock.go:11 +0x50
exit status 2
*/

func main() {
	println("Channels and deadlock")
	myCh := make(chan int, 5) // 5 is buffer channel here
	wg := &sync.WaitGroup{}
	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	// Read only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		close(myCh)
		// myCh <- 5
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
