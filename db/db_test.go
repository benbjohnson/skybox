package db_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/skybox/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that a DB can generate and save a secret key for cookies.
func TestDBSecret(t *testing.T) {
	withDB(func(db *DB) {
		// Make sure the key is generated.
		b, err := db.Secret()
		assert.Equal(t, len(b), 64)
		assert.Nil(t, err)

		// Make sure we get the same key next time.
		b2, _ := db.Secret()
		assert.True(t, bytes.Equal(b, b2))

		// Make sure the secret persists across saves.
		path := db.Path()
		db.Close()
		db.Open(path, 0666)
		b3, _ := db.Secret()
		assert.True(t, bytes.Equal(b, b3))
	})
}

// withDB executes a function with an open database.
func withDB(fn func(*DB)) {
	f, _ := ioutil.TempFile("", "skybox-")
	path := f.Name()
	f.Close()
	os.Remove(path)
	defer os.RemoveAll(path)

	var db DB
	if err := db.Open(path, 0666); err != nil {
		panic("db open error: " + err.Error())
	}
	defer db.Close()
	fn(&db)
}
