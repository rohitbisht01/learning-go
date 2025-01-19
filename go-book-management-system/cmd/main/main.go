package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/postgres"

	"github.com/rohitbisht/go-book-management-system/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookToStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:4000", r))
}
