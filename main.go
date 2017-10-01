package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

// var wsUpgrader = websocket.Upgrader{

// }

var upgrader = websocket.Upgrader{
      ReadBufferSize:  1024,
      WriteBufferSize: 1024,
      CheckOrigin: func(r *http.Request) bool {
          return true
      },
  }

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

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	} else {
		addClient(c)
	}

}

func main() {
	messages = append(messages, "Welcome to the server")
	messages = append(messages, "Go away")
	go handleClientMessages()
	flag.Parse()
	log.SetFlags(0)
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)
	http.HandleFunc("/ws", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
