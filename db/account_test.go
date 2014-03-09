package db_test

import (
	"testing"
	"time"

	. "github.com/skybox/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that an account can be deleted.
func TestAccountDelete(t *testing.T) {
	withDB(func(db *DB) {
		// Create account.
		err := db.Do(func(tx *Tx) error {
			return tx.CreateAccount(&Account{})
		})
		assert.NoError(t, err)

		// Retrieve and delete account.
		err = db.Do(func(tx *Tx) error {
			a, _ := tx.Account(1)
			return a.Delete()
		})
		assert.NoError(t, err)

		// Retrieve the account again.
		err = db.With(func(tx *Tx) error {
			_, err := tx.Account(1)
			return err
		})
		assert.Equal(t, err, ErrAccountNotFound)
	})
}

// Ensure that an account can retrieve all associated users.
func TestAccountUsers(t *testing.T) {
	withDB(func(db *DB) {
		// Create two accounts.
		db.Do(func(tx *Tx) error {
			a1 := &Account{}
			assert.NoError(t, tx.CreateAccount(a1))
			a2 := &Account{}
			assert.NoError(t, tx.CreateAccount(a2))
			return nil
		})

		// Add users to first account.
		db.Do(func(tx *Tx) error {
			a1, _ := tx.Account(1)
			assert.NoError(t, a1.CreateUser(&User{Email: "susyque@gmail.com", Password: "password"}))
			assert.NoError(t, a1.CreateUser(&User{Email: "johndoe@gmail.com", Password: "password"}))
			return nil
		})

		// Add users to second account.
		db.Do(func(tx *Tx) error {
			a2, _ := tx.Account(2)
			assert.NoError(t, a2.CreateUser(&User{Email: "billybob@gmail.com", Password: "password"}))
			return nil
		})

		// Check first account users.
		db.With(func(tx *Tx) error {
			a1, _ := tx.Account(1)
			users, err := a1.Users()
			if assert.NoError(t, err) && assert.Equal(t, len(users), 2) {
				assert.Equal(t, users[0].Tx, tx)
				assert.Equal(t, users[0].ID(), 2)
				assert.Equal(t, users[0].AccountID, 1)
				assert.Equal(t, users[0].Email, "johndoe@gmail.com")

				assert.Equal(t, users[1].Tx, tx)
				assert.Equal(t, users[1].ID(), 1)
				assert.Equal(t, users[1].AccountID, 1)
				assert.Equal(t, users[1].Email, "susyque@gmail.com")
			}
			return nil
		})

		// Check second account users.
		db.With(func(tx *Tx) error {
			a2, _ := tx.Account(2)
			users, err := a2.Users()
			if assert.NoError(t, err) && assert.Equal(t, len(users), 1) {
				assert.Equal(t, users[0].Tx, tx)
				assert.Equal(t, users[0].ID(), 3)
				assert.Equal(t, users[0].AccountID, 2)
				assert.Equal(t, users[0].Email, "billybob@gmail.com")
			}
			return nil
		})
	})
}

// Ensure that an account can generate a random API key.
func TestAccountGenerateAPIKey(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			a := &Account{}
			assert.NoError(t, tx.CreateAccount(a))

			// Check for an API key.
			assert.Equal(t, len(a.APIKey), 36)

			// Lookup account by API key.
			a2, err := tx.AccountByAPIKey(a.APIKey)
			assert.NoError(t, err)
			assert.Equal(t, a2.ID(), 1)

			// Regenerate key.
			apiKey := a.APIKey
			assert.NoError(t, a.GenerateAPIKey())

			// Make sure it's not the same as before.
			assert.NotEqual(t, a.APIKey, apiKey)

			// Make sure we can lookup by the new key and not the old.
			a3, err := tx.AccountByAPIKey(a.APIKey)
			assert.NoError(t, err)
			assert.Equal(t, a3.ID(), 1)

			a4, err := tx.AccountByAPIKey(apiKey)
			assert.Equal(t, err, ErrAccountNotFound)
			assert.Nil(t, a4)

			return nil
		})
	})
}

