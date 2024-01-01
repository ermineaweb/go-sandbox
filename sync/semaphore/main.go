package main

import (
	"fmt"
	"sync"
	"time"
)

var semaphore = make(chan struct{}, 2)
var wg sync.WaitGroup

func main() {
	start := time.Now()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()

	fmt.Printf("\ntotal time %vs\n", time.Since(start).Seconds())
}

func worker(id int) {
	semaphore <- struct{}{}
	fmt.Printf("worker %v start\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("worker %v stop\n", id)
	<-semaphore
	wg.Done()
}
