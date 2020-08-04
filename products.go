package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ProductsHandler handler for the products page
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	// qua trovo la var id presa da localhost:8080/products/{id}
	vars := mux.Vars(r)

	response := struct{ ID string }{ID: vars["id"]}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}
