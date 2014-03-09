package db

import (
	"sort"

	"github.com/boltdb/bolt"
)

type Tx struct {
	*bolt.Tx
	db *DB
}

// Account retrieves an Account from the database with the given identifier.
func (t *Tx) Account(id int) (*Account, error) {
	a := &Account{Tx: t, id: id}
	if err := a.Load(); err != nil {
		return nil, err
	}
	return a, nil
}

// Accounts retrieves all Account objects from the database.
func (t *Tx) Accounts() (Accounts, error) {
	accounts := make(Accounts, 0)
	err := t.Bucket("accounts").ForEach(func(k, v []byte) error {
		a := &Account{Tx: t, id: btoi(k)}
		unmarshal(v, &a)
		accounts = append(accounts, a)
		return nil
	})
	assert(err == nil, "accounts retrieval error: %s", err)
	sort.Sort(accounts)
	return accounts, nil
}

// User retrieves a User from the database with the given identifier.
func (t *Tx) User(id int) (*User, error) {
	u := &User{Tx: t, id: id}
	if err := u.Load(); err != nil {
		return nil, err
	}
	return u, nil
}

// UserByEmail retrieves a User from the database with the given Email.
func (t *Tx) UserByEmail(email string) (*User, error) {
	u := &User{Tx: t}
	if u.id = getUniqueIndex(t, "user.email", []byte(email)); u.id == 0 {
		return nil, ErrUserNotFound
	}
	if err := u.Load(); err != nil {
		return nil, err
	}
	return u, nil
}

// CreateAccount creates a new Account in the database.
func (t *Tx) CreateAccount(a *Account) error {
	assert(a.id == 0, "create account with a non-zero id: %d", a.ID)
	if err := a.Validate(); err != nil {
		return err
	}
	a.Tx = t

	var err error
	a.id, err = t.Bucket("accounts").NextSequence()
	assert(a.id > 0, "account sequence error: %s", err)
	if err := a.Save(); err != nil {
		return err
	}

	// Create Sky table.
	if err := a.Migrate(); err != nil {
		return err
	}

	return nil
}

// AccountByAPIKey retrieves an acocunt from the database with the given API key.
func (t *Tx) AccountByAPIKey(apiKey string) (*Account, error) {
	a := &Account{Tx: t}
	if a.id = getUniqueIndex(t, "accounts.APIKey", []byte(apiKey)); a.id == 0 {
		return nil, ErrAccountNotFound
	}
	if err := a.Load(); err != nil {
		return nil, err
	}
	return a, nil
}

// Funnel retrieves a Funnel from the database with the given identifier.
func (t *Tx) Funnel(id int) (*Funnel, error) {
	f := &Funnel{Tx: t, id: id}
	if err := f.Load(); err != nil {
		return nil, err
	}
	return f, nil
}
