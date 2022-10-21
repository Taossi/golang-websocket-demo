package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
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

	SendMessage(conn)
}

func SendMessage(conn *websocket.Conn) error {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		tp, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("write:", err)
			return err
		}
		log.Println("message type:" + fmt.Sprint(tp))
		log.Println("message:" + string(msg))
	}
	return nil
}
