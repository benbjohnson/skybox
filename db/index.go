package db

// getForeignKeyIndex retrieves a list of ids from a foreign index.
func getForeignKeyIndex(t *Tx, name string, key []byte) ids {
	// Retrieve index.
	v := t.Bucket(name).Get(key)

	// Unmarshal the index.
	var index ids
	if v != nil && len(v) > 0 {
		unmarshal(v, &index)
	}

	return index
}

// insertIntoForeignKeyIndex adds an id into a foreign key index within a transaction.
func insertIntoForeignKeyIndex(t *Tx, name string, key []byte, id int) {
	index := getForeignKeyIndex(t, name, key)
	index = index.insert(id)
	err := t.Bucket(name).Put(key, marshal(index))
	assert(err == nil, "foreign key index insert error: %s", err)
}

// removeFromForeignKeyIndex removes an id from a foreign key index within a transaction.
func removeFromForeignKeyIndex(t *Tx, name string, key []byte, id int) {
	index := getForeignKeyIndex(t, name, key)
	index = index.remove(id)
	err := t.Bucket(name).Put(key, marshal(index))
	assert(err == nil, "foreign key index remove error: %s", err)
}

// getUniqueIndex retrieves the id associated with a given value.
func getUniqueIndex(t *Tx, name string, key []byte) int {
	// Unmarshal the id.
	v := t.Bucket(name).Get(key)
	if v != nil && len(v) > 0 {
		return btoi(v)
	}
	return 0
}

// insertIntoUniqueIndex associates a value with an id.
// Panics if association already exists.
func insertIntoUniqueIndex(t *Tx, name string, key []byte, id int) {
	currentID := getUniqueIndex(t, name, key)
	assert(currentID == 0, "unique index overwrite: %d -> %d", currentID, id)
	err := t.Bucket(name).Put(key, itob(id))
	assert(err == nil, "unique index insert error: %s", err)
}

// removeFromUniqueIndex removes an association of a value to an id.
func removeFromUniqueIndex(t *Tx, name string, key []byte) {
	err := t.Bucket(name).Delete(key)
	assert(err == nil, "unique index remove error: %s", err)
}
