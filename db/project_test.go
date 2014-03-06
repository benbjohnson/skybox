package db_test

import (
	"testing"
	"time"

	. "github.com/benbjohnson/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that an account can create a project.
func TestProjectCreate(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Create an account and project.
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			p := &Project{Name: "Project X"}
			assert.NoError(t, a.CreateProject(p))
			assert.Equal(t, p.ID(), 1)

			// Retrieve the project.
			p2, err := txn.Project(1)
			if assert.NoError(t, err) && assert.NotNil(t, p2) {
				assert.Equal(t, p2.Transaction, txn)
				assert.Equal(t, p2.ID(), 1)
				assert.Equal(t, p2.AccountID, 1)
				assert.Equal(t, p2.Name, "Project X")
			}
			return nil
		})
	})
}

// Ensure that an account cannot create a project after it's deleted.
func TestProjectCreateAfterDeletion(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Create an account and delete it.
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			assert.NoError(t, a.Delete())

			// Attempt to create a project.
			err := a.CreateProject(&Project{Name: "Project X"})
			assert.Equal(t, err, ErrAccountNotFound)
			return nil
		})
	})
}

// Ensure that creating an invalid project returns an error.
func TestProjectCreateMissingName(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			assert.Equal(t, a.CreateProject(&Project{}), ErrProjectNameRequired)
			return nil
		})
	})
}

// Ensure that a project can update itself.
func TestProjectUpdate(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Create account and project.
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			p := &Project{Name: "Project X"}
			assert.NoError(t, a.CreateProject(p))

			// Update the project.
			p.Name = "Project Y"
			p.Save()

			// Retrieve the project.
			p2, err := txn.Project(1)
			if assert.NoError(t, err) && assert.NotNil(t, p2) {
				assert.Equal(t, p2.Name, "Project Y")
			}
			return nil
		})
	})
}

// Ensure that a user can be deleted.
func TestProjectDelete(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Create account and project.
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			p := &Project{Name: "Project X"}
			assert.NoError(t, a.CreateProject(p))

			// Delete the project.
			assert.NoError(t, p.Delete())

			// Retrieve the project again.
			_, err := txn.Project(1)
			assert.Equal(t, err, ErrProjectNotFound)
			return nil
		})
	})
}

// Ensure that a project can generate a random API key.
func TestProjectGenerateAPIKey(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Create an account and project.
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			p := &Project{Name: "Project X"}
			assert.NoError(t, a.CreateProject(p))

			// Check for an API key.
			assert.Equal(t, len(p.APIKey), 36)

			// Lookup project by API key.
			p2, err := txn.ProjectByAPIKey(p.APIKey)
			assert.NoError(t, err)
			assert.Equal(t, p2.ID(), 1)

			// Regenerate key.
			apiKey := p.APIKey
			assert.NoError(t, p.GenerateAPIKey())

			// Make sure it's not the same as before.
			assert.NotEqual(t, p.APIKey, apiKey)

			// Make sure we can lookup by the new key and not the old.
			p3, err := txn.ProjectByAPIKey(p.APIKey)
			assert.NoError(t, err)
			assert.Equal(t, p3.ID(), 1)

			p4, err := txn.ProjectByAPIKey(apiKey)
			assert.Equal(t, err, ErrProjectNotFound)
			assert.Nil(t, p4)

			return nil
		})
	})
}

