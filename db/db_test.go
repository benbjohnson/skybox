package db_test

import (
	"io/ioutil"
	"os"

	. "github.com/benbjohnson/skybox/db"
)

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
