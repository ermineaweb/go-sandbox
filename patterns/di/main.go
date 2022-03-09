package main

import (
	"fmt"

	"go-sandbox/patterns/di/user"
)

func main() {
	userService := user.Service{Repo: user.NewUserRepository()}
	fmt.Println(userService.GetAll())
	id := userService.Add(user.User{Username: "Dumbledore"})
	fmt.Println(userService.GetAll())
	fmt.Println(userService.GetById(id))

	fakeUserService := user.Service{Repo: user.NewMockRepository()}
	fmt.Println(fakeUserService.GetAll())
	fakeId := fakeUserService.Add(user.User{Username: "Marge"})
	fmt.Println(fakeUserService.GetAll())
	fmt.Println(fakeUserService.GetById(fakeId))
}
