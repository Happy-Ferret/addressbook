package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/JessTheBell/addressbook/model"
	//	"github.com/alecthomas/template"
)

var book model.Book

func main() {
	//	populate() // populate with example entries
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/get", getHandler)       // get and display
	http.HandleFunc("/add", addHandler)       // form to add an entry
	http.HandleFunc("/modify", modifyHandler) // modify an existing entry
	http.HandleFunc("/delete", deleteHandler) // delete an existing entry
	http.HandleFunc("/export", exportHandler) // export book to csv file
	http.HandleFunc("/import", importHandler) // import existing csv file
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func populate() {
	p := map[int][]string{
		0: {"Jessica", "Bellon", "214-555-9999", "Jess@Gmail.com"},
		1: {"Will", "McGinnis", "991-909-0123", "btown@email.com"},
		2: {"Tyler", "Higgins", "111-111-1111", "thiggs@mail.com"},
		3: {"Zelda", "Bellon", "232-909-9998", "Zelda@dogs.com"}}
	for _, e := range p {
		ent, err := model.NewEntry(e[0], e[1], e[2], e[3])
		if err != nil {
			break
		}
		book.Save(ent)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./pages/list.tmpl"))
	tmpl.Execute(w, book.All())
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Error parsing request")
		} else {
			entry, err := model.NewEntry(r.FormValue("first_name"), r.FormValue("last_name"),
				r.FormValue("phone"), r.FormValue("email"))
			if err != nil {
				fmt.Fprintf(w, "Error creating entry")
				return
			}
			book.Save(entry)
			http.ServeFile(w, r, "./pages/link.html")
		}
	default:
		http.ServeFile(w, r, "./pages/add.html")
	}
}

func modifyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("./pages/modify.tmpl"))
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "error parsing request")
		}

		strID := r.FormValue("id")
		id, err := strconv.ParseUint(strID, 10, 64)
		if err != nil {
			fmt.Fprintf(w, "error parsing id")
		}
		entry, _ := book.Get(id)
		tmpl.Execute(w, entry)
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "error parsing request")
		}
		strID := r.FormValue("id")
		id, err := strconv.ParseUint(strID, 10, 64)
		if err != nil {
			fmt.Fprintf(w, "error parsing id")
		}
		entry, ok := book.Get(id)
		if !ok {
			fmt.Fprintf(w, "Error getting book with id: %v", id)
			return
		}
		entry.FirstName = r.FormValue("first_name")
		entry.LastName = r.FormValue("last_name")
		entry.Phone = r.FormValue("phone")
		entry.Email = r.FormValue("email")
		book.Save(&entry)
		http.ServeFile(w, r, "./pages/link.html")
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "error parsing request")
	}
	strID := r.FormValue("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "error parsing id")
	}
	book.Delete(id)
	http.ServeFile(w, r, "./pages/link.html")
}

// handles get and display
func getHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./pages/get.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "error parsing request")
			return
		}
		if key := r.FormValue("name"); key != "" {
			// searching by name
			matches := book.Search(key)
			json.NewEncoder(w).Encode(matches)
		} else if key = r.FormValue("id"); key != "" {
			// searching by id
			id, err := strconv.ParseUint(key, 10, 64)
			if err != nil {
				fmt.Fprintf(w, "error parsing id")
			}
			entry, _ := book.Get(id)
			json.NewEncoder(w).Encode(entry)

		} else {
			fmt.Fprintf(w, "No entry was found with those search parameters")
		}
	}
}

func importHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./pages/upload.html")
	case "POST":
		file, _, err := r.FormFile("csvfile")
		if err != nil {
			fmt.Println("ERROR")
		}
		defer file.Close()
		csvReader := csv.NewReader(file)
		lines, _ := csvReader.ReadAll()
		if err != nil {
			fmt.Fprintf(w, "error parsing id")
		}
		for _, e := range lines {

			entry, err := model.NewEntry(e[0], e[1], e[2], e[3])
			if err != nil {
				break
			}
			book.Save(entry)
		}
		http.ServeFile(w, r, "./pages/link.html")
	}

}

func exportHandler(w http.ResponseWriter, r *http.Request) {

	filename := "./addressbook.csv"
	err := book.Export()
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=addressbook.csv")
	w.Header().Set("Content-Type", "text/csv")

	//stream the body to the client
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(w, bufio.NewReader(f))
	// delete the csv file off of the local machine when done
	err = os.Remove(filename)
	if err != nil {
		fmt.Println("Error deleting file: %v", err)
	}
}
