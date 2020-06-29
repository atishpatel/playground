package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"servertemplate/httperrors"
	"strings"
)

func buildJSONReq(method, url string, body interface{}) (*http.Request, error) {
	// encode body
	b := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(b).Encode(body)
		if err != nil {
			return nil, httperrors.Newf(http.StatusBadRequest, "failed to encode json: %w", err)
		}
	}
	// build request
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, httperrors.Newf(http.StatusBadRequest, "failed to create new request: %w", err)
	}
	// req.SetBasicAuth(c.apiKey, "")
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func decodeJSONResp(resp *http.Response, response interface{}) error {
	var err error
	contentType := resp.Header.Get("Content-Type")
	if resp.StatusCode == 204 || strings.Contains(contentType, "No Content") {
		return nil
	}
	if !strings.Contains(contentType, "json") {
		var b []byte
		var body string
		b, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			body = string(b)
		}
		return httperrors.Newf(http.StatusInternalServerError, "invalid Content-Type(%s) - StatusCode(%d) - %s", contentType, resp.StatusCode, body)
	}
	// has Content-Type json
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&response)
	if err != nil {
		return httperrors.Newf(http.StatusInternalServerError, "failed to decode json: %w", err)
	}
	return nil
}

func call3rdParty() {
	req, err := buildJSONReq(http.MethodGet, "", nil)
	if err != nil {

	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {

	}
	a := new(Answer)
	err = decodeJSONResp(resp, a)
	if err != nil {

	}
}

type Answer struct {
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Answer
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := todo // RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
