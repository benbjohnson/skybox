package db

import (
	"os"
	"sort"

	"github.com/boltdb/bolt"
)

// DB represents a Bolt-backed data store.
// The DB stores all non-event data.
type DB struct {
	bolt.DB
}

// Open initializes and opens the database.
func (db *DB) Open(path string, mode os.FileMode) error {
	if err := db.DB.Open(path, mode); err != nil {
		return err
	}

	// Create buckets.
	err := db.Do(func(txn *bolt.RWTransaction) error {
		err := txn.CreateBucketIfNotExists("accounts")
		assert(err == nil, "accounts bucket error: %s", err)
		return nil
	})

	if err != nil {
		db.Close()
		return err
	}

	return nil
}

// Account retrieves an Account from the database with the given identifier.
func (db *DB) Account(id int) (*Account, error) {
	a := &Account{db: db, id: id}
	if err := a.Load(); err != nil {
		return nil, err
	}
	return a, nil
}

// Accounts retrieves all Account objects from the database.
func (db *DB) Accounts() (Accounts, error) {
	accounts := make(Accounts, 0)
	err := db.ForEach("accounts", func(k, v []byte) error {
		a := &Account{db: db, id: btoi(k)}
		unmarshal(v, &a)
		accounts = append(accounts, a)
		return nil
	})
	assert(err == nil, "accounts retrieval error: %s", err)
	sort.Sort(accounts)
	return accounts, nil
}

// CreateAccount creates a new Account in the database.
func (db *DB) CreateAccount(a *Account) error {
	assert(a.id == 0, "account creation with a non-zero id: %d", a.Id)
	if err := a.Validate(); err != nil {
		return err
	}

	a.db = db
	return db.Do(func(txn *bolt.RWTransaction) error {
		var err error
		a.id, err = txn.NextSequence("accounts")
		assert(a.id > 0, "account sequence error: %s", err)
		return a.SaveTo(txn)
	})
}
