package main

import (
	"fmt"
	"time"
)

func main() {
	slowFakeRequest(time.Now())
	fastFakeRequest(time.Now())
}

func slowFakeRequest(now time.Time) {
	duration := 4 * time.Second
	time.Sleep(duration)
	fmt.Printf("slow request take %ss\n", time.Since(now))
}

func fastFakeRequest(now time.Time) {
	duration := 1 * time.Second
	time.Sleep(duration)
	fmt.Printf("fast request take %ss\n", time.Since(now))
}
