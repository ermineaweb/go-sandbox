package main

import (
	"fmt"
	"time"
)

type Job struct {
	id     int
	output bool
}

func main() {
	nWorkers := 4
	inputs := []int{3, 4, 5, 4, 3, 10, 39, 47, 54, 22, 23}
	nbInputs := len(inputs)
	inputChan := make(chan int, nbInputs)
	outputChan := make(chan int, nbInputs)

	fmt.Println("Workers count:", nWorkers)
	fmt.Println("Inputs count:", nbInputs)
	fmt.Println("-------------------------")

	fmt.Println("Create workers")
	for w := 1; w <= nWorkers; w++ {
		go worker(w, inputChan, outputChan)
	}

	fmt.Println("Send jobs to workers")
	for _, input := range inputs {
		inputChan <- input
	}
	close(inputChan)

	fmt.Println("Receive results")
	for range inputs {
		result := <-outputChan
		fmt.Println("<<< result", result)
	}
	close(outputChan)
}

func worker(id int, inputs chan int, results chan int) {
	for input := range inputs {
		fmt.Println("### worker", id, "with input", input)
		results <- calculate(input)
		// worker take a break
		time.Sleep(time.Millisecond * 1000)
	}
}

func calculate(value int) int {
	fmt.Println(">>> job in progress", value)
	// fake long time calculation
	time.Sleep(time.Millisecond * 4000)
	return value * 2
}
