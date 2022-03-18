package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go action("pecher", channel)
	go action("chasser", channel)
	for msg := range channel {
		fmt.Println(msg)
	}
}

func action(action string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- action
		time.Sleep(500 * time.Millisecond)
	}
	close(channel)
}
