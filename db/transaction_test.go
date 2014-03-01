package db_test

import (
	"testing"

	. "github.com/benbjohnson/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that the database can create an account.
func TestDBCreateAccount(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Create an account.
			a := &Account{Name: "Foo"}
			err := txn.CreateAccount(a)
			assert.NoError(t, err)
			assert.Equal(t, txn, a.Transaction)
			assert.Equal(t, a.ID(), 1)

			// Retrieve the account.
			a2, err := txn.Account(1)
			if assert.NoError(t, err) && assert.NotNil(t, a2) {
				assert.Equal(t, txn, a2.Transaction)
				assert.Equal(t, a2.Name, "Foo")
			}
			assert.True(t, a != a2)
			return nil
		})
	})
}

// Ensure that the database will reject an invalid account.
func TestDBCreateInvalidAccount(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			err := txn.CreateAccount(&Account{})
			assert.Equal(t, err, ErrAccountNameRequired)
			return nil
		})
	})
}

// Ensure that retrieving a missing account returns an error.
func TestDBAccountNotFound(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			a, err := txn.Account(1)
			assert.Equal(t, err, ErrAccountNotFound)
			assert.Nil(t, a)
			return nil
		})
	})
}

// Ensure that the database can return all accounts.
func TestDBAccounts(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			txn.CreateAccount(&Account{Name: "Foo"})
			txn.CreateAccount(&Account{Name: "Bar"})
			txn.CreateAccount(&Account{Name: "Baz"})

			// Retrieve the accounts.
			accounts, err := txn.Accounts()
			if assert.NoError(t, err) && assert.Equal(t, len(accounts), 3) {
				assert.Equal(t, accounts[0].Transaction, txn)
				assert.Equal(t, accounts[0].ID(), 2)
				assert.Equal(t, accounts[0].Name, "Bar")

				assert.Equal(t, accounts[1].Transaction, txn)
				assert.Equal(t, accounts[1].ID(), 3)
				assert.Equal(t, accounts[1].Name, "Baz")

				assert.Equal(t, accounts[2].Transaction, txn)
				assert.Equal(t, accounts[2].ID(), 1)
				assert.Equal(t, accounts[2].Name, "Foo")
			}
			return nil
		})
	})
}

// Ensure that retrieving a missing user returns an error.
func TestDBUserNotFound(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			u, err := txn.User(1)
			assert.Equal(t, err, ErrUserNotFound)
			assert.Nil(t, u)
			return nil
		})
	})
}

// Ensure that the database can retrieve a user by username.
func TestDBUserByUsername(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Add account and users.
			a := &Account{Name: "foo"}
			assert.NoError(t, txn.CreateAccount(a))
			assert.NoError(t, a.CreateUser(&User{Username: "johndoe", Password: "password"}))
			assert.NoError(t, a.CreateUser(&User{Username: "susyque", Password: "password"}))

			// Find user.
			u, _ := txn.UserByUsername("susyque")
			assert.Equal(t, u.ID(), 2)

			// Delete user and find.
			assert.NoError(t, u.Delete())
			_, err := txn.UserByUsername("susyque")
			assert.Equal(t, err, ErrUserNotFound)

			// Re-add and find again.
			assert.NoError(t, a.CreateUser(&User{Username: "susyque", Password: "foobar"}))
			u, _ = txn.UserByUsername("susyque")
			assert.Equal(t, u.ID(), 3)
			return nil
		})
	})
}