// Ensure that an account can retrieve all associated funnels.
func TestAccountFunnels(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			a1 := &Account{}
			a2 := &Account{}
			assert.NoError(t, tx.CreateAccount(a1))
			assert.NoError(t, tx.CreateAccount(a2))

			// Add funnels to first account.
			assert.NoError(t, a1.CreateFunnel(&Funnel{Name: "Funnel B", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}))
			assert.NoError(t, a1.CreateFunnel(&Funnel{Name: "Funnel A", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}))

			// Add funnels to second account.
			assert.NoError(t, a2.CreateFunnel(&Funnel{Name: "Funnel C", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}))
			return nil
		})

		// Check first account.
		db.With(func(tx *Tx) error {
			a, _ := tx.Account(1)
			funnels, err := a.Funnels()
			if assert.NoError(t, err) && assert.Equal(t, len(funnels), 2) {
				assert.Equal(t, funnels[0].Tx, tx)
				assert.Equal(t, funnels[0].ID(), 2)
				assert.Equal(t, funnels[0].AccountID, 1)
				assert.Equal(t, funnels[0].Name, "Funnel A")

				assert.Equal(t, funnels[1].Tx, tx)
				assert.Equal(t, funnels[1].ID(), 1)
				assert.Equal(t, funnels[1].AccountID, 1)
				assert.Equal(t, funnels[1].Name, "Funnel B")
			}

			// Make sure we can only get a1 funnels.
			f, err := a.Funnel(1)
			assert.NoError(t, err)
			assert.NotNil(t, f)
			f, err = a.Funnel(3)
			assert.Equal(t, err, ErrFunnelNotFound)
			assert.Nil(t, f)

			return nil
		})

		// Check second account's funnels.
		db.With(func(tx *Tx) error {
			a, _ := tx.Account(2)
			funnels, err := a.Funnels()
			if assert.NoError(t, err) && assert.Equal(t, len(funnels), 1) {
				assert.Equal(t, funnels[0].Tx, tx)
				assert.Equal(t, funnels[0].ID(), 3)
				assert.Equal(t, funnels[0].AccountID, 2)
				assert.Equal(t, funnels[0].Name, "Funnel C")
			}
			return nil
		})
	})
}

// Ensure that an account can track events and insert them into Sky.
func TestAccountTrack(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			a := &Account{}
			tx.CreateAccount(a)
			a.Reset()

			// Add some events.
			assert.NoError(t, a.Track(newTestEvent("2000-01-01T00:00:00Z", "john", "DEV0", "web", "/", "view", nil)))
			assert.NoError(t, a.Track(newTestEvent("2000-01-01T00:00:05Z", "john", "DEV1", "web", "/signup", "view", nil)))
			assert.NoError(t, a.Track(newTestEvent("2000-01-01T00:00:00Z", "susy", "DEV2", "web", "/cancel", "click", nil)))

			// Verify "john" events.
			events, err := a.Events("@john")
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
			events, err = a.Events("@susy")
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

// Ensure that an account can retrieve a list of unique resources.
func TestAccountResources(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			a1, a2, a3 := &Account{}, &Account{}, &Account{}
			tx.CreateAccount(a1)
			tx.CreateAccount(a2)
			tx.CreateAccount(a3)
			a1.Reset()
			a2.Reset()
			a3.Reset()

			// Add some events.
			assert.NoError(t, a1.Track(newTestEvent("2000-01-01T00:00:00Z", "john", "DEV0", "web", "/", "view", nil)))
			assert.NoError(t, a1.Track(newTestEvent("2000-01-01T00:00:01Z", "john", "DEV1", "web", "/signup", "view", nil)))
			assert.NoError(t, a1.Track(newTestEvent("2000-01-01T00:00:02Z", "susy", "DEV2", "web", "/cancel", "view", nil)))
			assert.NoError(t, a2.Track(newTestEvent("2000-01-01T00:00:02Z", "john", "DEV2", "web", "/blah", "view", nil)))

			// Retrieve resources.
			resources, err := a1.Resources()
			assert.NoError(t, err)
			assert.Equal(t, resources, []string{"/", "/cancel", "/signup"})

			resources, err = a2.Resources()
			assert.NoError(t, err)
			assert.Equal(t, resources, []string{"/blah"})

			resources, err = a3.Resources()
			assert.NoError(t, err)
			assert.Equal(t, resources, []string{})
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
