package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetAndPostHandler handler for get and post methods of 'GETandPOST' page
func GetAndPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getandpost")
	err := json.NewEncoder(w).Encode(
		struct {
			Method, Message, Host, Path string
		}{
			Method:  r.Method,
			Host:    r.URL.Host,
			Path:    r.URL.Path,
			Message: "This page only accepts GET and POST requests. In other cases it responds with 405.",
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
	}
}
