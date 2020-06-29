package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"servertemplate/logging"
	"time"
)

func main() {
	// findError()
	// workers()
	logging.Infof("Starting Server")
	s := newServer()
	http.Handle("/", s.BasicHandler(s.RequestLogger(http.HandlerFunc(s.hello))))
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

}

func (s *server) BasicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handle options requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		// handle panic, recover
		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("PANICKING in handler: %+v\n%s", r, debug.Stack())
				logging.Errorf("%w", err)
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(fmt.Sprintf("{\"code\":500,\"error\":\"Woops! Something went wrong. Please try again later.\"}")))
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// RequestLogger logs requests.
// Modeled after GIN's request logger but missing the cool colors for now.
func (s *server) RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start timer
		start := time.Now()
		path := r.URL.Path
		raw := r.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}
		writer := ResponseWriterWithStatus{ResponseWriter: w}
		// Process request
		next.ServeHTTP(&writer, r)
		// Stop timer
		timeStamp := time.Now()
		latency := timeStamp.Sub(start)
		clientAddr := r.RemoteAddr
		method := r.Method
		statusCode := writer.status
		// log request
		logging.Requestf("REQUEST: %s | %d | %13v | %s | %s | %s",
			timeStamp.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientAddr,
			method,
			path,
		)
	})
}

func (s *server) hello(w http.ResponseWriter, r *http.Request) {
	v := struct {
		Message string `json:"message"`
	}{
		Message: "hello world!",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		logging.Errorf("failed to encode json for hello: %v", err)
	}
}

type ResponseWriterWithStatus struct {
	http.ResponseWriter
	status int
}

func (w *ResponseWriterWithStatus) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
