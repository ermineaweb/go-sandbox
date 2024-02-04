package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	Retry(5, 3, RetryWrapper(DummyFunc))
}

func Retry(nbRetry int, delaySecond int, fn func() error) error {
	for i := 1; i <= nbRetry; i++ {
		err := fn()
		if err == nil {
			return nil
		}
		fmt.Println(err)
		fmt.Printf("retry...%v\n", i)
		time.Sleep(time.Duration(delaySecond) * time.Second)
	}
	return nil
}

func RetryWrapper(a func() (string, error)) func() error {
	return func() error {
		str, err := a()
		if err != nil {
			return err
		}
		fmt.Println(str)
		return nil
	}
}

var count = 0

func DummyFunc() (string, error) {
	if count < 3 {
		count++
		return "", errors.New("error")
	}
	return "Hello world", nil
}
