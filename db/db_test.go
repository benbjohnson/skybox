package db_test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/benbjohnson/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that the database can create an account.
func TestDBCreateAccount(t *testing.T) {
	withDB(func(db *DB) {
		// Create an account.
		a := &Account{Name: "Foo"}
		err := db.CreateAccount(a)
		assert.NoError(t, err)
		assert.Equal(t, db, a.DB())
		assert.Equal(t, a.ID(), 1)

		// Retrieve the account.
		a2, err := db.Account(1)
		if assert.NoError(t, err) && assert.NotNil(t, a2) {
			assert.Equal(t, db, a2.DB())
			assert.Equal(t, a2.Name, "Foo")
		}
		assert.True(t, a != a2)
	})
}

// Ensure that the database will reject an invalid account.
func TestDBCreateInvalidAccount(t *testing.T) {
	withDB(func(db *DB) {
		err := db.CreateAccount(&Account{})
		assert.Equal(t, err, ErrAccountNameRequired)
	})
}

// Ensure that retrieving a missing account returns an error.
func TestDBAccountNotFound(t *testing.T) {
	withDB(func(db *DB) {
		a, err := db.Account(1)
		assert.Equal(t, err, ErrAccountNotFound)
		assert.Nil(t, a)
	})
}

// Ensure that the database can return all accounts.
func TestDBAccounts(t *testing.T) {
	withDB(func(db *DB) {
		db.CreateAccount(&Account{Name: "Foo"})
		db.CreateAccount(&Account{Name: "Bar"})
		db.CreateAccount(&Account{Name: "Baz"})

		// Retrieve the accounts.
		accounts, err := db.Accounts()
		if assert.NoError(t, err) && assert.Equal(t, len(accounts), 3) {
			assert.Equal(t, accounts[0].DB(), db)
			assert.Equal(t, accounts[0].ID(), 2)
			assert.Equal(t, accounts[0].Name, "Bar")

			assert.Equal(t, accounts[1].DB(), db)
			assert.Equal(t, accounts[1].ID(), 3)
			assert.Equal(t, accounts[1].Name, "Baz")

			assert.Equal(t, accounts[2].DB(), db)
			assert.Equal(t, accounts[2].ID(), 1)
			assert.Equal(t, accounts[2].Name, "Foo")
		}
	})
}

// Ensure that retrieving a missing user returns an error.
func TestDBUserNotFound(t *testing.T) {
	withDB(func(db *DB) {
		u, err := db.User(1)
		assert.Equal(t, err, ErrUserNotFound)
		assert.Nil(t, u)
	})
}

// Ensure that the database can retrieve a user by username.
func TestDBUserByUsername(t *testing.T) {
	withDB(func(db *DB) {
		// Add account and users.
		a := &Account{Name: "foo"}
		assert.NoError(t, db.CreateAccount(a))
		assert.NoError(t, a.CreateUser(&User{Username: "johndoe", Password: "password"}))
		assert.NoError(t, a.CreateUser(&User{Username: "susyque", Password: "password"}))

		// Find user.
		u, _ := db.UserByUsername("susyque")
		assert.Equal(t, u.ID(), 2)

		// Delete user and find.
		assert.NoError(t, u.Delete())
		_, err := db.UserByUsername("susyque")
		assert.Equal(t, err, ErrUserNotFound)

		// Re-add and find again.
		assert.NoError(t, a.CreateUser(&User{Username: "susyque", Password: "foobar"}))
		u, _ = db.UserByUsername("susyque")
		assert.Equal(t, u.ID(), 3)
	})
}

// withDB executes a function with an open database.
func withDB(fn func(*DB)) {
	f, _ := ioutil.TempFile("", "skybox-")
	path := f.Name()
	f.Close()
	os.Remove(path)
	defer os.RemoveAll(path)

	var db DB
	if err := db.Open(path, 0666); err != nil {
		panic("db open error: " + err.Error())
	}
	defer db.Close()
	fn(&db)
}
