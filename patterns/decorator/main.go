package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan int)
	countWithLogger := withLogger(count)

	go countWithLogger(5, result)

	for res := range result {
		fmt.Println(res)
	}
}

type countFun func(int, chan int)

func withLogger(cb countFun) countFun {
	fmt.Println("log")
	return func(n int, c chan int) {
		cb(n, c)
	}
}

func count(max int, out chan int) {
	defer close(out)
	for n := 0; n <= max; n++ {
		time.Sleep(time.Millisecond * 200)
		out <- n
	}
}
