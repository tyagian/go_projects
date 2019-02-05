package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	router := mux.NewRouter()
	books = append(books,
		Book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
		Book{ID: 3, Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: 4, Title: "Golang concurrency", Author: "Mr. Concurrency", Year: "2013"},
		Book{ID: 5, Title: "Golang good parts", Author: "Mr. Good", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	//log.Println(reflect.TypeOf(params["i"]))
	// using strconv to convert str to int
	//log.Println(reflect.TypeOf(i))

	for _, book := range books {
		// throw error if we use id as int here in params
		// we need to convert id into string before using in params
		//  we have a package called reflect to check type of var
		// strconv used to convert str to integer
		//if book.ID == params["id"] {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}
func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add a book is called")
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update a book is called")
}
func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete a book is called")
}
