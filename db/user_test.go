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
		u := &User{Username: "bob", Password: "password"}
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
		u := &User{Username: "bob", Password: "password"}
		assert.NoError(t, a.CreateUser(u))

		// Delete the user.
		assert.NoError(t, u.Delete())

		// Retrieve the user again.
		_, err := db.User(1)
		assert.Equal(t, err, ErrUserNotFound)
	})
}

// Ensure that a user can be authenticated
func TestUserAuthenticate(t *testing.T) {
	withDB(func(db *DB) {
		// Create account and user.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		u := &User{Username: "bob", Password: "password"}
		assert.NoError(t, a.CreateUser(u))

		// Authenticate the user with the correct password.
		assert.Nil(t, u.Authenticate("password"))

		// Return error if authenticating with the wrong password.
		assert.Equal(t, u.Authenticate("not_the_right_password"), ErrUserNotAuthenticated)
	})
}
