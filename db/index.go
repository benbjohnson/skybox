package db

import (
	"github.com/boltdb/bolt"
)

// getForeignKeyIndex retrieves a list of ids from a foreign index.
func getForeignKeyIndex(txn *bolt.Transaction, name string, key []byte) ids {
	// Retrieve index.
	v, err := txn.Get(name, key)
	assert(err == nil, "foreign key index error: %s", err)

	// Unmarshal the index.
	var index ids
	if v != nil && len(v) > 0 {
		unmarshal(v, &index)
	}

	return index
}

// insertIntoForeignKeyIndex adds an id into a foreign key index within a transaction.
func insertIntoForeignKeyIndex(txn *bolt.RWTransaction, name string, key []byte, id int) {
	index := getForeignKeyIndex(&txn.Transaction, name, key)
	index = index.insert(id)
	err := txn.Put(name, key, marshal(index))
	assert(err == nil, "foreign key index insert error: %s", err)
}

// removeFromForeignKeyIndex removes an id from a foreign key index within a transaction.
func removeFromForeignKeyIndex(txn *bolt.RWTransaction, name string, key []byte, id int) {
	index := getForeignKeyIndex(&txn.Transaction, name, key)
	index = index.remove(id)
	err := txn.Put(name, key, marshal(index))
	assert(err == nil, "foreign key index remove error: %s", err)
}

// getUniqueIndex retrieves the id associated with a given value.
func getUniqueIndex(txn *bolt.Transaction, name string, key []byte) int {
	v, err := txn.Get(name, key)
	assert(err == nil, "index error: %s", err)

	// Unmarshal the id.
	if v != nil && len(v) > 0 {
		return btoi(v)
	}
	return 0
}

// insertIntoUniqueIndex associates a value with an id.
// Panics if association already exists.
func insertIntoUniqueIndex(txn *bolt.RWTransaction, name string, key []byte, id int) {
	currentId := getUniqueIndex(&txn.Transaction, name, key)
	assert(currentId == 0, "unique index overwrite: %d -> %d", currentId, id)
	err := txn.Put(name, key, itob(id))
	assert(err == nil, "unique index insert error: %s", err)
}

// removeFromUniqueIndex removes an association of a value to an id.
func removeFromUniqueIndex(txn *bolt.RWTransaction, name string, key []byte) {
	err := txn.Delete(name, key)
	assert(err == nil, "unique index remove error: %s", err)
}
