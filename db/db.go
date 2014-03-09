package db

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/gorilla/securecookie"
	"github.com/skydb/gosky"
)

// DB represents a Bolt-backed data store.
// The DB stores all non-event data.
type DB struct {
	bolt.DB
	SkyClient sky.Client
}

// Open initializes and opens the database.
func (db *DB) Open(path string, mode os.FileMode) error {
	if err := db.DB.Open(path, mode); err != nil {
		return err
	}

	// Create buckets.
	err := db.Do(func(tx *Tx) error {
		err := tx.CreateBucketIfNotExists("system")
		assert(err == nil, "system bucket error: %s", err)

		err = tx.CreateBucketIfNotExists("accounts")
		assert(err == nil, "accounts bucket error: %s", err)

		err = tx.CreateBucketIfNotExists("account.users")
		assert(err == nil, "account.users bucket error: %s", err)

		err = tx.CreateBucketIfNotExists("accounts.APIKey")
		assert(err == nil, "accounts.APIKey bucket error: %s", err)

		err = tx.CreateBucketIfNotExists("account.funnels")
		assert(err == nil, "account.funnels bucket error: %s", err)

		err = tx.CreateBucketIfNotExists("users")
		assert(err == nil, "users bucket error: %s", err)

		err = tx.CreateBucketIfNotExists("user.email")
		assert(err == nil, "user.email bucket error: %s", err)

		err = tx.CreateBucketIfNotExists("funnels")
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
func (db *DB) Do(fn func(*Tx) error) error {
	return db.DB.Do(func(tx *bolt.Tx) error {
		return fn(&Tx{tx, db})
	})
}

// With executes a function within the context of a read-only transaction.
func (db *DB) With(fn func(*Tx) error) error {
	return db.DB.With(func(tx *bolt.Tx) error {
		return fn(&Tx{tx, db})
	})
}

// Secret retrieves the secret key used for cookie storage.
func (db *DB) Secret() ([]byte, error) {
	var secret []byte
	err := db.Do(func(tx *Tx) error {
		b := tx.Bucket("system")
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

// assert will panic with a given formatted message if the given condition is false.
func assert(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf("assert failed: "+msg, v...))
	}
}

func warn(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
}

func warnf(msg string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", v...)
}
