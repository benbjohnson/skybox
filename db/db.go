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

		err = txn.CreateBucketIfNotExists("account.users")
		assert(err == nil, "account.users bucket error: %s", err)

		err = txn.CreateBucketIfNotExists("users")
		assert(err == nil, "users bucket error: %s", err)

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
	assert(a.id == 0, "create account with a non-zero id: %d", a.Id)
	if err := a.Validate(); err != nil {
		return err
	}

	a.db = db
	return db.Do(func(txn *bolt.RWTransaction) error {
		var err error
		a.id, err = txn.NextSequence("accounts")
		assert(a.id > 0, "account sequence error: %s", err)
		return a.save(txn)
	})
}

// User retrieves a User from the database with the given identifier.
func (db *DB) User(id int) (*User, error) {
	u := &User{db: db, id: id}
	if err := u.Load(); err != nil {
		return nil, err
	}
	return u, nil
}

// getIndex retrieves a list of ids from a named index.
func getIndex(txn *bolt.Transaction, name string, key []byte) ids {
	// Retrieve index.
	v, err := txn.Get(name, key)
	assert(err == nil, "index error: %s", err)

	// Unmarshal the index.
	var index ids
	if v != nil && len(v) > 0 {
		unmarshal(v, &index)
	}

	return index
}

// insertIntoIndex adds an id into a named index within a transaction.
func insertIntoIndex(txn *bolt.RWTransaction, name string, key []byte, id int) {
	index := getIndex(&txn.Transaction, name, key)
	index = index.insert(id)
	err := txn.Put(name, key, marshal(index))
	assert(err == nil, "index insert error: %s", err)
}

// removeFromIndex removes an id from a named index within a transaction.
func removeFromIndex(txn *bolt.RWTransaction, name string, key []byte, id int) {
	index := getIndex(&txn.Transaction, name, key)
	index = index.remove(id)
	err := txn.Put(name, key, marshal(index))
	assert(err == nil, "index remove error: %s", err)
}
