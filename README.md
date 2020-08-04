# Go GORILLA/MUX demo

## Risorse magari utili

[cors and middleware](https://stackoverflow.com/questions/51456253/how-to-set-http-responsewriter-content-type-header-globally-for-all-api-endpoint)

```go
func main() {
  port := ":3000"
  var router = mux.NewRouter()
  router.Use(commonMiddleware) // 💡

  router.HandleFunc("/m/{msg}", handleMessage).Methods("GET")
  router.HandleFunc("/n/{num}", handleNumber).Methods("GET")

  headersOk := handlers.AllowedHeaders([]string{"Authorization"})
  originsOk := handlers.AllowedOrigins([]string{"*"})
  methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

  fmt.Printf("Server is running at http://localhost%s\n", port)
  log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router))) // 💡
}

 // 💡
func commonMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    next.ServeHTTP(w, r)
  })
}
```
