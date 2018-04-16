package person

import (
	_ "github.com/JessTheBell/addressbook/book"
	"github.com/gogo/protobuf/proto"
)

// MarshalPerson encodes person to binary.
func MarshalPerson(p *book.Person) ([]byte, error) {
	return proto.Marshal(&Person{
		ID:        int32(p.ID),
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Phone:     p.Phone,
	})
}

// UnmarshalPerson decodes a dial from binary.
func UnmarshalPerson(data []byte, p *book.Person) error {

	var pb Person
	if err := proto.Unmarshal(data, &pb); err != nil {
		return err
	}

	p.ID = pb.ID
	p.FirstName = pb.FirstName
	p.LastName = pb.LastName
	p.Email = pb.Email
	p.Phone = pb.Phone

	return nil
}
