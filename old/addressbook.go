package main

// Person represents an entry in the addressbook.
//
type Person struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// Client creates a connection to the services.
type Client interface {
	AddressBookService() AddressBookService
}

// AddressBookService represents a service for managing the addressbook.
type AddressBookService interface {
	GetEntry(id uint) (*Person, error)
	AddEntry(p *Person) error
	DeleteEntry(id uint) error
	ModifyEntry(id uint) error
	ListAll() error
	//import
	//export
}
