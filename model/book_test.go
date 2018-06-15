package model

import "testing"

// populateBook populates an addressbook with test data
func populateBook(t *testing.T) *Book {
	b := NewBook()
	for _, i := range testPeople {
		entry := mustNewEntry(t, i.FirstName, i.LastName, i.Phone, i.Email)
		b.Save(entry)
	}
	return b
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

	b := populateBook(t)

	// returns a single entry which has a name "zelda"
	testEntry := b.GetFirst("Zelda")

	entry, found := b.Get(testEntry.ID)
	if !found {
		t.Error("Saved entry not found")
	}

	entry.LastName = "TheDog"
	t.Log(entry.ID)
	b.Save(&entry)

	entry, found = b.Get(testEntry.ID)
	if !found {
		t.Error("Saved entry not found")
	}
	if entry.LastName != "TheDog" {
		t.Error("Entry did not modify as expected")
	}
}

func TestGetFirst(t *testing.T) {

	testParam := "Bellon"
	b := populateBook(t)
	t1 := b.GetFirst(testParam)
	t2 := b.Search(testParam)[0]
	if t1 != t2 {
		t.Errorf("Expected the same response got: %v and %v", t1, t2)
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
	if *entry != ent {
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

func TestDelete(t *testing.T) {
	b := populateBook(t)

	lenB := len(b.All())
	entry := b.GetFirst("Bellon")

	b.Delete(entry.ID)
	if len(b.All()) == lenB {
		t.Error("Entry did not get deleted")
	}

	if _, found := b.Get(entry.ID); found {
		t.Error("Entry did not get deleted")

	}
}
