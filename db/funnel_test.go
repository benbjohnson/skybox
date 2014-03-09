package db_test

import (
	"strings"
	"testing"

	. "github.com/skybox/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that an account can create a funnel.
func TestFunnelCreate(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			// Create an account and funnel.
			a := &Account{}
			assert.NoError(t, tx.CreateAccount(a))
			f := &Funnel{Name: "Funnel Y", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}
			assert.NoError(t, a.CreateFunnel(f))
			assert.Equal(t, f.ID(), 1)

			// Retrieve the funnel.
			f2, err := tx.Funnel(1)
			if assert.NoError(t, err) && assert.NotNil(t, f2) {
				assert.Equal(t, f2.Tx, tx)
				assert.Equal(t, f2.ID(), 1)
				assert.Equal(t, f2.AccountID, 1)
				assert.Equal(t, f2.Name, "Funnel Y")
			}
			return nil
		})
	})
}

// Ensure that creating a funnel without a name returns an error.
func TestFunnelCreateMissingName(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			a := &Account{}
			assert.NoError(t, tx.CreateAccount(a))
			assert.Equal(t, a.CreateFunnel(&Funnel{Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}), ErrFunnelNameRequired)
			return nil
		})
	})
}

// Ensure that creating a funnel without steps returns an error.
func TestFunnelCreateMissingSteps(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			a := &Account{}
			assert.NoError(t, tx.CreateAccount(a))
			assert.Equal(t, a.CreateFunnel(&Funnel{Name: "Funnel Y"}), ErrFunnelStepsRequired)
			return nil
		})
	})
}

// Ensure that a funnel can update itself.
func TestFunnelUpdate(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			a := &Account{}
			assert.NoError(t, tx.CreateAccount(a))
			f := &Funnel{Name: "Funnel Y", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}
			assert.NoError(t, a.CreateFunnel(f))

			// Update the funnel.
			f.Name = "Funnel Z"
			f.Save()

			// Retrieve the funnel.
			f2, err := tx.Funnel(1)
			if assert.NoError(t, err) && assert.NotNil(t, f2) {
				assert.Equal(t, f2.Name, "Funnel Z")
			}
			return nil
		})
	})
}

// Ensure that a funnel can be deleted.
func TestFunnelDelete(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			// Create account and funnel.
			a := &Account{}
			assert.NoError(t, tx.CreateAccount(a))
			f := &Funnel{Name: "Funnel Y", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}
			assert.NoError(t, a.CreateFunnel(f))

			// Delete the funnel.
			assert.NoError(t, f.Delete())

			// Retrieve the funnel again.
			_, err := tx.Funnel(1)
			assert.Equal(t, err, ErrFunnelNotFound)
			return nil
		})
	})
}

// Ensure that a funnel can generate a correct query string.
func TestFunnelQueryString(t *testing.T) {
	f := &Funnel{
		Steps: []*FunnelStep{
			{Condition: "action == '/index.html'"},
			{Condition: "action == '/signup.html'"},
			{Condition: "action == '/checkout.html'"},
		},
	}

	exp := `
FOR EACH SESSION DELIMITED BY 2 HOURS
  FOR EACH EVENT
    WHEN action == '/index.html' THEN
      SELECT count() INTO "step0"
      WHEN action == '/signup.html' WITHIN 1..100000 STEPS THEN
        SELECT count() INTO "step1"
        WHEN action == '/checkout.html' WITHIN 1..100000 STEPS THEN
          SELECT count() INTO "step2"
        END
      END
    END
  END
END
`

	assert.Equal(t, strings.TrimSpace(f.QueryString()), strings.TrimSpace(exp))
}

// Ensure that a funnel query can be executed.
func TestFunnelQuery(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			a := &Account{}
			f := &Funnel{
				Name: "FUN",
				Steps: []*FunnelStep{
					{Condition: "resource == '/home'"},
					{Condition: "resource == '/signup'"},
					{Condition: "resource == '/checkout'"},
				},
			}
			assert.NoError(t, tx.CreateAccount(a))
			assert.NoError(t, a.CreateFunnel(f))
			a.Reset()

			// Track: "john" completes the whole checkout.
			assert.NoError(t, a.Track(newTestEvent("2000-01-01T00:00:00Z", "john", "", "web", "/home", "view", nil)))
			assert.NoError(t, a.Track(newTestEvent("2000-01-01T00:00:30Z", "john", "", "web", "/about", "view", nil)))
			assert.NoError(t, a.Track(newTestEvent("2000-01-01T00:01:00Z", "john", "", "web", "/signup", "view", nil)))
			assert.NoError(t, a.Track(newTestEvent("2000-01-01T00:02:00Z", "john", "", "web", "/checkout", "view", nil)))

			// Track: "susy" only completes the first step.
			assert.NoError(t, a.Track(newTestEvent("2000-01-02T00:00:00Z", "susy", "", "web", "/home", "view", nil)))

			// Track: "jim" completes the whole checkout but not in one session.
			assert.NoError(t, a.Track(newTestEvent("2000-01-01T00:00:00Z", "jim", "", "web", "/home", "view", nil)))
			assert.NoError(t, a.Track(newTestEvent("2000-01-01T00:01:00Z", "jim", "", "web", "/signup", "view", nil)))
			assert.NoError(t, a.Track(newTestEvent("2000-01-10T00:00:00Z", "jim", "", "web", "/checkout", "view", nil)))

			// Execute funnel query.
			results, err := f.Query()
			assert.NoError(t, err)
			if assert.NotNil(t, results) && assert.Equal(t, len(results.Steps), 3) {
				assert.Equal(t, results.Name, "FUN")
				assert.Equal(t, results.Steps[0].Condition, "resource == '/home'")
				assert.Equal(t, results.Steps[0].Count, 3) // john, susy, jim
				assert.Equal(t, results.Steps[1].Condition, "resource == '/signup'")
				assert.Equal(t, results.Steps[1].Count, 2) // john, jim
				assert.Equal(t, results.Steps[2].Condition, "resource == '/checkout'")
				assert.Equal(t, results.Steps[2].Count, 1) // john
			}

			return nil
		})
	})
}
