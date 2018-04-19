package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Entry struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

var entries []Entry

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/list", ListHandler)
	r.HandleFunc("/get", GetHandler)
	r.HandleFunc("/add", AddHandler)
	r.HandleFunc("/display", DisplayHandler)
	// r.HandleFunc("/modify", ModifyHandler )
	// r.HandleFunc("/delete", ModifyHandler )
	entries = append(entries, Entry{FirstName: "Jessica", LastName: "Bellon", Email: "J@B.com", Phone: "2148707789"})
	entries = append(entries, Entry{FirstName: "Will", LastName: "McGinnis", Email: "W@M.org", Phone: "1234449999"})
	entries = append(entries, Entry{FirstName: "Tyler", LastName: "Higgins", Email: "T@h.gov", Phone: "9876543210"})
	http.ListenAndServe(":8088", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index")
}

func ListHandler(w http.ResponseWriter, r *http.Request) {

	//t, err := template.ParseFiles("./list.html")
	t := template.Must(template.ParseFiles("./list.html"))
	t.Execute(w, entries)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./add.html") // display form
	case "POST": // process submitted form data
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Error with r.ParseForm()")
			return
		}
		e := Entry{}
		e.FirstName = r.FormValue("first_name")
		e.LastName = r.FormValue("last_name")
		e.Phone = r.FormValue("phone")
		e.Email = r.FormValue("email")

		entries = append(entries, e)
		fmt.Fprintf(w, "Success")

	default:
		fmt.Fprintf(w, "Attempting an unsupported action")
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./get.html")
}

func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	//		r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	for _, e := range entries {
		if e.FirstName == r.FormValue("name") || e.LastName == r.FormValue("name") {
			json.NewEncoder(w).Encode(e)
		}
	}
}
