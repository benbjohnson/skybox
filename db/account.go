package db

import (
	"fmt"
	"sort"

	"github.com/nu7hatch/gouuid"
	"github.com/skydb/gosky"
)

var (
	// ErrAccountNotFound is returned when an account with the given id does
	// not exist.
	ErrAccountNotFound = &Error{"account not found", nil}
)

// schema defines the required properties on the account's sky table.
var schema = []*sky.Property{
	{Name: "channel", Transient: true, DataType: sky.Factor},
	{Name: "resource", Transient: true, DataType: sky.Factor},
	{Name: "action", Transient: true, DataType: sky.Factor},
	{Name: "domain", Transient: true, DataType: sky.Factor},
	{Name: "path", Transient: true, DataType: sky.Factor},
}

// Account represents a collection of Users and Events.
type Account struct {
	Tx     *Tx
	id     int
	APIKey string `json:"apiKey"`
}

// ID returns the account identifier.
func (a *Account) ID() int {
	return a.id
}

// SkyTableName returns the name of the table used by Sky.
func (a *Account) SkyTableName() string {
	assert(a.id > 0, "uninitialized account does not have a sky table")
	return fmt.Sprintf("skybox-%d", a.id)
}

// SkyTable returns a reference to the table used by Sky.
func (a *Account) SkyTable() *sky.Table {
	return &sky.Table{
		Client: &a.Tx.db.SkyClient,
		Name:   a.SkyTableName(),
	}
}

// Validate validates all fields of the account.
func (a *Account) Validate() error {
	return nil
}

func (a *Account) get() ([]byte, error) {
	value := a.Tx.Bucket("accounts").Get(itob(a.id))
	if value == nil {
		return nil, ErrAccountNotFound
	}
	return value, nil
}

// Load retrieves an account from the database.
func (a *Account) Load() error {
	value, err := a.get()
	if err != nil {
		return err
	}
	unmarshal(value, &a)
	return nil
}

// Save commits the Account to the database.
func (a *Account) Save() error {
	assert(a.id > 0, "uninitialized account cannot be saved")

	// Autogenerate an API key if one does not exist.
	if len(a.APIKey) == 0 {
		if err := a.GenerateAPIKey(); err != nil {
			return err
		}
	}

	return a.Tx.Bucket("accounts").Put(itob(a.id), marshal(a))
}

// Delete removes the account from the database.
func (a *Account) Delete() error {
	err := a.Tx.Bucket("accounts").Delete(itob(a.id))
	assert(err == nil, "account delete error: %s", err)

	// TODO: Remove all users.

	return nil
}

// GenerateAPIKey creates a new API key for an account.
func (a *Account) GenerateAPIKey() error {
	// Remove old API key from index.
	if a.APIKey != "" {
		removeFromUniqueIndex(a.Tx, "accounts.APIKey", []byte(a.APIKey))
	}

	// Generate new API key.
	apiKey, err := uuid.NewV4()
	if err != nil {
		return err
	}
	a.APIKey = apiKey.String()

	// Update index.
	insertIntoUniqueIndex(a.Tx, "accounts.APIKey", []byte(a.APIKey), a.ID())

	return nil
}

// CreateUser creates a new User for this account.
func (a *Account) CreateUser(u *User) error {
	assert(u.id == 0, "create user with a non-zero id: %d", u.ID)
	assert(a.id > 0, "create user on unsaved account: %d", a.ID)
	if err := u.Validate(); err != nil {
		return err
	}

	// Generate password hash.
	if err := u.GenerateHash(); err != nil {
		return err
	}

	u.Tx = a.Tx
	u.AccountID = a.id

	// Verify account exists.
	if _, err := a.get(); err != nil {
		return err
	}

	// Verify that email is not taken.
	if id := getUniqueIndex(a.Tx, "user.email", []byte(u.Email)); id != 0 {
		return ErrUserEmailTaken
	}

	// Generate new id.
	u.id, _ = a.Tx.Bucket("users").NextSequence()
	assert(u.id > 0, "user sequence error")

	// Add user id to secondary index.
	insertIntoForeignKeyIndex(a.Tx, "account.users", itob(a.id), u.id)
	insertIntoUniqueIndex(a.Tx, "user.email", []byte(u.Email), u.id)

	// Save user.
	return u.Save()
}

// Users retrieves a list of all users for the account.
func (a *Account) Users() (Users, error) {
	users := make(Users, 0)
	index := getForeignKeyIndex(a.Tx, "account.users", itob(a.id))

	for _, id := range index {
		u := &User{Tx: a.Tx, id: id}
		err := u.Load()
		assert(err == nil, "user (%d) not found from account.users index (%d)", u.id, a.id)
		assert(u.AccountID == a.id, "user/account mismatch: %d (%d) not in %d", u.id, u.AccountID, a.id)
		users = append(users, u)
	}
	sort.Sort(users)

	return users, nil
}

