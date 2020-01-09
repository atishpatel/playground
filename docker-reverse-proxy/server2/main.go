package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/server2", hello)
	log.Println("starting server2 on port 9002")
	http.ListenAndServe(":9002", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("server2 hit")
	w.Write([]byte("hello from server2"))
}
