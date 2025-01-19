package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rohitbisht/go-book-management-system/pkg/models"
	"github.com/rohitbisht/go-book-management-system/pkg/utils"
)

var NewBook models.Book

// get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// get book by id
func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		log.Println("Error converting bookId to int64: ", err)
		return
	}

	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// create new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	book := CreateBook.CreateBook()

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		log.Println("Error converting bookId to int:", err)
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// update a book
func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		log.Println("Error converting bookId to int:", err)
		return
	}

	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)

	book, _ := models.GetBookById(id)

	if book.Name != "" {
		book.Name = updatedBook.Name
	}
	if book.Author != "" {
		book.Author = updatedBook.Author
	}
	if book.Publication != "" {
		book.Publication = updatedBook.Publication
	}

	// save the changes
	updatedBook, err = models.UpdateBook(book)
	res, _ := json.Marshal(updatedBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
