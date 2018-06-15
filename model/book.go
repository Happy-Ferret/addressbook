package model

import (
	"encoding/csv"
	"errors"
	"os"
)

// Book is an addressbook struct
type Book struct {
	entries []*Entry
	lastID  uint64 // lastID in use
}

// NewBook returns an empty book
func NewBook() *Book {
	return &Book{}
}

// Save saves the entry to the addressbook
func (b *Book) Save(entry *Entry) {
	if !b.Valid(entry) {
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

// DeleteEntry deletes an entry from the addressbook
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

// Get returns a copy of the the (single) entry with the given id and a bool representing found
func (b *Book) Get(ID uint64) (Entry, bool) {
	for _, e := range b.entries {
		if e.ID == ID {
			return *e, true
		}
	}
	return Entry{}, false
}

// Search returns a slice of entry copies that have (first/last) name matching search parameters
func (b *Book) Search(name string) []Entry {
	var matches []Entry
	for _, e := range b.entries {
		if e.FirstName == name || e.LastName == name {
			matches = append(matches, *e)
		}
	}
	return matches
}

// GetFirst returns a copy of the first entry that satisfies the search terms
func (b *Book) GetFirst(name string) Entry {
	return b.Search(name)[0]
}

// Valid returns true if entry is  not identical (same first, last, phone, and email)
// to an existing entry in the book
func (b *Book) Valid(entry *Entry) bool {
	for _, e := range b.entries {
		if e.FirstName == entry.FirstName && e.LastName == entry.LastName &&
			e.Phone == entry.Phone && e.Email == entry.Email {
			return false // already exists in book
		}
	}
	return true // does not already exist (safe to save)
}

func (b *Book) Export() error {
	file, err := os.Create("./addressbook.csv")
	if err != nil {
		return errors.New("Error creating CSV file")
	}
	defer file.Close()

	w := csv.NewWriter(file)

	for _, p := range b.entries {
		record := []string{p.FirstName, p.LastName, p.Phone, p.Email}
		w.Write(record)
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
