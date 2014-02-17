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

// Ensure that an account can create a user.
func TestAccountCreateUser(t *testing.T) {
	withDB(func(db *DB) {
		// Create an account and user.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		u := &User{Username: "johndoe"}
		assert.NoError(t, a.CreateUser(u))
		assert.Equal(t, u.Id(), 1)

		// Retrieve the user.
		u2, err := db.User(1)
		if assert.NoError(t, err) && assert.NotNil(t, u2) {
			assert.Equal(t, u2.DB(), db)
			assert.Equal(t, u2.Id(), 1)
			assert.Equal(t, u2.AccountId, 1)
			assert.Equal(t, u2.Username, "johndoe")
		}
	})
}

// Ensure that an account cannot create a user after it's deleted.
func TestAccountCreateUserAfterDeletion(t *testing.T) {
	withDB(func(db *DB) {
		// Create an account and delete it.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		assert.NoError(t, a.Delete())

		// Attempt to create a user.
		err := a.CreateUser(&User{Username: "johndoe"})
		assert.Equal(t, err, ErrAccountNotFound)
	})
}

// Ensure that creating an invalid user returns an error.
func TestAccountCreateUserMissingUsername(t *testing.T) {
	withDB(func(db *DB) {
		// Create an account and user.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		err := a.CreateUser(&User{})
		assert.Equal(t, err, ErrUserUsernameRequired)
	})
}

// Ensure that an account can retrieve all associated users.
func TestAccountUsers(t *testing.T) {
	withDB(func(db *DB) {
		// Create two accounts.
		a1 := &Account{Name: "foo"}
		assert.NoError(t, db.CreateAccount(a1))
		a2 := &Account{Name: "bar"}
		assert.NoError(t, db.CreateAccount(a2))

		// Add users to first account.
		assert.NoError(t, a1.CreateUser(&User{Username: "johndoe"}))
		assert.NoError(t, a1.CreateUser(&User{Username: "susyque"}))

		// Add users to second account.
		assert.NoError(t, a2.CreateUser(&User{Username: "billybob"}))

		// Check first account users.
		users, err := a1.Users()
		if assert.NoError(t, err) && assert.Equal(t, len(users), 2) {
			assert.Equal(t, users[0].DB(), db)
			assert.Equal(t, users[0].Id(), 1)
			assert.Equal(t, users[0].AccountId, 1)
			assert.Equal(t, users[0].Username, "johndoe")

			assert.Equal(t, users[1].DB(), db)
			assert.Equal(t, users[1].Id(), 2)
			assert.Equal(t, users[1].AccountId, 1)
			assert.Equal(t, users[1].Username, "susyque")
		}

		// Check second account users.
		users, err = a2.Users()
		if assert.NoError(t, err) && assert.Equal(t, len(users), 1) {
			assert.Equal(t, users[0].DB(), db)
			assert.Equal(t, users[0].Id(), 3)
			assert.Equal(t, users[0].AccountId, 2)
			assert.Equal(t, users[0].Username, "billybob")
		}
	})
}
