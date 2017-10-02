package mysocket

import (
	"log"
	"net/http"
	"main/document"
	"github.com/gorilla/websocket"
)

var clients []*websocket.Conn

var d document.Document

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func sendUpdate(){
	for _,client := range(clients){
		err := client.WriteMessage(1, []byte(d.GetDocumentValue()))
		if err != nil {
			log.Println(err)
			break;
		}
	}
}

func addClient(c *websocket.Conn){
	defer c.Close()

	clients = append(clients,c);
	c.WriteMessage(1, []byte("Connected"))

	for {
		_, message, err := c.ReadMessage()
		d.OverwriteText(string(message));
		sendUpdate();
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
			go addClient(c)
		}
	}

func Init(){
	d = document.New("asdf");
}
