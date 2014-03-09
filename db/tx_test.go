package db_test

import (
	"testing"

	. "github.com/skybox/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that the database can create an account.
func TestDBCreateAccount(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			// Create an account.
			a := &Account{}
			err := tx.CreateAccount(a)
			assert.NoError(t, err)
			assert.Equal(t, tx, a.Tx)
			assert.Equal(t, a.ID(), 1)

			// Retrieve the account.
			a2, err := tx.Account(1)
			if assert.NoError(t, err) && assert.NotNil(t, a2) {
				assert.Equal(t, tx, a2.Tx)
				assert.Equal(t, a2.ID(), 1)
			}
			assert.True(t, a != a2)
			return nil
		})
	})
}

// Ensure that retrieving a missing account returns an error.
func TestDBAccountNotFound(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			a, err := tx.Account(1)
			assert.Equal(t, err, ErrAccountNotFound)
			assert.Nil(t, a)
			return nil
		})
	})
}

// Ensure that the database can return all accounts.
func TestDBAccounts(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			tx.CreateAccount(&Account{})
			tx.CreateAccount(&Account{})
			tx.CreateAccount(&Account{})

			// Retrieve the accounts.
			accounts, err := tx.Accounts()
			if assert.NoError(t, err) && assert.Equal(t, len(accounts), 3) {
				assert.Equal(t, accounts[0].Tx, tx)
				assert.Equal(t, accounts[0].ID(), 1)

				assert.Equal(t, accounts[1].Tx, tx)
				assert.Equal(t, accounts[1].ID(), 2)

				assert.Equal(t, accounts[2].Tx, tx)
				assert.Equal(t, accounts[2].ID(), 3)
			}
			return nil
		})
	})
}

// Ensure that retrieving a missing user returns an error.
func TestDBUserNotFound(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			u, err := tx.User(1)
			assert.Equal(t, err, ErrUserNotFound)
			assert.Nil(t, u)
			return nil
		})
	})
}

// Ensure that the database can retrieve a user by email.
func TestDBUserByEmail(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(tx *Tx) error {
			// Add account and users.
			a := &Account{}
			assert.NoError(t, tx.CreateAccount(a))
			assert.NoError(t, a.CreateUser(&User{Email: "johndoe@gmail.com", Password: "password"}))
			assert.NoError(t, a.CreateUser(&User{Email: "susyque@gmail.com", Password: "password"}))

			// Find user.
			u, _ := tx.UserByEmail("susyque@gmail.com")
			assert.Equal(t, u.ID(), 2)

			// Delete user and find.
			assert.NoError(t, u.Delete())
			_, err := tx.UserByEmail("susyque@gmail.com")
			assert.Equal(t, err, ErrUserNotFound)

			// Re-add and find again.
			assert.NoError(t, a.CreateUser(&User{Email: "susyque@gmail.com", Password: "foobar"}))
			u, _ = tx.UserByEmail("susyque@gmail.com")
			assert.Equal(t, u.ID(), 3)
			return nil
		})
	})
}
