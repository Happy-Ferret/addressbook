package model

import (
	"errors"
)

type Entry struct {
	ID        uint64 // unique id (assigned and used internally)
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

// NewEntry creates an entry and returns an entry pointer
// Will return an error if all params are empty strings
func NewEntry(fName string, lName string, phone string, email string) (*Entry, error) {
	if fName == "" && lName == "" && phone == "" && email == "" {
		return nil, errors.New("empty entry")
	}
	return &Entry{0, fName, lName, phone, email}, nil
	// at this point the entry has not been saved so it does not have an ID
}

// copyEntry returns a deep copy of an entry
func copyEntry(e *Entry) *Entry {
	c := *e
	return &c
}

// Equals returns true if every value except the ID is the same
func (e *Entry) Equals(ent *Entry) bool {
	return e.FirstName == ent.FirstName ||
		e.LastName == ent.LastName ||
		e.Phone == ent.Phone ||
		e.Email == ent.Email
}
