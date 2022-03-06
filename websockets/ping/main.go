package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const ping = 5 * time.Second

var upgrader = websocket.Upgrader{}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/ws", serveWebsocket)
	fmt.Println("server start")
	http.ListenAndServe(":5000", nil)
}

func serveWebsocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("client connected")

	// upgrade http connection to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// ping client
	ticker := time.NewTicker(ping)
	go func() {
		for tick := range ticker.C {
			fmt.Println("ping", tick.Local())
			if err := conn.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				return
			}
		}
	}()
}
