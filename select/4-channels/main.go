package main

import (
	"fmt"
	"time"
)

func main() {
	res := make(chan time.Duration)
	go slowFakeRequest(time.Now(), res)
	go fastFakeRequest(time.Now(), res)
	<-res
	<-res
}

func slowFakeRequest(now time.Time, res chan time.Duration) {
	duration := 4 * time.Second
	time.Sleep(duration)
	fmt.Printf("slow request take %ss\n", time.Since(now))
	res <- time.Since(now)
}

func fastFakeRequest(now time.Time, res chan time.Duration) {
	duration := 1 * time.Second
	time.Sleep(duration)
	fmt.Printf("fast request take %ss\n", time.Since(now))
	res <- time.Since(now)
}
