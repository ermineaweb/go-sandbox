package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

	// parallel works
	go func() { ch4 <- work4() }()
	go func() { ch2 <- work2() }()
	go func() { ch1 <- work1() }()
	go func() { ch3 <- work3() }()

	// ordered results
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	fmt.Println(<-ch3)
	fmt.Println(<-ch4)

	fmt.Printf("\ntime elapsed %v\n", time.Since(start).Seconds())
}

func work1() string { time.Sleep(4 * time.Second); return "work1" }
func work2() string { time.Sleep(2 * time.Second); return "work2" }
func work3() string { time.Sleep(1 * time.Second); return "work3" }
func work4() string { time.Sleep(3 * time.Second); return "work4" }
