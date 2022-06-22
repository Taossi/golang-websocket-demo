package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

var addr = flag.String("addr", "localhost:8088", "http server")

func main() {
	flag.Parse()
	log.SetFlags(0)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
		return
	}
	defer conn.Close()

	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("read message error", err)
		return
	}
	log.Println(string(msg))
}
