package server_test

import (
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/skybox/skybox/db"
	. "github.com/skybox/skybox/server"
	"github.com/stretchr/testify/assert"
)

// Ensure that the handler returns a home page.
func TestHandlerIndex(t *testing.T) {
	withHandler(func(h *Handler) {
		var resp = httptest.NewRecorder()
		h.ServeHTTP(resp, MustParseRequest("GET", "http://localhost/", nil))
		assert.Equal(t, resp.Code, 200)
	})
}

// Ensure that the tracking pixel records an anonymous event.
func TestHandlerTrackAnonymous(t *testing.T) {
	withHandlerAndAccount(func(h *Handler, a *db.Account) {
		var resp = httptest.NewRecorder()
		h.ServeHTTP(resp, MustParseRequest("GET", "/track.png?device.id=device0&channel=web&resource=/users/:id/items&action=view&domain=foo.com&path=/users/123/items&apiKey="+a.APIKey, nil))
		assert.Equal(t, resp.Code, 200)

		// Check that Sky recorded it.
		events, err := a.SkyTable().Events("device0")
		assert.NoError(t, err)
		if assert.Equal(t, len(events), 1) {
			assert.True(t, time.Now().Sub(events[0].Timestamp) < time.Second)
			assert.Equal(t, events[0].Data["channel"], "web")
			assert.Equal(t, events[0].Data["resource"], "/users/:id/items")
			assert.Equal(t, events[0].Data["action"], "view")
			assert.Equal(t, events[0].Data["domain"], "foo.com")
			assert.Equal(t, events[0].Data["path"], "/users/123/items")
		}
	})
}

// Ensure that the tracking pixel records a known user event.
func TestHandlerTrackKnownUser(t *testing.T) {
	withHandlerAndAccount(func(h *Handler, a *db.Account) {
		var resp = httptest.NewRecorder()
		h.ServeHTTP(resp, MustParseRequest("GET", "/track.png?user.id=user0&device.id=device0&channel=web&resource=/users/:id/items&action=view&domain=foo.com&path=/users/123/items&apiKey="+a.APIKey, nil))
		assert.Equal(t, resp.Code, 200)

		// Check that Sky recorded it.
		events, err := a.SkyTable().Events("@user0")
		assert.NoError(t, err)
		if assert.Equal(t, len(events), 1) {
			assert.True(t, time.Now().Sub(events[0].Timestamp) < time.Second)
			assert.Equal(t, events[0].Data["resource"], "/users/:id/items")
		}
	})
}

// Ensure that the tracking pixel returns a "bad request" if the API key is invalid.
func TestHandlerTrackWithInvalidAPIKey(t *testing.T) {
	withHandler(func(h *Handler) {
		var resp = httptest.NewRecorder()
		h.ServeHTTP(resp, MustParseRequest("GET", "/track.png?apiKey=no_such_key", nil))
		assert.Equal(t, resp.Code, 400)
	})
}

// withHandler executes a function with an open handler.
func withHandler(fn func(*Handler)) {
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

	// Initialize handler.
	h, err := NewHandler(&db)
	if err != nil {
		panic("handler error: " + err.Error())
	}
	fn(h)
}

// withHandlerAndAccount executes a function with an open handler and created account.
func withHandlerAndAccount(fn func(*Handler, *db.Account)) {
	withHandler(func(h *Handler) {
		a := &db.Account{}
		err := h.DB().Do(func(tx *db.Tx) error {
			if err := tx.CreateAccount(a); err != nil {
				panic("create account error: " + err.Error())
			}
			if err := a.Reset(); err != nil {
				panic("reset account error: " + err.Error())
			}
			return nil
		})
		if err != nil {
			panic("init error: " + err.Error())
		}
		fn(h, a)
	})
}
