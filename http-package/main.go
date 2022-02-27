package main

import (
	"go-sandbox/http-package/handlers"
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.Handle("/p", handlers.NewProduct())
	server.Handle("/u", handlers.NewUser())

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatal("failed to start server")
	}
}
