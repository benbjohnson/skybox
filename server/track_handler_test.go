package server_test

import (
	"testing"
	"time"

	"github.com/benbjohnson/skybox/db"
	. "github.com/benbjohnson/skybox/server"
	"github.com/stretchr/testify/assert"
)

// Ensure that the tracking pixel records an anonymous event.
func TestServerTrackAnonymous(t *testing.T) {
	withServerAndProject(func(s *Server, p *db.Project) {
		// Track event.
		status, body := getHTML("/track.png?device.id=device0&channel=web&resource=/users/:id/projects&action=view&domain=foo.com&path=/users/123/projects&apiKey=" + p.APIKey)
		assert.Equal(t, status, 200)
		assert.NotNil(t, body)

		// Check that Sky recorded it.
		events, err := p.SkyTable().Events("@device0")
		assert.NoError(t, err)
		if assert.Equal(t, len(events), 1) {
			assert.True(t, time.Now().Sub(events[0].Timestamp) < time.Second)
			assert.Equal(t, events[0].Data["channel"], "web")
			assert.Equal(t, events[0].Data["resource"], "/users/:id/projects")
			assert.Equal(t, events[0].Data["action"], "view")
			assert.Equal(t, events[0].Data["domain"], "foo.com")
			assert.Equal(t, events[0].Data["path"], "/users/123/projects")
		}
	})
}

// Ensure that the tracking pixel records a known user event.
func TestServerTrackKnownUser(t *testing.T) {
	withServerAndProject(func(s *Server, p *db.Project) {
		// Track event.
		status, body := getHTML("/track.png?user.id=user0&device.id=device0&channel=web&resource=/users/:id/projects&action=view&domain=foo.com&path=/users/123/projects&apiKey=" + p.APIKey)
		assert.Equal(t, status, 200)
		assert.NotNil(t, body)

		// Check that Sky recorded it.
		events, err := p.SkyTable().Events("user0")
		assert.NoError(t, err)
		if assert.Equal(t, len(events), 1) {
			assert.True(t, time.Now().Sub(events[0].Timestamp) < time.Second)
			assert.Equal(t, events[0].Data["resource"], "/users/:id/projects")
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
