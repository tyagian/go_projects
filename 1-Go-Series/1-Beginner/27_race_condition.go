package main

import (
	"fmt"
	"sync"
)

/*
How to check if code has race condition
 go run --race 27_race_condition.go
Race condition
One R
Two R
Three R
==================
WARNING: DATA RACE
Read at 0x00c000118000 by goroutine 7:
  main.main.func2()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:22 +0x78
  main.main.func5()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:24 +0x44

Previous write at 0x00c000118000 by goroutine 6:
  main.main.func1()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:16 +0xec
  main.main.func4()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:18 +0x44

Goroutine 7 (running) created at:
  main.main()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:20 +0x2d4

Goroutine 6 (finished) created at:
  main.main()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:14 +0x1d4
==================
==================
WARNING: DATA RACE
Read at 0x00c00000e0d8 by goroutine 8:
  runtime.growslice()
      /usr/local/go/src/runtime/slice.go:157 +0x0
  main.main.func3()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:27 +0xb0
  main.main.func6()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:29 +0x44

Previous write at 0x00c00000e0d8 by goroutine 6:
  main.main.func1()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:16 +0xd4
  main.main.func4()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:18 +0x44

Goroutine 8 (running) created at:
  main.main()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:25 +0x3d4

Goroutine 6 (finished) created at:
  main.main()
      /Users/swatityagi/go/src/go_projects/1-Go-Series/1-Beginner/27_race_condition.go:14 +0x1d4
==================
Found 2 data race(s)
exit status 66

*/
func main() {
	println("Race condition")

	// mutex added to fix race condition
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	mut.Rlock()
	var score = []int{0}
	mut.RUnlock()
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	//wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
}
