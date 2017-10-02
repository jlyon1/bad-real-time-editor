package main

import (
	"flag"
	"log"
	"net/http"
	"main/mysocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
func test(w http.ResponseWriter, r *http.Request){
	log.Printf("asdf");
}

func main() {

	mysocket.Init();
  fs := http.FileServer(http.Dir("static"))

  http.Handle("/", fs)
	http.HandleFunc("/test", test)

	http.HandleFunc("/ws", mysocket.ReceiveClient)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
