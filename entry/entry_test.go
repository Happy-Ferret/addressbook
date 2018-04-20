package entry

import (
	"testing"
)

var testPeople = map[int]Entry{
	0: Entry{0, "Jessica", "Bellon", "214-555-9999", "Jess@Gmail.com"},
	1: Entry{0, "Will", "McGinnis", "991-909-0123", "btown@email.com"},
	2: Entry{0, "Tyler", "Higgins", "111-111-1111", "thiggs@mail.com"},
	3: Entry{0, "Zelda", "Bellon", "232-909-9998", "Zelda@dogs.com"},
}

func mustNewEntry(t *testing.T, fname string, lname string, phone string, email string) *Entry {
	entry, err := NewEntry(fname, lname, phone, email)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	return entry
}

func TestNewEntry(t *testing.T) {
	person := testPeople[1]
	entry := mustNewEntry(t, person.FirstName, person.LastName, person.Phone, person.Email)
	if entry.ID != 0 {
		t.Errorf("Expected ID %q, got %q", 0, entry.ID)
	}

	if entry.FirstName != person.FirstName {
		t.Errorf("Expected First Name %q, got %q", person.FirstName, entry.FirstName)
	}
	if entry.LastName != person.LastName {
		t.Errorf("Expected Last Name %q, got %q", person.LastName, entry.LastName)
	}

	if entry.Phone != person.Phone {
		t.Errorf("Expected Phone number %q, got %q", person.Phone, entry.Phone)
	}
	if entry.Email != person.Email {
		t.Errorf("Expected Email address %q, got %q", person.Email, entry.Email)
	}
}

func TestNewEntryFailing(t *testing.T) {
	_, err := NewEntry("", "", "", "")
	if err == nil {
		t.Errorf("expected 'empty entry' error")
	}

}

func TestSaveAndRetrieve(t *testing.T) {
	person := testPeople[1]
	entry := mustNewEntry(t, person.FirstName, person.LastName, person.Phone, person.Email)

	b := NewBook()
	b.Save(entry)

	all := b.All()
	if len(all) != 1 {
		t.Errorf("expected 1 entry, got %v", len(all))
	}
	if *all[0] != *entry {
		t.Errorf("expected %v, got %v", entry, all[0])
	}
}

func TestSaveAndRetrieveTwo(t *testing.T) {
	person := testPeople[1]
	entry1 := mustNewEntry(t, person.FirstName, person.LastName, person.Phone, person.Email)

	person = testPeople[2]
	entry2 := mustNewEntry(t, person.FirstName, person.LastName, person.Phone, person.Email)

	b := NewBook()
	b.Save(entry1)
	b.Save(entry2)

	all := b.All()

	if len(all) != 2 {
		t.Errorf("expected two entries, got %v", len(all))
	}
	if *all[0] != *entry1 && *all[1] != *entry1 {
		t.Errorf("missing entry: %v", entry1)
	}
	if *all[0] != *entry2 && *all[1] != *entry2 {
		t.Errorf("missing entry: %v", entry2)
	}

}

func TestSaveModifyAndRetrieve(t *testing.T) {
	person := testPeople[0]
	entry := mustNewEntry(t, person.FirstName, person.LastName, person.Phone, person.Email)

	b := NewBook()
	b.Save(entry)

	entry.FirstName = "Zelda"
	if b.All()[0].FirstName != person.FirstName {
		t.Errorf("Something Went Wrong. Saved entry's name was %v", person.FirstName)
	}
}

func TestSaveTwice(t *testing.T) {

	person := testPeople[2]
	entry := mustNewEntry(t, person.FirstName, person.LastName, person.Phone, person.Email)

	b := NewBook()
	b.Save(entry)
	b.Save(entry)

	all := b.All()
	if len(all) != 1 {
		t.Errorf("expected 1 entry, got %v", len(all))
	}
	if *all[0] != *entry {
		t.Errorf("expected entry %v, got %v", entry, all[0])

	}
}

func TestSaveAndGet(t *testing.T) {

	person := testPeople[1]
	entry := mustNewEntry(t, person.FirstName, person.LastName, person.Phone, person.Email)
	b := NewBook()
	b.Save(entry)
	ent, ok := b.Get(entry.ID)
	if !ok {
		t.Errorf("Entry not found")
	}
	if *entry != *ent {
		t.Errorf("expected %v, got %v", entry, ent)
	}
}

func TestFailingGet(t *testing.T) {
	b := NewBook()
	f1, ok := b.Get(1)
	if ok {
		t.Errorf("expected to not find, found %v", f1)
	}
}

func TestFind(t *testing.T) {
	b := NewBook()
	for i := range testPeople {
		person := testPeople[i]
		entry := mustNewEntry(t, person.FirstName, person.LastName, person.Phone, person.Email)
		b.Save(entry)
	}

	if len(b.All()) != len(testPeople) {
		t.Errorf("Expected %v entries, got %v", len(testPeople), len(b.All()))
	}

	matches := b.Search("Bellon")
	if len(matches) != 2 {
		t.Errorf("Expected 2 matches, got %v", len(matches))
	}
}

// 100 % coverage?
