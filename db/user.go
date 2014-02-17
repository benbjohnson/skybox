package db

import (
	// "code.google.com/p/go.crypto/bcrypt"
	"github.com/boltdb/bolt"
)

var (
	// ErrUserNotFound is returned when a user does not exist.
	ErrUserNotFound = &Error{"user not found", nil}

	// ErrUserUsernameRequired is returned when an user has a blank name.
	ErrUserUsernameRequired = &Error{"user username required", nil}
)

// User represents a user within the system.
// A User belongs to an Account and can access all Projects within the Account.
type User struct {
	db        *DB
	id        int
	AccountId int
	Username  string
	Password  string
	Email     string
}

// DB returns the database that created the user.
func (u *User) DB() *DB {
	return u.db
}

// Id returns the user identifier.
func (u *User) Id() int {
	return u.id
}

// Validate validates all fields of the user.
func (u *User) Validate() error {
	if len(u.Username) == 0 {
		return ErrUserUsernameRequired
	}
	return nil
}

func (u *User) get(txn *bolt.Transaction) ([]byte, error) {
	value, err := txn.Get("users", itob(u.id))
	assert(err == nil, "get user error: %s", err)
	if value == nil {
		return nil, ErrUserNotFound
	}
	return value, nil
}

// Load retrieves a user from the database.
func (u *User) Load() error {
	return u.db.With(func(txn *bolt.Transaction) error {
		return u.load(txn)
	})
}

func (u *User) load(txn *bolt.Transaction) error {
	value, err := u.get(txn)
	if err != nil {
		return err
	}
	unmarshal(value, &u)
	return nil
}

// Save commits the User to the database.
func (u *User) Save() error {
	return u.db.Do(func(txn *bolt.RWTransaction) error {
		return u.save(txn)
	})
}

func (u *User) save(txn *bolt.RWTransaction) error {
	assert(u.id > 0, "uninitialized user cannot be saved")
	return txn.Put("users", itob(u.id), marshal(u))
}

// Delete removes the User from the database.
func (u *User) Delete() error {
	return u.db.Do(func(txn *bolt.RWTransaction) error {
		return u.del(txn)
	})
}

func (u *User) del(txn *bolt.RWTransaction) error {
	// Remove user entry.
	err := txn.Delete("users", itob(u.id))
	assert(err == nil, "user delete error: %s", err)

	// Remove user id from secondary index.
	removeFromIndex(txn, "account.users", itob(u.AccountId), u.id)

	return nil
}

type Users []*User

func (s Users) Len() int           { return len(s) }
func (s Users) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Users) Less(i, j int) bool { return s[i].Username < s[j].Username }
