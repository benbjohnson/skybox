package db_test

import (
	"strings"
	"testing"

	. "github.com/benbjohnson/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that an account can create a user.
func TestUserCreate(t *testing.T) {
	withDB(func(db *DB) {
		// Create an account and user.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		u := &User{Username: "johndoe", Password: "mybirthday"}
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
func TestUserCreateAfterDeletion(t *testing.T) {
	withDB(func(db *DB) {
		// Create an account and delete it.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		assert.NoError(t, a.Delete())

		// Attempt to create a user.
		err := a.CreateUser(&User{Username: "johndoe", Password: "password"})
		assert.Equal(t, err, ErrAccountNotFound)
	})
}

// Ensure that creating an invalid user returns an error.
func TestUserCreateMissingUsername(t *testing.T) {
	withDB(func(db *DB) {
		// Create an account and user.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		err := a.CreateUser(&User{Password: "password"})
		assert.Equal(t, err, ErrUserUsernameRequired)
	})
}

// Ensure that creating a user without a password returns an error.
func TestUserCreateMissingPassword(t *testing.T) {
	withDB(func(db *DB) {
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		err := a.CreateUser(&User{Username: "johndoe"})
		assert.Equal(t, err, ErrUserPasswordRequired)
	})
}

// Ensure that creating a user with a short password returns an error.
func TestUserCreatePasswordTooShort(t *testing.T) {
	withDB(func(db *DB) {
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		err := a.CreateUser(&User{Username: "johndoe", Password: "abc"})
		assert.Equal(t, err, ErrUserPasswordTooShort)
	})
}

// Ensure that creating a user with a long password returns an error.
func TestUserCreatePasswordTooLong(t *testing.T) {
	withDB(func(db *DB) {
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		err := a.CreateUser(&User{Username: "johndoe", Password: strings.Repeat("*", 51)})
		assert.Equal(t, err, ErrUserPasswordTooLong)
	})
}

// Ensure that creating a user with an already taken username returns an error.
func TestUserCreateUsernameTaken(t *testing.T) {
	withDB(func(db *DB) {
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		err := a.CreateUser(&User{Username: "johndoe", Password: "password"})
		assert.NoError(t, err)
		err = a.CreateUser(&User{Username: "johndoe", Password: "foobar"})
		assert.Equal(t, err, ErrUserUsernameTaken)
	})
}

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
