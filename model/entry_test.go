package model

import "testing"

var testPeople = map[int]Entry{
	0: Entry{0, "Jessica", "Bellon", "214-555-9999", "Jess@Gmail.com"},
	1: Entry{0, "Will", "McGinnis", "991-909-0123", "btown@email.com"},
	2: Entry{0, "Tyler", "Higgins", "111-111-1111", "thiggs@mail.com"},
	3: Entry{0, "Zelda", "Bellon", "232-909-9998", "Zelda@dogs.com"},
}

// mustNewEntry basic testing block that must pass
func mustNewEntry(t *testing.T, fname string, lname string, phone string, email string) *Entry {
	entry, err := NewEntry(fname, lname, phone, email)
	if err != nil {
		t.Fatalf("Error: %v", err) // stops execution of test suite
	}
	return entry
}

func TestEquals(t *testing.T) {
	person0 := mustNewEntry(t, testPeople[0].FirstName, testPeople[0].LastName, testPeople[0].Phone, testPeople[0].Email)
	personZero := mustNewEntry(t, testPeople[0].FirstName, testPeople[0].LastName, testPeople[0].Phone, testPeople[0].Email)
	person1 := mustNewEntry(t, testPeople[1].FirstName, testPeople[1].LastName, testPeople[1].Phone, testPeople[1].Email)

	if !person0.Equals(personZero) || person0.Equals(person1) {
		t.Error("error in Equal function")
	}
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

func TestCopy(t *testing.T) {
	entry := testPeople[0]
	copy := copyEntry(&entry)

	if !entry.Equals(copy) {
		t.Errorf("Copy did not work")
	}
}
