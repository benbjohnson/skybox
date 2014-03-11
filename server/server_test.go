package server_test

import (
	"io"
	"net/http"
)

// MustParseRequest creates a new HTTP request and panics if there is an error.
func MustParseRequest(method, urlStr string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		panic("request error: " + err.Error())
	}
	return req
}
