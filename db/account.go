package db

import (
	"github.com/boltdb/bolt"
	"sort"
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

func (a *Account) get(txn *bolt.Transaction) ([]byte, error) {
	value, err := txn.Get("accounts", itob(a.id))
	assert(err == nil, "get account error: %s", a.id)
	if value == nil {
		return nil, ErrAccountNotFound
	}
	return value, nil
}

// Load retrieves an account from the database.
func (a *Account) Load() error {
	return a.db.With(func(txn *bolt.Transaction) error {
		return a.load(txn)
	})
}

func (a *Account) load(txn *bolt.Transaction) error {
	value, err := a.get(txn)
	if err != nil {
		return err
	}
	unmarshal(value, &a)
	return nil
}

// Save commits the Account to the database.
func (a *Account) Save() error {
	return a.db.Do(func(txn *bolt.RWTransaction) error {
		return a.save(txn)
	})
}

func (a *Account) save(txn *bolt.RWTransaction) error {
	assert(a.id > 0, "uninitialized account cannot be saved")
	return txn.Put("accounts", itob(a.id), marshal(a))
}

// Delete removes the account from the database.
func (a *Account) Delete() error {
	return a.db.Do(func(txn *bolt.RWTransaction) error {
		return a.del(txn)
	})
}

func (a *Account) del(txn *bolt.RWTransaction) error {
	err := txn.Delete("accounts", itob(a.id))
	assert(err == nil, "account delete error: %s", err)

	// TODO: Remove all users.
	// TODO: Remove all projects.

	return nil
}

// CreateUser creates a new User for this account.
func (a *Account) CreateUser(u *User) error {
	assert(u.id == 0, "create user with a non-zero id: %d", u.Id)
	assert(a.id > 0, "create user on unsaved account: %d", a.Id)
	if err := u.Validate(); err != nil {
		return err
	}

	// Generate password hash.
	if err := u.GenerateHash(); err != nil {
		return err
	}

	u.db = a.db
	u.AccountId = a.id

	return u.db.Do(func(txn *bolt.RWTransaction) error {
		// Verify account exists.
		if _, err := a.get(&txn.Transaction); err != nil {
			return err
		}

		// Verify that username is not taken.
		if id := getUniqueIndex(&txn.Transaction, "user.username", []byte(u.Username)); id != 0 {
			return ErrUserUsernameTaken
		}

		// Generate new id.
		u.id, _ = txn.NextSequence("users")
		assert(u.id > 0, "user sequence error")

		// Add user id to secondary index.
		insertIntoForeignKeyIndex(txn, "account.users", itob(a.id), u.id)
		insertIntoUniqueIndex(txn, "user.username", []byte(u.Username), u.id)

		// Save user.
		return u.save(txn)
	})
}

// Users retrieves a list of all users for the account.
func (a *Account) Users() (Users, error) {
	users := make(Users, 0)
	err := a.db.With(func(txn *bolt.Transaction) error {
		index := getForeignKeyIndex(txn, "account.users", itob(a.id))

		for _, id := range index {
			u := &User{db: a.db, id: id}
			err := u.load(txn)
			assert(err == nil, "user (%d) not found from account.users index (%d)", u.id, a.id)
			assert(u.AccountId == a.id, "user/account mismatch: %d (%d) not in %d", u.id, u.AccountId, a.id)
			users = append(users, u)
		}
		return nil
	})
	assert(err == nil, "users retrieval error: %s", err)
	sort.Sort(users)
	return users, nil
}

// CreateProject creates a new Project for this account.
func (a *Account) CreateProject(p *Project) error {
	assert(p.id == 0, "create project with a non-zero id: %d", p.id)
	assert(a.id > 0, "create project on unsaved account: %d", a.id)
	if err := p.Validate(); err != nil {
		return err
	}

	p.db = a.db
	p.AccountId = a.id

	return p.db.Do(func(txn *bolt.RWTransaction) error {
		// Verify account exists.
		if _, err := a.get(&txn.Transaction); err != nil {
			return err
		}

		// Generate new id.
		p.id, _ = txn.NextSequence("projects")
		assert(p.id > 0, "project sequence error")

		// Add project id to secondary index.
		insertIntoForeignKeyIndex(txn, "account.projects", itob(a.id), p.id)

		// Save project.
		return p.save(txn)
	})
}

// Projects retrieves a list of all projects for the account.
func (a *Account) Projects() (Projects, error) {
	projects := make(Projects, 0)
	err := a.db.With(func(txn *bolt.Transaction) error {
		index := getForeignKeyIndex(txn, "account.projects", itob(a.id))

		for _, id := range index {
			p := &Project{db: a.db, id: id}
			err := p.load(txn)
			assert(err == nil, "project (%d) not found from account.projects index (%d)", p.id, a.id)
			assert(p.AccountId == a.id, "project/account mismatch: %d (%d) not in %d", p.id, p.AccountId, a.id)
			projects = append(projects, p)
		}
		return nil
	})
	assert(err == nil, "projects retrieval error: %s", err)
	sort.Sort(projects)
	return projects, nil
}

type Accounts []*Account

func (s Accounts) Len() int           { return len(s) }
func (s Accounts) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Accounts) Less(i, j int) bool { return s[i].Name < s[j].Name }
