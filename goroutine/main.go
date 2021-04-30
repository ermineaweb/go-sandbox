package main

import "fmt"

func main() {

	c := make(chan int)
	go count(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("end")

}

func count(c chan int) {
	for n := 0; n < 5; n++ {
		c <- n
	}
}
