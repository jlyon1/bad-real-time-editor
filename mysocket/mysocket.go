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
	for _,client := range(clients){
		err := client.WriteMessage(1, []byte(d.GetDocumentValue()))
		if err != nil {
			log.Println(err)
			continue;
		}
	}

}

func closeClient(c *websocket.Conn, pos int){
	log.Println("close requested");
	c.Close()
}

func addClient(c *websocket.Conn){
	clients = append(clients,c);
	pos := len(clients)
	defer closeClient(c,pos);

	c.WriteMessage(1, []byte(d.GetDocumentValue()))

	for {
		_, message, err := c.ReadMessage()
		d.OverwriteText(string(message));
		if(string(message) != ""){
			sendUpdate(c);
		}
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
