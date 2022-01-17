package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sandbox/data"
)

type User struct {
	logger *log.Logger
}

func NewUser() User {
	return User{logger: log.New(os.Stdout, "user-api-", log.LstdFlags)}
}

func (u User) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	u.logger.Println("request on user handler")
	u.getUsers(rw, r)
}

func (u User) getUsers(rw http.ResponseWriter, r *http.Request) {
	u.logger.Println("getUsers")
	json := json.NewEncoder(rw)
	json.Encode(data.GetUsers)
}
