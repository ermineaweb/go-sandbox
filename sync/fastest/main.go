package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("winner: %v\n", quickestSorcerer())
	fmt.Printf("\ntime elapsed %v\n", time.Since(start).Seconds())
}

func quickestSorcerer() string {
	results := make(chan string)
	go func() { results <- launchSpell("hermione") }()
	go func() { results <- launchSpell("harry") }()
	go func() { results <- launchSpell("ron") }()
	return <-results
}

func launchSpell(name string) string {
	rnd := rand.Intn(4)
	fmt.Printf("%v launch a spell, it take %vs\n", name, rnd)
	time.Sleep(time.Duration(rnd) * time.Second)
	return name
}
