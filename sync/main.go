package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for v := 0; v <= 20; v++ {
		wg.Add(1)
		// asynchronous
		go func(number int) {
			fmt.Println(longFunction(number))
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func longFunction(number int) int {
	defer fmt.Println("long time function end")
	fmt.Println("long time function start")
	time.Sleep(time.Millisecond * 2000)
	return number
}
