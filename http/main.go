package main

import (
	"go-sandbox/http/handlers"
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.Handle("/product", handlers.NewProduct())
	server.Handle("/user", handlers.NewUser())

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatal("failed to start server")
	}
}
