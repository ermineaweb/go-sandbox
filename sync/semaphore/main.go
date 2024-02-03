package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	var semaphore = make(chan bool, 2)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- true
			worker(i)
			<-semaphore
		}(i)
	}

	wg.Wait()
	fmt.Printf("\ntime elapsed %v\n", time.Since(start).Seconds())
}

func worker(id int) {
	fmt.Printf("worker %v start\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("worker %v stop\n", id)
}
