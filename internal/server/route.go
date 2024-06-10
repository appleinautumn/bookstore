package server

import (
	"fmt"
	"net/http"
	"os"

	"gotu/bookstore/internal/handler"

	"github.com/gorilla/mux"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, os.Getenv("APP_NAME")+" "+os.Getenv("APP_ENV")+" "+os.Getenv("APP_VERSION"))
}

func NewRouter(apiHandler *handler.ApiHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/books", apiHandler.ListBooks).Methods("GET")
	r.HandleFunc("/", root)

	return r
}