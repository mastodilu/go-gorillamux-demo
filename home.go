package main

import (
	"fmt"
	"net/http"
)

// HomeHandler handler of the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home")
}
