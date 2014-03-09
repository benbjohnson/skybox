package server_test

import (
	"testing"
	"time"

	"github.com/skybox/skybox/db"
	. "github.com/skybox/skybox/server"
	"github.com/stretchr/testify/assert"
)

// Ensure that the tracking pixel records an anonymous event.
func TestServerTrackAnonymous(t *testing.T) {
	withServerAndAccount(func(s *Server, a *db.Account) {
		// Track event.
		status, body := getHTML("/track.png?device.id=device0&channel=web&resource=/users/:id/items&action=view&domain=foo.com&path=/users/123/items&apiKey=" + a.APIKey)
		assert.Equal(t, status, 200)
		assert.NotNil(t, body)

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
func TestServerTrackKnownUser(t *testing.T) {
	withServerAndAccount(func(s *Server, a *db.Account) {
		// Track event.
		status, body := getHTML("/track.png?user.id=user0&device.id=device0&channel=web&resource=/users/:id/items&action=view&domain=foo.com&path=/users/123/items&apiKey=" + a.APIKey)
		assert.Equal(t, status, 200)
		assert.NotNil(t, body)

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
func TestServerTrackWithInvalidAPIKey(t *testing.T) {
	withServer(func(s *Server) {
		status, body := getHTML("/track.png?apiKey=no_such_key")
		assert.Equal(t, status, 400)
		assert.NotNil(t, body)
	})
}
