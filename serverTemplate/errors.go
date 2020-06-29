package main

import (
	"errors"
	"fmt"
	"net/http"
	"servertemplate/httperrors"
	"servertemplate/logging"
)

var (
	ErrNotFound = errors.New("not found")
)

func wrapError() error {
	err := httperrors.New(http.StatusNotFound, ErrNotFound)
	return fmt.Errorf("wrapping: %w", err)
}

func findError() {
	err := wrapError()
	logging.Infof("err: %v", err)
	var httpErr *httperrors.HTTPError
	if errors.As(err, &httpErr) {
		logging.Infof("found httpErr: %v", httpErr)
		if httpErr.Err == ErrNotFound {
			logging.Infof("httpError.Err matches ErrNotFound")
		}
	}
}
