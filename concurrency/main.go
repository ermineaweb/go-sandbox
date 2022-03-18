package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		for {
			chan1 <- "2000ms"
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	go func() {
		for {
			chan2 <- "500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		select {
		case c := <-chan1:
			fmt.Println(c)
		case c := <-chan2:
			fmt.Println(c)
		}
	}
}
