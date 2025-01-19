package routes

import (
	"github.com/gorilla/mux"
	"github.com/rohitbisht/go-book-management-system/pkg/controllers" // absolute path
)

var RegisterBookToStore = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBookById).Methods("PUT")
}
