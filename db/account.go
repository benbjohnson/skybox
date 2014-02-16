package db

import (
	"github.com/boltdb/bolt"
)

var (
	// ErrAccountNotFound is returned when an account with the given id does
	// not exist.
	ErrAccountNotFound = &Error{"account not found", nil}

	// ErrAccountNameRequired is returned when an account has a blank name.
	ErrAccountNameRequired = &Error{"account name required", nil}
)

// Account represents a collection of Users and Projects.
type Account struct {
	db   *DB
	id   int
	Name string
}

// DB returns the database that created the account.
func (a *Account) DB() *DB {
	return a.db
}

// Id returns the account identifier.
func (a *Account) Id() int {
	return a.id
}

// Validate validates all fields of the account.
func (a *Account) Validate() error {
	if len(a.Name) == 0 {
		return ErrAccountNameRequired
	}
	return nil
}

// Load retrieves an account from the database.
func (a *Account) Load() error {
	value, err := a.db.Get("accounts", itob(a.id))
	if err != nil {
		return err
	} else if value == nil {
		return ErrAccountNotFound
	}

	unmarshal(value, &a)
	return nil
}

// Save commits the Account to the database.
func (a *Account) Save() error {
	return a.db.Do(func(txn *bolt.RWTransaction) error {
		return a.SaveTo(txn)
	})
}

// SaveTo commits the Account to an open transaction.
func (a *Account) SaveTo(txn *bolt.RWTransaction) error {
	assert(a.id > 0, "uninitialized account cannot be saved")
	return txn.Put("accounts", itob(a.id), marshal(a))
}

// Delete removes the account from the database.
func (a *Account) Delete() error {
	return a.db.Do(func(txn *bolt.RWTransaction) error {
		return a.DeleteFrom(txn)
	})
}

// DeleteFrom removes the account from an open transaction.
func (a *Account) DeleteFrom(txn *bolt.RWTransaction) error {
	err := txn.Delete("accounts", itob(a.id))
	assert(err == nil, "account delete error: %s", err)

	// TODO: Remove all projects.
	// TODO: Remove all users.

	return nil
}

// Project a project within this account by id.
func (a *Account) Project(id int) (*Project, error) {
	return nil, nil // TODO
}

// Projects retrieves a list of projects for the account.
func (a *Account) Projects() (Projects, error) {
	return nil, nil // TODO
}

// CreateProject creates a new Project for this account.
func (a *Account) CreateProject(p *Project) error {
	return nil // TODO
}

type Accounts []*Account

func (s Accounts) Len() int           { return len(s) }
func (s Accounts) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Accounts) Less(i, j int) bool { return s[i].Name < s[j].Name }
