package server_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/benbjohnson/skybox/db"
	. "github.com/benbjohnson/skybox/server"
	"github.com/stretchr/testify/assert"
)

const testAddr = ":7000"

// Ensure that the server returns a home page.
func TestServerIndex(t *testing.T) {
	withServer(func(s *Server) {
		status, body := getHTML("/")
		assert.Equal(t, status, 200)
		assert.Equal(t, body, "-")
	})
}

// withServer executes a function with an open server.
func withServer(fn func(*Server)) {
	f, _ := ioutil.TempFile("", "skybox-")
	path := f.Name()
	f.Close()
	os.Remove(path)
	defer os.RemoveAll(path)

	// Create database.
	var db db.DB
	if err := db.Open(path, 0666); err != nil {
		panic("db open error: " + err.Error())
	}
	defer db.Close()

	// Start server.
	var s Server
	s.Addr = testAddr
	s.DB = &db

	go s.ListenAndServe()
	defer s.Close()

	// Execute function.
	fn(&s)
}
