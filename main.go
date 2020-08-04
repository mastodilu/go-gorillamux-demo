package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func apiMiddleware(next http.Handler) http.Handler {
	fmt.Println("using api middleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	// creo il router
	r := mux.NewRouter()
	subr := r.PathPrefix("/api").Subrouter()
	subr.Use(apiMiddleware)

	// definisco le rotte sul mio nuovo e scintillante router
	r.HandleFunc("/", HomeHandler)

	// queste sono le rotte /api/.... che restituiscono sempre JSON
	// il content type application json lo setto nel middleware
	subr.HandleFunc("/onlyget", OnlyGetHandler).Methods("GET")
	subr.HandleFunc("/onlypost", OnlyPostHandler).Methods("POST")
	subr.HandleFunc("/getandpost", GetAndPostHandler).Methods("GET", "POST")
	subr.HandleFunc("/api/products/", ProductsHandler)
	subr.HandleFunc("/products/{id:[0-9]+}", ProductsHandler) // <-- è utile usare le variabili nelle rotte per quando si fa RESTful routing

	// definisco che voglio usare il mio apprezzatissimo router
	http.Handle("/", r)

	// resto in ascolto
	fmt.Println("Listening on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
