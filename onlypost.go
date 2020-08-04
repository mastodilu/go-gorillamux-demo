package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// OnlyPostHandler handler function for 'only POST' page
func OnlyPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	err := json.NewEncoder(w).Encode(struct {
		RequestMethod, Message string
	}{
		RequestMethod: r.Method,
		Message:       "this page only accepts POST requests.\nIt sends status code 405 to other request types.",
	})
	if err != nil {
		log.Println(err)
	}
}
