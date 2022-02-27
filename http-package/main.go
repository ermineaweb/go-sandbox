package main

import (
	"log"
	"net/http"
	"sandbox/handlers"
)

func main() {
	server := http.NewServeMux()
	server.Handle("/p", handlers.NewProduct())
	server.Handle("/u", handlers.NewUser())

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatal("failed to start server")
	}
}
