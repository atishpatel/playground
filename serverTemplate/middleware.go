package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"servertemplate/httperrors"
	"servertemplate/logging"
	"time"
)

func (s *server) JSONResponse(f apiHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v, err := f(w, r)
		if err != nil {
			// return error as json
			var httpErr *httperrors.HTTPError
			if errors.As(err, &httpErr) {
				logging.Errorf("httpErr: %v", httpErr)
			} else {
				// This shouldn't happen if errors are handled properly. Wrap error in httperror for now.
				logging.Errorf("unknown error context that needs to be wrapped: %v", err)
				httpErr = &httperrors.HTTPError{
					Status:    http.StatusInternalServerError,
					Err:       err,
					ErrString: err.Error(),
				}
			}
			w.WriteHeader(httpErr.Status)
			err = json.NewEncoder(w).Encode(httpErr)
			if err != nil {
				logging.Errorf("failed to encode json for hello: %v", err)
			}
			return
		}
		// return value as json
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		err = json.NewEncoder(w).Encode(v)
		if err != nil {
			logging.Errorf("failed to encode json for hello: %v", err)
		}
	})
}

func (s *server) BasicHandler(next http.Handler) http.Handler {
	return s.RequestLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handle options requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		// handle panic, recover
		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("PANICKING in handler: %+v\n%s", r, debug.Stack())
				logging.Errorf("%v", err)
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("{\"status\":500,\"error\":\"Woops! Something went wrong. Please try again later.\"}"))
				return
			}
		}()
		next.ServeHTTP(w, r)
	}))
}

// RequestLogger logs requests with status and latency.
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

type ResponseWriterWithStatus struct {
	http.ResponseWriter
	status int
}

func (w *ResponseWriterWithStatus) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
