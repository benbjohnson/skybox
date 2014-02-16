package db

import (
// "github.com/boltdb/bolt"
)

var (
	// ErrAccountNameRequired is returned when an account has a blank name.
	ErrAccountNameRequired = &Error{"account name required", nil}
)

// Account represents a collection of Users and Projects.
type Account struct {
	db   *DB
	Id   int
	Name string
}

// Accounts is a list of Account objects.
type Accounts []*Account

// DB returns the database that created the account.
func (a *Account) DB() *DB {
	return a.db
}

// Validate validates all fields of the account.
func (a *Account) Validate() error {
	if len(a.Name) == 0 {
		return ErrAccountNameRequired
	}
	return nil
}

// Update updates all fields in the account using a map.
func (a *Account) Update(values map[string]interface{}) error {
	return nil // TODO
}

// Delete removes the account from the database.
func (a *Account) Delete() error {
	// TODO: Remove all projects.
	// TODO: Remove all users.
	return nil // TODO
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
