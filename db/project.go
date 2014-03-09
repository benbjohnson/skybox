package db

import (
	"fmt"
	"sort"

	"github.com/nu7hatch/gouuid"
	"github.com/skydb/gosky"
)

var (
	// ErrProjectNotFound is returned when a project does not exist.
	ErrProjectNotFound = &Error{"project not found", nil}

	// ErrProjectNameRequired is returned when a project has a blank name.
	ErrProjectNameRequired = &Error{"project name required", nil}
)

// schema defines the required properties on the project's sky table.
var schema = []*sky.Property{
	{Name: "channel", Transient: true, DataType: sky.Factor},
	{Name: "resource", Transient: true, DataType: sky.Factor},
	{Name: "action", Transient: true, DataType: sky.Factor},
	{Name: "domain", Transient: true, DataType: sky.Factor},
	{Name: "path", Transient: true, DataType: sky.Factor},
}

// Project represents a collection of Persons and their events.
// A Project belongs to an Account.
type Project struct {
	Tx        *Tx
	id        int
	AccountID int    `json:"accountID"`
	Name      string `json:"name"`
	APIKey    string `json:"apiKey"`
}

// ID returns the project identifier.
func (p *Project) ID() int {
	return p.id
}

// SkyTableName returns the name of the table used by Sky.
func (p *Project) SkyTableName() string {
	assert(p.id > 0, "uninitialized project does not have a sky table")
	return fmt.Sprintf("skybox-%d", p.id)
}

// SkyTable returns a reference to the table used by Sky.
func (p *Project) SkyTable() *sky.Table {
	return &sky.Table{
		Client: &p.Tx.db.SkyClient,
		Name:   p.SkyTableName(),
	}
}

// Validate validates all fields of the user.
func (p *Project) Validate() error {
	if len(p.Name) == 0 {
		return ErrProjectNameRequired
	}
	return nil
}

func (p *Project) get() ([]byte, error) {
	value := p.Tx.Bucket("projects").Get(itob(p.id))
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
		if err := p.GenerateAPIKey(); err != nil {
			return err
		}
	}

	return p.Tx.Bucket("projects").Put(itob(p.id), marshal(p))
}

// Delete removes the Project from the database.
func (p *Project) Delete() error {
	// Remove project entry.
	err := p.Tx.Bucket("projects").Delete(itob(p.id))
	assert(err == nil, "project delete error: %s", err)

	// Remove project id from indices.
	removeFromForeignKeyIndex(p.Tx, "account.projects", itob(p.AccountID), p.id)

	return nil
}

// GenerateAPIKey creates a new API key for a project.
func (p *Project) GenerateAPIKey() error {
	// Remove old API key from index.
	if p.APIKey != "" {
		removeFromUniqueIndex(p.Tx, "projects.APIKey", []byte(p.APIKey))
	}

	// Generate new API key.
	apiKey, err := uuid.NewV4()
	if err != nil {
		return err
	}
	p.APIKey = apiKey.String()

	// Update index.
	insertIntoUniqueIndex(p.Tx, "projects.APIKey", []byte(p.APIKey), p.ID())

	return nil
}

// Funnel retrieves a funnel with a given ID.
// Only funnels associated with this project will be returned.
func (p *Project) Funnel(id int) (*Funnel, error) {
	assert(p.id > 0, "find funnel on unsaved project: %d", p.id)
	f, err := p.Tx.Funnel(id)
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

	f.Tx = p.Tx
	f.ProjectID = p.id

	// Generate new id.
	f.id, _ = p.Tx.Bucket("funnels").NextSequence()
	assert(p.id > 0, "funnel sequence error")

	// Add funnel id to secondary index.
	insertIntoForeignKeyIndex(p.Tx, "project.funnels", itob(p.id), f.id)

	// Save funnel.
	return f.Save()
}

// Funnels retrieves a list of all funnels for the project.
func (p *Project) Funnels() (Funnels, error) {
	funnels := make(Funnels, 0)
	index := getForeignKeyIndex(p.Tx, "project.funnels", itob(p.id))

	for _, id := range index {
		f := &Funnel{Tx: p.Tx, id: id}
		err := f.Load()
		assert(err == nil, "funnel (%d) not found from project.funnels index (%d)", f.id, p.id)
		assert(f.ProjectID == p.id, "funnel/project mismatch: %d (%d) not in %d", f.id, f.ProjectID, p.id)
		funnels = append(funnels, f)
	}
	sort.Sort(funnels)
	return funnels, nil
}

// Track sends a single event to Sky.
func (p *Project) Track(e *Event) error {
	t := p.SkyTable()

	// Serialize event and insert event into Sky.
	id := e.ID()
	skyEvent := e.Serialize()
	if err := t.InsertEvent(id, skyEvent); err != nil {
		return err
	}

	// TODO(benbjohnson): Merge timelines if necessary.
	// t.Merge(id, t.DeviceID

	return nil
}

// Events returns a list of events for a given ID.
func (p *Project) Events(id string) ([]*Event, error) {
	t := p.SkyTable()
	skyEvents, err := t.Events(id)
	if err != nil {
		return nil, err
	}

	events := make([]*Event, 0, len(skyEvents))
	for _, skyEvent := range skyEvents {
		event := &Event{}
		event.Deserialize(id, skyEvent)
		events = append(events, event)
	}

	return events, nil
}

// Migrate creates a Sky table and updates the schema, if necessary.
func (p *Project) Migrate() error {
	name := p.SkyTableName()
	if err := p.createSkyTableIfNotExists(); err != nil {
		return fmt.Errorf("migrate table error: %s: %s", name, err)
	}
	for _, property := range schema {
		if err := p.createSkyPropertyIfNotExists(property); err != nil {
			return fmt.Errorf("migrate property error: %s: %s: %s", name, property.Name, err)
		}
	}
	return nil
}

// Reset drops the sky table and recreates it.
func (p *Project) Reset() error {
	c := &p.Tx.db.SkyClient
	c.DeleteTable(p.SkyTableName())
	return p.Migrate()
}

func (p *Project) createSkyTableIfNotExists() error {
	c := &p.Tx.db.SkyClient
	if t, err := c.Table(p.SkyTableName()); t != nil && err == nil {
		return err
	}
	return c.CreateTable(p.SkyTable())
}

func (p *Project) createSkyPropertyIfNotExists(property *sky.Property) error {
	t := p.SkyTable()
	if tmp, err := t.Property(property.Name); tmp != nil && err == nil {
		return err
	}
	return t.CreateProperty(property)
}

type Projects []*Project

func (s Projects) Len() int           { return len(s) }
func (s Projects) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Projects) Less(i, j int) bool { return s[i].Name < s[j].Name }
