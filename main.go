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
		fmt.Fprintf(w, "AddressBook v0.2")
	})

	// List all entries
	r.HandleFunc("/entry", GetAll)
	// Get Specific Entry
	r.HandleFunc("/entry/{id:[0-9]+}", GetEntry).Methods("GET")      // single entry by ID
	r.HandleFunc("/entry/{name:[a-zA-Z]+}", GetEntry).Methods("GET") // One or more entries by first/last name
	// Create an Entry (using simple form)
	r.HandleFunc("/add", CreateEntry)
	// TODO modify (PUT)
	// r.HandleFunc("/entry/{id}", ModifyEntry).Methods("PUT")
	// TODO delete (DELETE)
	//r.HandleFunc("/entry/{id}", DeleteEntry).Methods("DELETE")

	// TODO import from csv
	// TODO export to csv
	//entries = append(entries, Entry{ID: "1", FirstName: "John", LastName: "Doe", Email: "jd@gmail.com", Phone: "214-009-9000"})
	//entries = append(entries, Entry{ID: "2", FirstName: "Jane", LastName: "Doe", Email: "djd@gmail.com", Phone: "432-222-2122"})
	//entries = append(entries, Entry{ID: "3", FirstName: "Jessica", LastName: "Bellon", Email: "jess@jessicabellon.com", Phone: "214-870-7789"})
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

	switch r.Method {
	case "GET": // display html form
		fmt.Fprintf(w, POSTPage)
	case "POST": // process form data
		// ParseForm() parses raw query data and updates r.PostForm and r.Form
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Error with r.ParseForm")
			return
		}
		e := Entry{ID: string(len(entries))} // TODO this is aweful
		e.FirstName = r.FormValue("first_name")
		e.LastName = r.FormValue("last_name")
		// e.Phone
		// e.Email
		entries = append(entries, e)
		fmt.Fprintf(w, "Successfully added an entry. Visit /entry to see the full list")
	default:
		fmt.Fprintf(w, "Attempting an unsupported action, Gorilla Mux should have prevented this")
	}
}
func ModifyEntry(w http.ResponseWriter, r *http.Request) {}
func DeleteEntry(w http.ResponseWriter, r *http.Request) {}

var POSTPage = `
<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8" />
	</head>
	<body>
		<div>
			<form method="POST" action="/add">
				<p>First Name: </p>
				<input name="first_name" type="text" >
				<br>
				<p> Last Name: </p>
				<input name="last_name" type="text" >
				<input type="submit" value="submit"/>
				</form>
		</div>
	</body>
	</html>
`
