package db

import (
	"os"

	"github.com/boltdb/bolt"
	"github.com/gorilla/securecookie"
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
	err := db.Do(func(txn *Transaction) error {
		err := txn.CreateBucketIfNotExists("system")
		assert(err == nil, "system bucket error: %s", err)

		err = txn.CreateBucketIfNotExists("accounts")
		assert(err == nil, "accounts bucket error: %s", err)

		err = txn.CreateBucketIfNotExists("account.users")
		assert(err == nil, "account.users bucket error: %s", err)

		err = txn.CreateBucketIfNotExists("account.projects")
		assert(err == nil, "account.projects bucket error: %s", err)

		err = txn.CreateBucketIfNotExists("users")
		assert(err == nil, "users bucket error: %s", err)

		err = txn.CreateBucketIfNotExists("user.email")
		assert(err == nil, "user.email bucket error: %s", err)

		err = txn.CreateBucketIfNotExists("projects")
		assert(err == nil, "projects bucket error: %s", err)

		err = txn.CreateBucketIfNotExists("project.funnels")
		assert(err == nil, "project.funnels bucket error: %s", err)

		err = txn.CreateBucketIfNotExists("funnels")
		assert(err == nil, "funnels bucket error: %s", err)

		return nil
	})

	if err != nil {
		db.Close()
		return err
	}

	return nil
}

// Do executes a function within the context of a writable transaction.
func (db *DB) Do(fn func(*Transaction) error) error {
	return db.DB.Do(func(t *bolt.RWTransaction) error {
		return fn(&Transaction{&t.Transaction, t})
	})
}

// With executes a function within the context of a read-only transaction.
func (db *DB) With(fn func(*Transaction) error) error {
	return db.DB.With(func(t *bolt.Transaction) error {
		return fn(&Transaction{t, nil})
	})
}

// Secret retrieves the secret key used for cookie storage.
func (db *DB) Secret() ([]byte, error) {
	var secret []byte
	err := db.Do(func(t *Transaction) error {
		b := t.Bucket("system")
		secret = b.Get([]byte("secret"))

		if secret == nil {
			secret = securecookie.GenerateRandomKey(64)
			err := b.Put([]byte("secret"), secret)
			assert(err == nil, "secret gen error: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return secret, nil
}
