package server_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Creates a new HTTP client with KeepAlive disabled.
func NewHTTPClient() *http.Client {
	return &http.Client{Transport: &http.Transport{DisableKeepAlives: true}}
}

func ReadBody(resp *http.Response) []byte {
	if resp == nil {
		return []byte{}
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body
}

// Reads the body from the response and parses it as JSON.
func ReadBodyJSON(resp *http.Response) map[string]interface{} {
	m := make(map[string]interface{})
	b := ReadBody(resp)
	if err := json.Unmarshal(b, &m); err != nil {
		panic(fmt.Sprintf("HTTP body JSON parse error: %v", err))
	}
	return m
}

func getHTML(path string) (status int, body string) {
	c := NewHTTPClient()
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost%s%s", 7000, path), nil)
	resp, err := c.Do(req)
	if err != nil {
		panic("get http error: " + err.Error())
	}
	status = resp.StatusCode
	b, _ := ioutil.ReadAll(resp.Body)
	body = string(b)
	return
}
