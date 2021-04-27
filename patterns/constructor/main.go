package main

import (
	"fmt"
	"strconv"

	"github.com/ermineaweb/sandbox/patterns/constructor/user"
)

func main() {
	harry := user.NewUser("harry")
	fmt.Println("the best sorcer is " + harry.GetUsername())
	fmt.Println("his ID is " + strconv.Itoa(harry.GetId()))
}
