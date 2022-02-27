package main

import (
	"fmt"

	"go-sandbox/patterns/di/user"
)

func main() {
	userService := user.Service{Repo: user.Repository{}}
	fmt.Println(userService.GetAll())
	id := userService.Add(user.User{Username: "Dumbledore"})
	fmt.Println(userService.GetAll())
	fmt.Println(userService.GetById(id))
}
