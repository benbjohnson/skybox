package db_test

import (
	"testing"

	. "github.com/benbjohnson/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that a project can create a funnel.
func TestFunnelCreate(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			// Create an account, project, and funnel.
			a := &Account{}
			assert.NoError(t, tx.CreateAccount(a))
			p := &Project{Name: "Project X"}
			assert.NoError(t, a.CreateProject(p))
			f := &Funnel{Name: "Funnel Y", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}
			assert.NoError(t, p.CreateFunnel(f))
			assert.Equal(t, f.ID(), 1)

			// Retrieve the funnel.
			f2, err := tx.Funnel(1)
			if assert.NoError(t, err) && assert.NotNil(t, f2) {
				assert.Equal(t, f2.Tx, tx)
				assert.Equal(t, f2.ID(), 1)
				assert.Equal(t, f2.ProjectID, 1)
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
			p := &Project{Name: "Project X"}
			assert.NoError(t, a.CreateProject(p))
			assert.Equal(t, p.CreateFunnel(&Funnel{Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}), ErrFunnelNameRequired)
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
			p := &Project{Name: "Project X"}
			assert.NoError(t, a.CreateProject(p))
			assert.Equal(t, p.CreateFunnel(&Funnel{Name: "Funnel Y"}), ErrFunnelStepsRequired)
			return nil
		})
	})
}

// Ensure that a funnel can update itself.
func TestFunnelUpdate(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			// Create account and project.
			a := &Account{}
			assert.NoError(t, tx.CreateAccount(a))
			p := &Project{Name: "Project X"}
			assert.NoError(t, a.CreateProject(p))
			f := &Funnel{Name: "Funnel Y", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}
			assert.NoError(t, p.CreateFunnel(f))

			// Update the funnel.
			f.Name = "Funnel Z"
			f.Save()

			// Retrieve the project.
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
			// Create account, project, and funnel.
			a := &Account{}
			assert.NoError(t, tx.CreateAccount(a))
			p := &Project{Name: "Project X"}
			assert.NoError(t, a.CreateProject(p))
			f := &Funnel{Name: "Project Y", Steps: []*FunnelStep{{Condition: "action == 'foo'"}}}
			assert.NoError(t, p.CreateFunnel(f))

			// Delete the funnel.
			assert.NoError(t, f.Delete())

			// Retrieve the funnel again.
			_, err := tx.Funnel(1)
			assert.Equal(t, err, ErrFunnelNotFound)
			return nil
		})
	})
}
