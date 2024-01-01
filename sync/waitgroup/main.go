package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	var count = 4
	wg.Add(count)

	for v := 1; v <= count; v++ {
		// wg.Add(1)
		go func(number int) {
			fmt.Printf("result %v\n", longFunction(number))
			wg.Done()
		}(v)
	}
	wg.Wait()

	fmt.Printf("time elapsed: %vs\n", time.Since(start).Seconds())
}

func longFunction(number int) int {
	fmt.Printf("-work start %v\n", number)
	time.Sleep(time.Second * 2)
	fmt.Printf("-work end %v\n", number)
	return number
}
