package main

import (
	"fmt"
	"time"
)

type Input struct {
	str string
	ch  chan bool
}

func main() {
	start := time.Now()

	// ordered inputs
	inputs := []Input{
		{str: "work1", ch: make(chan bool)},
		{str: "work2", ch: make(chan bool)},
		{str: "work3", ch: make(chan bool)},
		{str: "work4", ch: make(chan bool)},
	}

	// parallel works
	for _, in := range inputs {
		go func(in Input) {
			in.str = work(in.str)
			in.ch <- true
		}(in)
	}

	// ordered results
	for _, in := range inputs {
		<-in.ch
		fmt.Println(in.str)
	}

	fmt.Printf("\ntime elapsed %v\n", time.Since(start).Seconds())
}

func work(input string) string { time.Sleep(3 * time.Second); return fmt.Sprintf("### %s", input) }