// Ensure that a project can retrieve all associated funnels.
func TestProjectFunnels(t *testing.T) {
	withDB(func(db *DB) {
		// Create two projects.
		db.Do(func(txn *Transaction) error {
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			p1 := &Project{Name: "Project X"}
			assert.NoError(t, a.CreateProject(p1))
			p2 := &Project{Name: "Project Y"}
			assert.NoError(t, a.CreateProject(p2))

			// Add funnels to first project.
			assert.NoError(t, p1.CreateFunnel(&Funnel{Name: "Funnel B", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}))
			assert.NoError(t, p1.CreateFunnel(&Funnel{Name: "Funnel A", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}))

			// Add projects to second account.
			assert.NoError(t, p2.CreateFunnel(&Funnel{Name: "Funnel C", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}))
			return nil
		})

		// Check first project.
		db.With(func(txn *Transaction) error {
			p, _ := txn.Project(1)
			funnels, err := p.Funnels()
			if assert.NoError(t, err) && assert.Equal(t, len(funnels), 2) {
				assert.Equal(t, funnels[0].Transaction, txn)
				assert.Equal(t, funnels[0].ID(), 2)
				assert.Equal(t, funnels[0].ProjectID, 1)
				assert.Equal(t, funnels[0].Name, "Funnel A")

				assert.Equal(t, funnels[1].Transaction, txn)
				assert.Equal(t, funnels[1].ID(), 1)
				assert.Equal(t, funnels[1].ProjectID, 1)
				assert.Equal(t, funnels[1].Name, "Funnel B")
			}

			// Make sure we can only get p1 funnels.
			f, err := p.Funnel(1)
			assert.NoError(t, err)
			assert.NotNil(t, f)
			f, err = p.Funnel(3)
			assert.Equal(t, err, ErrFunnelNotFound)
			assert.Nil(t, f)

			return nil
		})

		// Check second project's funnels.
		db.With(func(txn *Transaction) error {
			p, _ := txn.Project(2)
			funnels, err := p.Funnels()
			if assert.NoError(t, err) && assert.Equal(t, len(funnels), 1) {
				assert.Equal(t, funnels[0].Transaction, txn)
				assert.Equal(t, funnels[0].ID(), 3)
				assert.Equal(t, funnels[0].ProjectID, 2)
				assert.Equal(t, funnels[0].Name, "Funnel C")
			}
			return nil
		})
	})
}

// Ensure that a project can track events and insert them into Sky.
func TestProjectTrack(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			a, p := &Account{}, &Project{Name: "prj"}
			txn.CreateAccount(a)
			a.CreateProject(p)
			p.Reset()

			// Add some events.
			err := p.Track(newTestEvent("2000-01-01T00:00:00Z", "john", "DEV0", "web", "/", "view", nil))
			assert.NoError(t, err)
			err = p.Track(newTestEvent("2000-01-01T00:00:05Z", "john", "DEV1", "web", "/signup", "view", nil))
			assert.NoError(t, err)
			err = p.Track(newTestEvent("2000-01-01T00:00:00Z", "susy", "DEV2", "web", "/cancel", "click", nil))
			assert.NoError(t, err)

			// Verify "john" events.
			events, err := p.Events("john")
			assert.NoError(t, err)
			if assert.Equal(t, len(events), 2) {
				assert.Equal(t, events[0].Timestamp, mustParseTime("2000-01-01T00:00:00Z"))
				assert.Equal(t, events[0].Channel, "web")
				assert.Equal(t, events[0].Resource, "/")
				assert.Equal(t, events[0].Action, "view")

				assert.Equal(t, events[1].Timestamp, mustParseTime("2000-01-01T00:00:05Z"))
				assert.Equal(t, events[1].Channel, "web")
				assert.Equal(t, events[1].Resource, "/signup")
				assert.Equal(t, events[1].Action, "view")
			}

			// Verify "susy" events.
			events, err = p.Events("susy")
			assert.NoError(t, err)
			if assert.Equal(t, len(events), 1) {
				assert.Equal(t, events[0].Timestamp, mustParseTime("2000-01-01T00:00:00Z"))
				assert.Equal(t, events[0].Channel, "web")
				assert.Equal(t, events[0].Resource, "/cancel")
				assert.Equal(t, events[0].Action, "click")
			}
			return nil
		})
	})
}

func newTestEvent(timestamp, userID, deviceID, channel, resource, action string, data map[string]interface{}) *Event {
	return &Event{
		Timestamp: mustParseTime(timestamp),
		UserID:    userID,
		DeviceID:  deviceID,
		Channel:   channel,
		Resource:  resource,
		Action:    action,
		Data:      data,
	}
}

func mustParseTime(timestamp string) time.Time {
	ts, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		panic("invalid timestamp: " + err.Error())
	}
	return ts.UTC()
}
