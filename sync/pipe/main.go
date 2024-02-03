package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	counts := make(chan int)
	squares := make(chan int)
	negatives := make(chan int)

	go counter(counts)
	go squarer(counts, squares)
	go negativer(squares, negatives)
	printer(negatives)

	fmt.Println(time.Since(start).Seconds())
}

func counter(out chan<- int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("counter %v\n", i)
		time.Sleep(1 * time.Second)
		out <- i
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		fmt.Printf("squarer %v\n", x)
		time.Sleep(2 * time.Second)
		out <- x * x
	}
	close(out)
}

func negativer(in <-chan int, out chan<- int) {
	for x := range in {
		fmt.Printf("negativer %v\n", x)
		time.Sleep(3 * time.Second)
		out <- -x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		time.Sleep(1 * time.Second)
		fmt.Println(x)
	}
}
