package routes

import (
	"github.com/gorilla/mux"
	"github.com/rendevv/bookstore/backend/pkg/controllers"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	s := router.PathPrefix("/book").Subrouter()

	s.HandleFunc("/", controllers.CreateBook).Methods("POST")
	s.HandleFunc("/", controllers.GetBook).Methods("GET")
	s.HandleFunc("/{bookId}", controllers.GetBookById).Methods("GET")
	s.HandleFunc("/{bookId}", controllers.UpdateBook).Methods("PUT")
	s.HandleFunc("/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
