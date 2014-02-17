package db

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/boltdb/bolt"
)

var (
	// ErrUserNotFound is returned when a user does not exist.
	ErrUserNotFound = &Error{"user not found", nil}

	// ErrUserUsernameRequired is returned when a user has a blank username.
	ErrUserUsernameRequired = &Error{"user username required", nil}

	// ErrUserUsernameTaken is returned when a username already exists.
	ErrUserUsernameTaken = &Error{"user username is already taken", nil}

	// ErrUserPasswordRequired is returned when a user has a blank password.
	ErrUserPasswordRequired = &Error{"user password required", nil}

	// ErrUserPasswordTooShort is returned when a password is too short.
	ErrUserPasswordTooShort = &Error{"user password too short", nil}

	// ErrUserPasswordTooLong is returned when a password is too long.
	ErrUserPasswordTooLong = &Error{"user password too long", nil}

	// ErrUserNotAuthenticated is returned when a password doesn't match the hash.
	ErrUserNotAuthenticated = &Error{"user not authenticated", nil}
)

const (
	// MinPasswordLength is the shortest a password can be.
	MinPasswordLength = 6

	// MaxPasswordLength is the longest a password can be.
	MaxPasswordLength = 50
)

// User represents a user within the system.
// A User belongs to an Account and can access all Projects within the Account.
type User struct {
	db        *DB
	id        int
	AccountId int    `json:"accountId"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Hash      []byte `json:"hash"`
	Email     string `json:"email"`
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
	} else if u.id == 0 && len(u.Password) == 0 {
		return ErrUserPasswordRequired
	} else if len(u.Password) < MinPasswordLength {
		return ErrUserPasswordTooShort
	} else if len(u.Password) > MaxPasswordLength {
		return ErrUserPasswordTooLong
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

	// Remove user id from indices.
	removeFromForeignKeyIndex(txn, "account.users", itob(u.AccountId), u.id)
	removeFromUniqueIndex(txn, "user.username", []byte(u.Username))

	return nil
}

// GenerateHash generates a hashed password from the currently set password.
func (u *User) GenerateHash() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 0)
	if err != nil {
		return err
	}
	u.Hash = hash
	return nil
}

// Authenticate checks if a plaintext password matches the hash.
func (u *User) Authenticate(password string) error {
	if err := bcrypt.CompareHashAndPassword(u.Hash, []byte(password)); err != nil {
		return ErrUserNotAuthenticated
	}
	return nil
}

type Users []*User

func (s Users) Len() int           { return len(s) }
func (s Users) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Users) Less(i, j int) bool { return s[i].Username < s[j].Username }
