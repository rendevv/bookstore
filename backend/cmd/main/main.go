package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rendevv/bookstore/backend/pkg/config"
	"github.com/rendevv/bookstore/backend/pkg/routes"
)

func main() {
	config.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to my Book API")
	}).Methods("GET")

	routes.RegisterBookStoreRoutes(r)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
