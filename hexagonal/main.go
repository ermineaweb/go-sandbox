package main

import (
	"fmt"
	"go-sandbox/hexagonal/repository/memory"
	"go-sandbox/hexagonal/shortener"
)

func main() {
	shortenerRepo := memory.Repository{}
	shortenerService := shortener.NewService(shortenerRepo)

	fmt.Println(shortenerService.FindAll())

	shortUrl, err := shortenerService.Create("badUrlFormat")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(shortenerService.FindAll())

	shortUrl, err = shortenerService.Create("http://reddit.com")
	fmt.Println(shortenerService.FindAll())

	url, err := shortenerService.Find("inexistantCode")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(url.Url)

	url, err = shortenerService.Find(shortUrl.Code)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(url.Url)

}
