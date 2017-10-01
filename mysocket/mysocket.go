package mysocket

import (
	"log"
	"net/http"
	"main/document"
	"github.com/gorilla/websocket"
)

var clients []*websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func addClient(c *websocket.Conn){
	defer c.Close()

	clients = append(clients,c);
	c.WriteMessage(1, []byte("Connected"))
	var d = document.New("asdf");
	log.Print(d.GetDocumentValue());
	for {
		_, message, err := c.ReadMessage()
		log.Println(string(message));
		if err != nil {
			log.Println("read:", err)
			break
		}
	}
}

func ReceiveClient(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
		} else {
			addClient(c)
		}
	}
