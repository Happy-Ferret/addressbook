package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"

	pb "github.com/JessTheBell/addressbook/pb"
	"github.com/golang/protobuf/proto"
)

func main() {
	list()
	add()
	list()

}

type length int64

const (
	sizeOfLength = 8
	dbFile       = "./addressbook.pb"
)

var endianness = binary.LittleEndian

// add prompts user for a new entry into the address book
// for now just have a default entry for testing
func Add() error {
	entry := &pb.Person{
		FirstName: "Ted",
		LastName:  "Tedson",
		Email:     "tted@gmail.com",
		Phone:     "312-222-5564",
	}

	b, err := proto.Marshal(entry)
	if err != nil {
		return fmt.Errorf("could not encode: %v", err)
	}

	file, err := os.OpenFile(dbFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return fmt.Errorf("Could not open %s: %v", dbFile, err)
	}
	if err := binary.Write(file, endianness, length(len(b))); err != nil {
		return fmt.Errorf("Could not encode length of message: %v", err)
	}
	_, err = file.Write(b)
	if err != nil {
		return fmt.Errorf("Could not write to file: %v", err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("Could not close file %s: %v", dbFile, err)
	}
	return nil
}

// list
func list() error {

	// does db file exist?
	// if no then just return
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		fmt.Println("*There are no entries in your addressbook*")
		return nil
	}

	b, err := ioutil.ReadFile(dbFile)
	if err != nil {
		return fmt.Errorf("Could not read %s: %v", dbFile, err)
	}
	// infinite loop (exited when dbfile is completely read)
	for {
		if len(b) == 0 {
			return nil
		} else if len(b) < sizeOfLength {
			return fmt.Errorf("extra %d bytes remaining, this isnt good", len(b))
		}

		var l length
		if err := binary.Read(bytes.NewReader(b[:sizeOfLength]), endianness, &l); err != nil {
			return fmt.Errorf("could not decode message length: %v", err)
		}
		b = b[sizeOfLength:]

		var entry pb.Person
		if err := proto.Unmarshal(b[:l], &entry); err != nil {
			return fmt.Errorf("Could not read entry: %v", err)
		}
		b = b[l:]

		fmt.Printf("Name: %s %s \nEmail: %s \nPhone Number: %s\n",
			entry.FirstName, entry.LastName, entry.Email, entry.Phone)
	}
}
func find(string) error {

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return nil
	}
	b, err := ioutil.ReadFile(dbFile)
	if err != nil {
		return fmt.Errorf("Could not read %s: %v", dbFile, err)
	}
	//matches := []pb.Person{}
	for {
		if len(b) == 0 {
			return nil
		} else if len(b) < sizeOfLength {
			return fmt.Errorf("extra %d bytes remaining, this isnt good", len(b))
		}

		var l length
		if err := binary.Read(bytes.NewReader(b[:sizeOfLength]), endianness, &l); err != nil {
			return fmt.Errorf("could not decode message length: %v", err)
		}
		b = b[sizeOfLength:]

		var entry pb.Person
		if err := proto.Unmarshal(b[:l], &entry); err != nil {
			return fmt.Errorf("Could not read entry: %v", err)
		}
		b = b[l:]
		for i, j := range entry {
			fmt.Println(i, j)
		}

	}

}

// list
// modify
// delete
// import from csv
// export to csv
