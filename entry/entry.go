package entry

import (
	"encoding/csv"
	"errors"
	"os"
)

type Entry struct {
	ID        uint64 // unique id (assigned and used internally)
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

// NewEntry creates an entry given first name, last name, phone number, and email address
// returns an entry pointer and error
// Will reject if all params are empty strings
func NewEntry(fName string, lName string, phone string, email string) (*Entry, error) {
	if fName == "" && lName == "" && phone == "" && email == "" {
		return nil, errors.New("empty entry")
	}
	return &Entry{0, fName, lName, phone, email}, nil
	// at this point the entry has not been saved to the address book, so it does not have an id
}

// copyEntry returns a deep copy of an entry
func copyEntry(e *Entry) *Entry {
	c := *e
	return &c
}

// Book is an addressbook struct
type Book struct {
	entries []*Entry
	lastID  uint64 // LastID in use
}

// NewBook returns an empty book
func NewBook() *Book {
	return &Book{}
}

// Save saves the entry to the addressbook
func (b *Book) Save(entry *Entry) {
	if !b.validateInput(entry) {
		return
	}
	if entry.ID == 0 { // new entry
		b.lastID++
		entry.ID = b.lastID
		b.entries = append(b.entries, copyEntry(entry))
	} else { // existing entry (update)
		for i, e := range b.entries {
			if entry.ID == e.ID {
				b.entries[i] = copyEntry(entry)
				return
			}
		}
	}
}

// DeleteEntry deletes an entry from the addressvook
func (b *Book) Delete(id uint64) {
	for i, e := range b.entries {
		if e.ID == id {
			b.entries[i] = b.entries[len(b.All())-1]
			b.entries[len(b.All())-1] = nil
			b.entries = b.entries[:len(b.All())-1]
			return
		}
	}
}

// All returns a slice of entry pointers that represents the entire contents of the addressbook
func (b *Book) All() []*Entry {
	return b.entries
}

// Get returns the a pointer to the (single) entry with the given id and a bool representing found
func (b *Book) Get(ID uint64) (*Entry, bool) {
	for _, e := range b.entries {
		if e.ID == ID {
			return e, true
		}
	}
	return nil, false
}

// Search returns a slice of entry pointers that have (first/last) name matching search parameters
func (b *Book) Search(name string) []*Entry {
	var matches []*Entry
	for _, e := range b.entries {
		if e.FirstName == name || e.LastName == name {
			matches = append(matches, e)
		}
	}
	return matches
}

// validateInput returns true if entry is  not identical (same first, last, phone, and email)
// to an existing entry in the book
func (b *Book) validateInput(entry *Entry) bool {
	for _, e := range b.entries {
		if e.FirstName == entry.FirstName && e.LastName == entry.LastName &&
			e.Phone == entry.Phone && e.Email == entry.Email {
			return false // already exists in book
		}
	}
	return true // does not already exist (safe to save)
}

// Export
func (b *Book) Export(f string) error {
	file, err := os.Create(f)
	if err != nil {
		return errors.New("Error creating CSV file")
	}
	defer file.Close()

	w := csv.NewWriter(file)

	for _, e := range b.entries {
		row := []string{e.FirstName, e.LastName, e.Phone, e.Email}
		w.Write(row)
		if err != nil {
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}
