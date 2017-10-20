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

func sendUpdate(fromClient *websocket.Conn){
	for i,client := range(clients){
		err := client.WriteMessage(1, []byte(d.GetDocumentValue()))

		if err != nil {
			clients = append(clients[:i], clients[i+1:]...)
			log.Println(err)
			continue;
		}
	}

}

func processMessages(c *websocket.Conn){
	defer closeClient(c);

	c.WriteMessage(1, []byte(d.GetDocumentValue()))

	for {
		t, mess, err := c.ReadMessage()
		if(t != -1){
			d.OverwriteText(string(mess));
			sendUpdate(c);
		}
		if err != nil {
			log.Println("read:", err)
			break
		}
	}
}

func closeClient(c *websocket.Conn){
	log.Println("close requested");
	c.Close()
}

func addClient(c *websocket.Conn){
	clients = append(clients,c);
	go processMessages(c);
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
