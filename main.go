package main

import (
	"flag"
	"log"
	"net/http"
	"main/mysocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")


func main() {

	mysocket.Init();
  fs := http.FileServer(http.Dir("static"))

  http.Handle("/", fs)
	http.HandleFunc("/ws", mysocket.ReceiveClient)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
