package db

import (
	"github.com/boltdb/bolt"
)

var (
	// ErrProjectNotFound is returned when a project does not exist.
	ErrProjectNotFound = &Error{"project not found", nil}

	// ErrProjectNameRequired is returned when a project has a blank name.
	ErrProjectNameRequired = &Error{"project name required", nil}
)

// Project represents a collection of Persons and their events.
// A Project belongs to an Account.
type Project struct {
	db        *DB
	id        int
	AccountId int    `json:"accountId"`
	Name      string `json:"name"`
}

// DB returns the database that created the project.
func (p *Project) DB() *DB {
	return p.db
}

// Id returns the project identifier.
func (p *Project) Id() int {
	return p.id
}

// Validate validates all fields of the user.
func (p *Project) Validate() error {
	if len(p.Name) == 0 {
		return ErrProjectNameRequired
	}
	return nil
}

func (p *Project) get(txn *bolt.Transaction) ([]byte, error) {
	value, err := txn.Get("projects", itob(p.id))
	assert(err == nil, "get project error: %s", err)
	if value == nil {
		return nil, ErrProjectNotFound
	}
	return value, nil
}

// Load retrieves a project from the database.
func (p *Project) Load() error {
	return p.db.With(func(txn *bolt.Transaction) error {
		return p.load(txn)
	})
}

func (p *Project) load(txn *bolt.Transaction) error {
	value, err := p.get(txn)
	if err != nil {
		return err
	}
	unmarshal(value, &p)
	return nil
}

// Save commits the Project to the database.
func (p *Project) Save() error {
	return p.db.Do(func(txn *bolt.RWTransaction) error {
		return p.save(txn)
	})
}

func (p *Project) save(txn *bolt.RWTransaction) error {
	assert(p.id > 0, "uninitialized project cannot be saved")
	return txn.Put("projects", itob(p.id), marshal(p))
}

// Delete removes the Project from the database.
func (p *Project) Delete() error {
	return p.db.Do(func(txn *bolt.RWTransaction) error {
		return p.del(txn)
	})
}

func (p *Project) del(txn *bolt.RWTransaction) error {
	// Remove project entry.
	err := txn.Delete("projects", itob(p.id))
	assert(err == nil, "project delete error: %s", err)

	// Remove project id from indices.
	removeFromForeignKeyIndex(txn, "account.projects", itob(p.AccountId), p.id)

	return nil
}

type Projects []*Project

func (s Projects) Len() int           { return len(s) }
func (s Projects) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Projects) Less(i, j int) bool { return s[i].Name < s[j].Name }
