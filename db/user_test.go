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
		db.Do(func(txn *Transaction) error {
			// Create an account and user.
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			u := &User{Email: "johndoe@gmail.com", Password: "mybirthday"}
			assert.NoError(t, a.CreateUser(u))
			assert.Equal(t, u.ID(), 1)

			// Retrieve the user.
			u2, err := txn.User(1)
			if assert.NoError(t, err) && assert.NotNil(t, u2) {
				assert.Equal(t, u2.Transaction, txn)
				assert.Equal(t, u2.ID(), 1)
				assert.Equal(t, u2.AccountID, 1)
				assert.Equal(t, u2.Email, "johndoe@gmail.com")
			}
			return nil
		})
	})
}

// Ensure that an account cannot create a user after it's deleted.
func TestUserCreateAfterDeletion(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Create an account and delete it.
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			assert.NoError(t, a.Delete())

			// Attempt to create a user.
			err := a.CreateUser(&User{Email: "johndoe@gmail.com", Password: "password"})
			assert.Equal(t, err, ErrAccountNotFound)
			return nil
		})
	})
}

// Ensure that creating an invalid user returns an error.
func TestUserCreateMissingEmail(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Create an account and user.
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			err := a.CreateUser(&User{Password: "password"})
			assert.Equal(t, err, ErrUserEmailRequired)
			return nil
		})
	})
}

// Ensure that creating a user without a password returns an error.
func TestUserCreateMissingPassword(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			err := a.CreateUser(&User{Email: "johndoe@gmail.com"})
			assert.Equal(t, err, ErrUserPasswordRequired)
			return nil
		})
	})
}

// Ensure that creating a user with a short password returns an error.
func TestUserCreatePasswordTooShort(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			err := a.CreateUser(&User{Email: "johndoe@gmail.com", Password: "abc"})
			assert.Equal(t, err, ErrUserPasswordTooShort)
			return nil
		})
	})
}

// Ensure that creating a user with a long password returns an error.
func TestUserCreatePasswordTooLong(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			err := a.CreateUser(&User{Email: "johndoe@gmail.com", Password: strings.Repeat("*", 51)})
			assert.Equal(t, err, ErrUserPasswordTooLong)
			return nil
		})
	})
}

// Ensure that creating a user with an already taken email returns an error.
func TestUserCreateEmailTaken(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			err := a.CreateUser(&User{Email: "johndoe@gmail.com", Password: "password"})
			assert.NoError(t, err)
			err = a.CreateUser(&User{Email: "johndoe@gmail.com", Password: "foobar"})
			assert.Equal(t, err, ErrUserEmailTaken)
			return nil
		})
	})
}

// Ensure that a user can update itself.
func TestUserUpdate(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Create account and user.
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			u := &User{Email: "bob@gmail.com", Password: "password"}
			assert.NoError(t, a.CreateUser(u))

			// Update the user.
			u.Email = "jim@gmail.com"
			u.Save()

			// Retrieve the user.
			u2, err := txn.User(1)
			if assert.NoError(t, err) && assert.NotNil(t, u2) {
				assert.Equal(t, u2.Email, "jim@gmail.com")
			}
			return nil
		})
	})
}

// Ensure that an account can be deleted.
func TestUserDelete(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Create account and user.
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			u := &User{Email: "bob@gmail.com", Password: "password"}
			assert.NoError(t, a.CreateUser(u))

			// Delete the user.
			assert.NoError(t, u.Delete())

			// Retrieve the user again.
			_, err := txn.User(1)
			assert.Equal(t, err, ErrUserNotFound)
			return nil
		})
	})
}

// Ensure that a user can be authenticated
func TestUserAuthenticate(t *testing.T) {
	withDB(func(db *DB) {
		db.Do(func(txn *Transaction) error {
			// Create account and user.
			a := &Account{}
			assert.NoError(t, txn.CreateAccount(a))
			u := &User{Email: "bob@gmail.com", Password: "password"}
			assert.NoError(t, a.CreateUser(u))

			// Authenticate the user with the correct password.
			assert.Nil(t, u.Authenticate("password"))

			// Return error if authenticating with the wrong password.
			assert.Equal(t, u.Authenticate("not_the_right_password"), ErrUserNotAuthenticated)
			return nil
		})
	})
}
