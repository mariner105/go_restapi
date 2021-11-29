package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Book Struct (Model)
type Book struct {
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	TITLE string `json:"title"`
	AUTHOR *Author `json:"author"`
}

// Author Struct
type Author struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}


func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock data - @todo - implement DB
	books = append(books, Book{ID: "1", ISBN: "448743", TITLE: "Book One", AUTHOR: &Author{
		FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", ISBN: "8477564", TITLE: "Book Two", AUTHOR: &Author{
		FirstName: "Steve", LastName: "Smith"}})

	// Route handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}