package httperrors

import "fmt"

// HTTPError is an error type that contains the status code to return with the request.
type HTTPError struct {
	Status int
	Err    error
}

// Error is the error associated with the HTTPError.
func (e *HTTPError) Error() string {
	return e.Err.Error()
}

// Unwrap returns the orginal internal error.
func (e *HTTPError) Unwrap() error {
	return e.Err
}

// New returns a new HTTPError.
func New(status int, err error) error {
	return &HTTPError{
		Status: status,
		Err:    err,
	}
}

// Newf formats and wraps an error and returns a new HTTPError.
// Please remember to us %w if there is an error in the argument.
func Newf(status int, format string, args ...interface{}) error {
	return &HTTPError{
		Status: status,
		Err:    fmt.Errorf(format, args...),
	}
}
