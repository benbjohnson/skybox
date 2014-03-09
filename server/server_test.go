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
		assert.NotNil(t, body)
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

	c := make(chan bool)
	go func() {
		c <- true
		s.ListenAndServe()
	}()
	<-c

	// Execute function.
	fn(&s)

	s.Close()
}

// withServerAndProject executes a function with an open server and created project.
func withServerAndProject(fn func(*Server, *db.Project)) {
	withServer(func(s *Server) {
		p := &db.Project{Name: "My Project"}
		err := s.DB.Do(func(tx *db.Tx) error {
			a := &db.Account{}
			if err := tx.CreateAccount(a); err != nil {
				panic("create account error: " + err.Error())
			}
			if err := a.CreateProject(p); err != nil {
				panic("create project error: " + err.Error())
			}
			if err := p.Reset(); err != nil {
				panic("reset project error: " + err.Error())
			}
			return nil
		})
		if err != nil {
			panic("init error: " + err.Error())
		}
		fn(s, p)
	})
}
