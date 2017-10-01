package main

import (
	"flag"
	"log"
	"net/http"
	"main/mysocket"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var clients []*websocket.Conn
var messages []string

func handleClientMessages() {
	oldLen := 0
	length := len(messages)
	for true {
		length = len(messages);
		if length != oldLen{
			log.Println(len(clients))
			for _, e := range clients {
				log.Println("ur here");
				for _, message := range messages[oldLen:]{
					err := e.WriteMessage(1, []byte(message))

					if err != nil {
						log.Println("read:", err)
						break
					}
				}
			}
			oldLen = len(messages);
		}
	}

}

func addClient(c *websocket.Conn) {
	defer c.Close()

	clients = append(clients,c);
	for _, message := range messages{
		err := c.WriteMessage(1, []byte(message))

		if err != nil {
			log.Println("read:", err)
			break
		}
	}
	for {

		_, message, err := c.ReadMessage()
		messages = append(messages,string(message));
		if err != nil {
			log.Println("read:", err)
			break
		}
	}
}


func main() {
	messages = append(messages, "Welcome to the server")
	messages = append(messages, "Go away")

	go handleClientMessages()

  fs := http.FileServer(http.Dir("static"))

  http.Handle("/", fs)
	http.HandleFunc("/ws", mysocket.ReceiveClient)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
