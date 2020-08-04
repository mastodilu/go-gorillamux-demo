package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// OnlyGetHandler handler for the 'only-GET' page
func OnlyGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	err := json.NewEncoder(w).Encode(struct {
		RequestMethod, Message string
	}{
		RequestMethod: r.Method,
		Message:       "this page only accepts GET requests.\nIt sends status code 405 to other request types.",
	})
	if err != nil {
		log.Println(err)
	}
}
