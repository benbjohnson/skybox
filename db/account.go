package db

import (
	"sort"
)

var (
	// ErrAccountNotFound is returned when an account with the given id does
	// not exist.
	ErrAccountNotFound = &Error{"account not found", nil}
)

// Account represents a collection of Users and Projects.
type Account struct {
	Tx *Tx
	id int
}

// ID returns the account identifier.
func (a *Account) ID() int {
	return a.id
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
	return a.Tx.Bucket("accounts").Put(itob(a.id), marshal(a))
}

// Delete removes the account from the database.
func (a *Account) Delete() error {
	err := a.Tx.Bucket("accounts").Delete(itob(a.id))
	assert(err == nil, "account delete error: %s", err)

	// TODO: Remove all users.
	// TODO: Remove all projects.

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

// Project retrieves a project with a given ID.
// Only projects associated with this account will be returned.
func (a *Account) Project(id int) (*Project, error) {
	assert(a.id > 0, "find project on unsaved account: %d", a.id)
	p, err := a.Tx.Project(id)
	if err != nil {
		return nil, err
	} else if p.AccountID != a.ID() {
		return nil, ErrProjectNotFound
	}
	return p, nil
}

// CreateProject creates a new Project for this account.
func (a *Account) CreateProject(p *Project) error {
	assert(p.id == 0, "create project with a non-zero id: %d", p.id)
	assert(a.id > 0, "create project on unsaved account: %d", a.id)
	if err := p.Validate(); err != nil {
		return err
	}

	p.Tx = a.Tx
	p.AccountID = a.id

	// Verify account exists.
	if _, err := a.get(); err != nil {
		return err
	}

	// Generate new id.
	p.id, _ = a.Tx.Bucket("projects").NextSequence()
	assert(p.id > 0, "project sequence error")

	// Add project id to secondary index.
	insertIntoForeignKeyIndex(a.Tx, "account.projects", itob(a.id), p.id)

	// Save project.
	if err := p.Save(); err != nil {
		return err
	}

	// Create Sky table.
	if err := p.Migrate(); err != nil {
		return err
	}

	return nil
}

// Projects retrieves a list of all projects for the account.
func (a *Account) Projects() (Projects, error) {
	projects := make(Projects, 0)
	index := getForeignKeyIndex(a.Tx, "account.projects", itob(a.id))

	for _, id := range index {
		p := &Project{Tx: a.Tx, id: id}
		err := p.Load()
		assert(err == nil, "project (%d) not found from account.projects index (%d)", p.id, a.id)
		assert(p.AccountID == a.id, "project/account mismatch: %d (%d) not in %d", p.id, p.AccountID, a.id)
		projects = append(projects, p)
	}
	sort.Sort(projects)
	return projects, nil
}

type Accounts []*Account

func (s Accounts) Len() int           { return len(s) }
func (s Accounts) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Accounts) Less(i, j int) bool { return s[i].id < s[j].id }
