package main

import (
	"log"
	"net/http"
)

var templates *templates.Templates

func main() {
	var err error
	// templates, err = got.New("templates", ".html")
	if err != nil {
		log.Fatalf("failed to load templates, %+v", err)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/server1", index)
	http.HandleFunc("/checkout/step1", checkout)
	log.Println("starting server1 on port 9001")
	http.ListenAndServe(":9001", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("server1 hit")
	data := Page{}
	err := templates.Render(w, "home", data, http.StatusOK)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func checkout(w http.ResponseWriter, r *http.Request) {
	log.Println("server1 hit")
	data := Page{}
	err := templates.Render(w, "checkout", data, http.StatusOK)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
