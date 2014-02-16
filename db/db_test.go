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
		assert.Equal(t, a.Id, 1)

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
