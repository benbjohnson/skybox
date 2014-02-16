package db_test

import (
	"testing"

	. "github.com/benbjohnson/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that an account can update itself.
func TestAccountUpdate(t *testing.T) {
	withDB(func(db *DB) {
		// Create and update account.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		a.Name = "Bar"
		assert.NoError(t, a.Save())

		// Retrieve the account.
		a2, err := db.Account(1)
		if assert.NoError(t, err) && assert.NotNil(t, a2) {
			assert.Equal(t, a2.Name, "Bar")
		}
	})
}

// Ensure that an account can be deleted.
func TestAccountDelete(t *testing.T) {
	withDB(func(db *DB) {
		// Create account.
		assert.NoError(t, db.CreateAccount(&Account{Name: "Foo"}))

		// Retrieve and delete account.
		a, _ := db.Account(1)
		assert.NoError(t, a.Delete())

		// Retrieve the account again.
		_, err := db.Account(1)
		assert.Equal(t, err, ErrAccountNotFound)
	})
}
