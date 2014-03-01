package db_test

import (
	"testing"

	. "github.com/benbjohnson/skybox/db"
	"github.com/stretchr/testify/assert"
)

// Ensure that an account can be deleted.
func TestAccountDelete(t *testing.T) {
	withDB(func(db *DB) {
		// Create account.
		err := db.Do(func(txn *Transaction) error {
			return txn.CreateAccount(&Account{})
		})
		assert.NoError(t, err)

		// Retrieve and delete account.
		err = db.Do(func(txn *Transaction) error {
			a, _ := txn.Account(1)
			return a.Delete()
		})
		assert.NoError(t, err)

		// Retrieve the account again.
		err = db.With(func(txn *Transaction) error {
			_, err := txn.Account(1)
			return err
		})
		assert.Equal(t, err, ErrAccountNotFound)
	})
}

// Ensure that an account can retrieve all associated users.
func TestAccountUsers(t *testing.T) {
	withDB(func(db *DB) {
		// Create two accounts.
		db.Do(func(txn *Transaction) error {
			a1 := &Account{}
			assert.NoError(t, txn.CreateAccount(a1))
			a2 := &Account{}
			assert.NoError(t, txn.CreateAccount(a2))
			return nil
		})

		// Add users to first account.
		db.Do(func(txn *Transaction) error {
			a1, _ := txn.Account(1)
			assert.NoError(t, a1.CreateUser(&User{Email: "susyque@gmail.com", Password: "password"}))
			assert.NoError(t, a1.CreateUser(&User{Email: "johndoe@gmail.com", Password: "password"}))
			return nil
		})

		// Add users to second account.
		db.Do(func(txn *Transaction) error {
			a2, _ := txn.Account(2)
			assert.NoError(t, a2.CreateUser(&User{Email: "billybob@gmail.com", Password: "password"}))
			return nil
		})

		// Check first account users.
		db.With(func(txn *Transaction) error {
			a1, _ := txn.Account(1)
			users, err := a1.Users()
			if assert.NoError(t, err) && assert.Equal(t, len(users), 2) {
				assert.Equal(t, users[0].Transaction, txn)
				assert.Equal(t, users[0].ID(), 2)
				assert.Equal(t, users[0].AccountID, 1)
				assert.Equal(t, users[0].Email, "johndoe@gmail.com")

				assert.Equal(t, users[1].Transaction, txn)
				assert.Equal(t, users[1].ID(), 1)
				assert.Equal(t, users[1].AccountID, 1)
				assert.Equal(t, users[1].Email, "susyque@gmail.com")
			}
			return nil
		})

		// Check second account users.
		db.With(func(txn *Transaction) error {
			a2, _ := txn.Account(2)
			users, err := a2.Users()
			if assert.NoError(t, err) && assert.Equal(t, len(users), 1) {
				assert.Equal(t, users[0].Transaction, txn)
				assert.Equal(t, users[0].ID(), 3)
				assert.Equal(t, users[0].AccountID, 2)
				assert.Equal(t, users[0].Email, "billybob@gmail.com")
			}
			return nil
		})
	})
}

// Ensure that an account can retrieve all associated projects.
func TestAccountProjects(t *testing.T) {
	withDB(func(db *DB) {
		// Create two accounts.
		db.Do(func(txn *Transaction) error {
			a1 := &Account{}
			assert.NoError(t, txn.CreateAccount(a1))
			a2 := &Account{}
			assert.NoError(t, txn.CreateAccount(a2))

			// Add projects to first account.
			assert.NoError(t, a1.CreateProject(&Project{Name: "Project Y"}))
			assert.NoError(t, a1.CreateProject(&Project{Name: "Project X"}))

			// Add projects to second account.
			assert.NoError(t, a2.CreateProject(&Project{Name: "Project A"}))
			return nil
		})

		// Check first account projects.
		db.With(func(txn *Transaction) error {
			a1, _ := txn.Account(1)
			projects, err := a1.Projects()
			if assert.NoError(t, err) && assert.Equal(t, len(projects), 2) {
				assert.Equal(t, projects[0].Transaction, txn)
				assert.Equal(t, projects[0].ID(), 2)
				assert.Equal(t, projects[0].AccountID, 1)
				assert.Equal(t, projects[0].Name, "Project X")

				assert.Equal(t, projects[1].Transaction, txn)
				assert.Equal(t, projects[1].ID(), 1)
				assert.Equal(t, projects[1].AccountID, 1)
				assert.Equal(t, projects[1].Name, "Project Y")
			}
			return nil
		})

		// Check second account projects.
		db.With(func(txn *Transaction) error {
			a2, _ := txn.Account(2)
			projects, err := a2.Projects()
			if assert.NoError(t, err) && assert.Equal(t, len(projects), 1) {
				assert.Equal(t, projects[0].Transaction, txn)
				assert.Equal(t, projects[0].ID(), 3)
				assert.Equal(t, projects[0].AccountID, 2)
				assert.Equal(t, projects[0].Name, "Project A")
			}
			return nil
		})
	})
}
