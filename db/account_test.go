package db_test

import (
	"testing"

	. "github.com/benbjohnson/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that an account can update itself.
func TestAccountUpdate(t *testing.T) {
	withDB(func(db *DB) {
		// Create and update account.
		a := &Account{Name: "Foo"}
		assert.NoError(t, db.CreateAccount(a))
		a.Name = "Bar"
		assert.NoError(t, a.Save())

		// Retrieve the account.
		a2, err := db.Account(1)
		if assert.NoError(t, err) && assert.NotNil(t, a2) {
			assert.Equal(t, a2.Name, "Bar")
		}
	})
}

// Ensure that an account can be deleted.
func TestAccountDelete(t *testing.T) {
	withDB(func(db *DB) {
		// Create account.
		assert.NoError(t, db.CreateAccount(&Account{Name: "Foo"}))

		// Retrieve and delete account.
		a, _ := db.Account(1)
		assert.NoError(t, a.Delete())

		// Retrieve the account again.
		_, err := db.Account(1)
		assert.Equal(t, err, ErrAccountNotFound)
	})
}

// Ensure that an account can retrieve all associated users.
func TestAccountUsers(t *testing.T) {
	withDB(func(db *DB) {
		// Create two accounts.
		a1 := &Account{Name: "foo"}
		assert.NoError(t, db.CreateAccount(a1))
		a2 := &Account{Name: "bar"}
		assert.NoError(t, db.CreateAccount(a2))

		// Add users to first account.
		assert.NoError(t, a1.CreateUser(&User{Username: "susyque", Password: "password"}))
		assert.NoError(t, a1.CreateUser(&User{Username: "johndoe", Password: "password"}))

		// Add users to second account.
		assert.NoError(t, a2.CreateUser(&User{Username: "billybob", Password: "password"}))

		// Check first account users.
		users, err := a1.Users()
		if assert.NoError(t, err) && assert.Equal(t, len(users), 2) {
			assert.Equal(t, users[0].DB(), db)
			assert.Equal(t, users[0].Id(), 2)
			assert.Equal(t, users[0].AccountId, 1)
			assert.Equal(t, users[0].Username, "johndoe")

			assert.Equal(t, users[1].DB(), db)
			assert.Equal(t, users[1].Id(), 1)
			assert.Equal(t, users[1].AccountId, 1)
			assert.Equal(t, users[1].Username, "susyque")
		}

		// Check second account users.
		users, err = a2.Users()
		if assert.NoError(t, err) && assert.Equal(t, len(users), 1) {
			assert.Equal(t, users[0].DB(), db)
			assert.Equal(t, users[0].Id(), 3)
			assert.Equal(t, users[0].AccountId, 2)
			assert.Equal(t, users[0].Username, "billybob")
		}
	})
}

// Ensure that an account can retrieve all associated projects.
func TestAccountProjects(t *testing.T) {
	withDB(func(db *DB) {
		// Create two accounts.
		a1 := &Account{Name: "foo"}
		assert.NoError(t, db.CreateAccount(a1))
		a2 := &Account{Name: "bar"}
		assert.NoError(t, db.CreateAccount(a2))

		// Add projects to first account.
		assert.NoError(t, a1.CreateProject(&Project{Name: "Project Y"}))
		assert.NoError(t, a1.CreateProject(&Project{Name: "Project X"}))

		// Add projects to second account.
		assert.NoError(t, a2.CreateProject(&Project{Name: "Project A"}))

		// Check first account projects.
		projects, err := a1.Projects()
		if assert.NoError(t, err) && assert.Equal(t, len(projects), 2) {
			assert.Equal(t, projects[0].DB(), db)
			assert.Equal(t, projects[0].Id(), 2)
			assert.Equal(t, projects[0].AccountId, 1)
			assert.Equal(t, projects[0].Name, "Project X")

			assert.Equal(t, projects[1].DB(), db)
			assert.Equal(t, projects[1].Id(), 1)
			assert.Equal(t, projects[1].AccountId, 1)
			assert.Equal(t, projects[1].Name, "Project Y")
		}

		// Check second account projects.
		projects, err = a2.Projects()
		if assert.NoError(t, err) && assert.Equal(t, len(projects), 1) {
			assert.Equal(t, projects[0].DB(), db)
			assert.Equal(t, projects[0].Id(), 3)
			assert.Equal(t, projects[0].AccountId, 2)
			assert.Equal(t, projects[0].Name, "Project A")
		}
	})
}
