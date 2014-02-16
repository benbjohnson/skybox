package db

import (
// "github.com/boltdb/bolt"
)

// Project represents a collection of Persons and their events.
type Project struct {
	db   *DB
	Name string
}

// Projects is a list of Project objects.
type Projects []*Project

// Validate validates all fields of the project.
func (p *Project) Validate() error {
	return nil // TODO
}

// Update updates all fields in the project using a map.
func (p *Project) Update(values map[string]interface{}) error {
	return nil // TODO
}

// Delete removes the project from the database.
func (p *Project) Delete() error {
	// TODO: Remove Sky table.
	return nil // TODO
}
