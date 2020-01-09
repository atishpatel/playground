package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/server1", hello)
	log.Println("starting server1 on port 9001")
	http.ListenAndServe(":9001", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println("server1 hit")
	w.Write([]byte("hello from server1"))
}
