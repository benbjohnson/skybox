package db

import (
	"sort"

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

// Funnel retrieves a funnel with a given ID.
// Only funnels associated with this project will be returned.
func (p *Project) Funnel(id int) (*Funnel, error) {
	assert(p.id > 0, "find funnel on unsaved project: %d", p.id)
	f, err := p.Transaction.Funnel(id)
	if err != nil {
		return nil, err
	} else if f.ProjectID != p.ID() {
		return nil, ErrFunnelNotFound
	}
	return f, nil
}

// CreateFunnel creates a new Funnel for this project.
func (p *Project) CreateFunnel(f *Funnel) error {
	assert(f.id == 0, "create funnel with a non-zero id: %d", f.id)
	assert(p.id > 0, "create funnel on unsaved project: %d", p.id)
	if err := f.Validate(); err != nil {
		return err
	}

	// Verify project exists.
	if _, err := p.get(); err != nil {
		return err
	}

	f.Transaction = p.Transaction
	f.ProjectID = p.id

	// Generate new id.
	f.id, _ = p.Transaction.Bucket("funnels").NextSequence()
	assert(p.id > 0, "funnel sequence error")

	// Add funnel id to secondary index.
	insertIntoForeignKeyIndex(p.Transaction, "project.funnels", itob(p.id), f.id)

	// Save funnel.
	return f.Save()
}

// Funnels retrieves a list of all funnels for the project.
func (p *Project) Funnels() (Funnels, error) {
	funnels := make(Funnels, 0)
	index := getForeignKeyIndex(p.Transaction, "project.funnels", itob(p.id))

	for _, id := range index {
		f := &Funnel{Transaction: p.Transaction, id: id}
		err := f.Load()
		assert(err == nil, "funnel (%d) not found from project.funnels index (%d)", f.id, p.id)
		assert(f.ProjectID == p.id, "funnel/project mismatch: %d (%d) not in %d", f.id, f.ProjectID, p.id)
		funnels = append(funnels, f)
	}
	sort.Sort(funnels)
	return funnels, nil
}

type Projects []*Project

func (s Projects) Len() int           { return len(s) }
func (s Projects) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Projects) Less(i, j int) bool { return s[i].Name < s[j].Name }
