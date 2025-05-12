package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rendevv/bookstore/backend/pkg/config"
	"github.com/rendevv/bookstore/backend/pkg/routes"
)

func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	config.Connect()
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(mux.CORSMethodMiddleware(r))
	routes.RegisterBookStoreRoutes(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to my Book API")
	}).Methods("GET")

	// Wrap the router with CORS middleware
	handlerWithCORS := enableCORS(r)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCORS))
}

