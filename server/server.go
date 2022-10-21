package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8088", "http server")

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // upgrade to websocket
	if err != nil {
		log.Println("upgrade error", err)
		return
	}
	defer conn.Close()

	message := "websocket success"
	SendMessage(message, conn)
	return
}

func SendMessage(message string, conn *websocket.Conn) error {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("write:", err)
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("http success")
}
