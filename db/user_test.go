package db_test

import (
	"testing"

	. "github.com/benbjohnson/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that a user can update itself.
func TestUserUpdate(t *testing.T) {
	withDB(func(db *DB) {
		// Create account and user.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		u := &User{Username: "bob"}
		assert.NoError(t, a.CreateUser(u))

		// Update the user.
		u.Username = "jim"
		u.Save()

		// Retrieve the user.
		u2, err := db.User(1)
		if assert.NoError(t, err) && assert.NotNil(t, u2) {
			assert.Equal(t, u2.Username, "jim")
		}
	})
}

// Ensure that an account can be deleted.
func TestUserDelete(t *testing.T) {
	withDB(func(db *DB) {
		// Create account and user.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		u := &User{Username: "bob"}
		assert.NoError(t, a.CreateUser(u))

		// Delete the user.
		assert.NoError(t, u.Delete())

		// Retrieve the user again.
		_, err := db.User(1)
		assert.Equal(t, err, ErrUserNotFound)
	})
}
