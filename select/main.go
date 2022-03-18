package main

import (
	"fmt"
	"time"
)

func main() {
	resSlow := make(chan time.Duration)
	resFast := make(chan time.Duration)

	// execute 3 times each request
	for i := 1; i <= 3; i++ {
		go slowFakeRequest(i, time.Now(), resSlow)
		go fastFakeRequest(i, time.Now(), resFast)
	}

	for i := 1; i <= 6; i++ {
		select {
		case res1 := <-resSlow:
			fmt.Printf("received slow result %s", res1)
		case res2 := <-resFast:
			fmt.Printf("received fast result %s", res2)
		}
	}
}

func slowFakeRequest(requestCount int, now time.Time, res chan time.Duration) {
	duration := 4 * time.Second
	time.Sleep(duration)
	fmt.Printf("%d slow request take %ss\n", requestCount, time.Since(now))
	res <- time.Since(now)
}

func fastFakeRequest(requestCount int, now time.Time, res chan time.Duration) {
	duration := 1 * time.Second
	time.Sleep(duration)
	fmt.Printf("%d fast request take %ss\n", requestCount, time.Since(now))
	res <- time.Since(now)
}
