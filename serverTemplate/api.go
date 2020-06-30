package main

import (
	"net/http"
	"servertemplate/httperrors"
)

type apiHandler func(http.ResponseWriter, *http.Request) (interface{}, error)

func (s *server) success(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	v := struct {
		Message string `json:"message"`
	}{
		Message: "hello world!",
	}
	return v, nil
}

func (s *server) fail(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	v := struct {
		Message string `json:"message"`
	}{
		Message: "hello world!",
	}
	return v, httperrors.Newf(http.StatusBadRequest, "failed with bad request")
	// return v, errors.New("bad error")
}
