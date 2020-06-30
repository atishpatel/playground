package main

import (
	"net/http"
	"servertemplate/logging"
)

func main() {
	// findError()
	// workers()
	logging.Infof("Starting Server")
	s := newServer()
	http.Handle("/", s.BasicHandler(s.JSONResponse((s.success))))
	http.Handle("/fail", s.BasicHandler(s.JSONResponse((s.fail))))
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		logging.Errorf("failed to ListenAndServe: %v", err)
	}
}

func newServer() *server {
	s := &server{}
	return s
}

type server struct {
	// db
	// logger
}
