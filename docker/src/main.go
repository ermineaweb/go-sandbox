package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	err := HttpError{status: 404, message: "not found"}
	fmt.Println(err.Error())
}

type HttpError struct {
	status  int
	message string
}

func (err *HttpError) Error() string {
	return fmt.Sprintf("%d %s", err.status, err.message)
}
