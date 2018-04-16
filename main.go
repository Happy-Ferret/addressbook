package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Entry is an entry in the addressbook
type Entry struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

var entries []Entry

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "AddressBook v0.1")
	})

	// List all entries
	r.HandleFunc("/entry", GetAll).Methods("GET")
	// Get Specific Entry
	r.HandleFunc("/entry/{id:[0-9]+}", GetEntry).Methods("GET")
	r.HandleFunc("/entry/{name:[a-zA-Z]+}", GetEntry).Methods("GET")

	// TODO modify (PUT)
	// r.HandleFunc("/entry/{id}", ModifyEntry).Methods("PUT")
	// TODO delete (DELETE)
	//r.HandleFunc("/entry/{id}", DeleteEntry).Methods("DELETE")

	// TODO import from csv
	// TODO export to csv
	entries = append(entries, Entry{ID: "1", FirstName: "John", LastName: "Doe", Email: "jd@gmail.com", Phone: "214-009-9000"})
	entries = append(entries, Entry{ID: "2", FirstName: "Jane", LastName: "Doe", Email: "djd@gmail.com", Phone: "432-222-2122"})
	entries = append(entries, Entry{ID: "3", FirstName: "Jessica", LastName: "Bellon", Email: "jess@jessicabellon.com", Phone: "214-870-7789"})
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}

func GetEntry(w http.ResponseWriter, r *http.Request) {
	var matches []Entry
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	_, hasName := params["name"]
	_, hasID := params["id"]
	for _, e := range entries {
		if hasName {
			if e.FirstName == params["name"] || e.LastName == params["name"] {
				matches = append(matches, e)
			}
		}
		if hasID && e.ID == params["id"] {
			matches = append(matches, e)
		}

	}
	json.NewEncoder(w).Encode(matches)
}

func CreateEntry(w http.ResponseWriter, r *http.Request) {
}
func ModifyEntry(w http.ResponseWriter, r *http.Request) {}
func DeleteEntry(w http.ResponseWriter, r *http.Request) {}
