package db_test

import (
	"testing"

	. "github.com/benbjohnson/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that an account can create a project.
func TestProjectCreate(t *testing.T) {
	withDB(func(db *DB) {
		// Create an account and project.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		u := &Project{Name: "Project X"}
		assert.NoError(t, a.CreateProject(u))
		assert.Equal(t, u.ID(), 1)

		// Retrieve the project.
		u2, err := db.Project(1)
		if assert.NoError(t, err) && assert.NotNil(t, u2) {
			assert.Equal(t, u2.DB(), db)
			assert.Equal(t, u2.ID(), 1)
			assert.Equal(t, u2.AccountID, 1)
			assert.Equal(t, u2.Name, "Project X")
		}
	})
}

// Ensure that an account cannot create a project after it's deleted.
func TestProjectCreateAfterDeletion(t *testing.T) {
	withDB(func(db *DB) {
		// Create an account and delete it.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		assert.NoError(t, a.Delete())

		// Attempt to create a project.
		err := a.CreateProject(&Project{Name: "Project X"})
		assert.Equal(t, err, ErrAccountNotFound)
	})
}

// Ensure that creating an invalid project returns an error.
func TestProjectCreateMissingName(t *testing.T) {
	withDB(func(db *DB) {
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		assert.Equal(t, a.CreateProject(&Project{}), ErrProjectNameRequired)
	})
}

// Ensure that a project can update itself.
func TestProjectUpdate(t *testing.T) {
	withDB(func(db *DB) {
		// Create account and project.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		p := &Project{Name: "Project X"}
		assert.NoError(t, a.CreateProject(p))

		// Update the project.
		p.Name = "Project Y"
		p.Save()

		// Retrieve the project.
		p2, err := db.Project(1)
		if assert.NoError(t, err) && assert.NotNil(t, p2) {
			assert.Equal(t, p2.Name, "Project Y")
		}
	})
}

// Ensure that a user can be deleted.
func TestProjectDelete(t *testing.T) {
	withDB(func(db *DB) {
		// Create account and project.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		p := &Project{Name: "Project X"}
		assert.NoError(t, a.CreateProject(p))

		// Delete the project.
		assert.NoError(t, p.Delete())

		// Retrieve the project again.
		_, err := db.Project(1)
		assert.Equal(t, err, ErrProjectNotFound)
	})
}
