package db

import (
	"github.com/nu7hatch/gouuid"
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
	Transaction *Transaction
	id          int
	AccountID   int    `json:"accountID"`
	Name        string `json:"name"`
	APIKey      string `json:"apiKey"`
}

// ID returns the project identifier.
func (p *Project) ID() int {
	return p.id
}

// Validate validates all fields of the user.
func (p *Project) Validate() error {
	if len(p.Name) == 0 {
		return ErrProjectNameRequired
	}
	return nil
}

func (p *Project) get() ([]byte, error) {
	value := p.Transaction.Bucket("projects").Get(itob(p.id))
	if value == nil {
		return nil, ErrProjectNotFound
	}
	return value, nil
}

// Load retrieves a project from the database.
func (p *Project) Load() error {
	value, err := p.get()
	if err != nil {
		return err
	}
	unmarshal(value, &p)
	return nil
}

// Save commits the Project to the database.
func (p *Project) Save() error {
	assert(p.id > 0, "uninitialized project cannot be saved")

	// Autogenerate an API key if one does not exist.
	if len(p.APIKey) == 0 {
		apiKey, err := uuid.NewV4()
		if err != nil {
			return err
		}
		p.APIKey = apiKey.String()
	}

	return p.Transaction.Bucket("projects").Put(itob(p.id), marshal(p))
}

// Delete removes the Project from the database.
func (p *Project) Delete() error {
	// Remove project entry.
	err := p.Transaction.Bucket("projects").Delete(itob(p.id))
	assert(err == nil, "project delete error: %s", err)

	// Remove project id from indices.
	removeFromForeignKeyIndex(p.Transaction, "account.projects", itob(p.AccountID), p.id)

	return nil
}

type Projects []*Project

func (s Projects) Len() int           { return len(s) }
func (s Projects) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Projects) Less(i, j int) bool { return s[i].Name < s[j].Name }