// Funnel retrieves a funnel with a given ID.
// Only funnels associated with this account will be returned.
func (a *Account) Funnel(id int) (*Funnel, error) {
	assert(a.id > 0, "find funnel on unsaved account: %d", a.id)
	f, err := a.Tx.Funnel(id)
	if err != nil {
		return nil, err
	} else if f.AccountID != a.ID() {
		return nil, ErrFunnelNotFound
	}
	return f, nil
}

// CreateFunnel creates a new Funnel for this account.
func (a *Account) CreateFunnel(f *Funnel) error {
	assert(f.id == 0, "create funnel with a non-zero id: %d", f.id)
	assert(a.id > 0, "create funnel on unsaved account: %d", a.id)
	if err := f.Validate(); err != nil {
		return err
	}

	// Verify account exists.
	if _, err := a.get(); err != nil {
		return err
	}

	f.Tx = a.Tx
	f.AccountID = a.id

	// Generate new id.
	f.id, _ = a.Tx.Bucket("funnels").NextSequence()
	assert(a.id > 0, "funnel sequence error")

	// Add funnel id to secondary index.
	insertIntoForeignKeyIndex(a.Tx, "account.funnels", itob(a.id), f.id)

	// Save funnel.
	return f.Save()
}

// Funnels retrieves a list of all funnels for the account.
func (a *Account) Funnels() (Funnels, error) {
	funnels := make(Funnels, 0)
	index := getForeignKeyIndex(a.Tx, "account.funnels", itob(a.id))

	for _, id := range index {
		f := &Funnel{Tx: a.Tx, id: id}
		err := f.Load()
		assert(err == nil, "funnel (%d) not found from account.funnels index (%d)", f.id, a.id)
		assert(f.AccountID == a.id, "funnel/account mismatch: %d (%d) not in %d", f.id, f.AccountID, a.id)
		funnels = append(funnels, f)
	}
	sort.Sort(funnels)
	return funnels, nil
}

// Track sends a single event to Sky.
func (a *Account) Track(e *Event) error {
	t := a.SkyTable()

	// Serialize event and insert event into Sky.
	id := e.ID()
	skyEvent := e.Serialize()
	if err := t.InsertEvent(id, skyEvent); err != nil {
		return err
	}

	// TODO(benbjohnson): Merge timelines if necessary.
	// t.Merge(id, t.DeviceID)

	return nil
}

// Events returns a list of events for a given ID.
func (a *Account) Events(id string) ([]*Event, error) {
	t := a.SkyTable()
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

// Resources returns a list unique resources on the account.
func (a *Account) Resources() ([]string, error) {
	t := a.SkyTable()
	results, err := t.Query("SELECT count() GROUP BY resource")
	if err != nil {
		return nil, err
	}

	resources := make([]string, 0, len(results))
	if results, ok := results["resource"].(map[string]interface{}); ok {
		for resource, _ := range results {
			resources = append(resources, resource)
		}
	}
	sort.Sort(sort.StringSlice(resources))
	return resources, nil
}

// Migrate creates a Sky table and updates the schema, if necessary.
func (a *Account) Migrate() error {
	name := a.SkyTableName()
	if err := a.createSkyTableIfNotExists(); err != nil {
		return fmt.Errorf("migrate table error: %s: %s", name, err)
	}
	for _, property := range schema {
		if err := a.createSkyPropertyIfNotExists(property); err != nil {
			return fmt.Errorf("migrate property error: %s: %s: %s", name, property.Name, err)
		}
	}
	return nil
}

// Reset drops the sky table and recreates it.
func (a *Account) Reset() error {
	c := &a.Tx.db.SkyClient
	c.DeleteTable(a.SkyTableName())
	return a.Migrate()
}

func (a *Account) createSkyTableIfNotExists() error {
	c := &a.Tx.db.SkyClient
	if t, err := c.Table(a.SkyTableName()); t != nil && err == nil {
		return err
	}
	return c.CreateTable(a.SkyTable())
}

func (a *Account) createSkyPropertyIfNotExists(property *sky.Property) error {
	t := a.SkyTable()
	if tmp, err := t.Property(property.Name); tmp != nil && err == nil {
		return err
	}
	return t.CreateProperty(property)
}

type Accounts []*Account

func (s Accounts) Len() int           { return len(s) }
func (s Accounts) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Accounts) Less(i, j int) bool { return s[i].id < s[j].id }
