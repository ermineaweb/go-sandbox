package main

import (
	"fmt"
	"time"
)

func main() {
	nWorkers := 4
	jobs := make(chan int, 10)
	results := make(chan bool, 10)

	// create workers
	for w := 1; w <= nWorkers; w++ {
		go worker(w, jobs, results)
	}

	// send jobs to workers
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	// receive results
	for r := 1; r <= 10; r++ {
		result := <-results
		fmt.Println("result", result)
	}
}

func worker(id int, jobs chan int, results chan bool) {
	for j := range jobs {
		fmt.Println("worker", id, "start", j)
		results <- calculate(j)
		fmt.Println("worker", id, "end")
		// worker take a break
		time.Sleep(time.Millisecond * 1000)
	}
}

func calculate(value int) bool {
	// long time calculation
	time.Sleep(time.Millisecond * 3000)
	return value < 5
}
