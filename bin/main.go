package main

import (

  "net/http"

  "fmt"

  "github.com/gorilla/mux"

  "github.com/rs/cors"
)


func main() {
  router := mux.NewRouter()

  router.HandleFunc("/", GetHome).Methods("GET")

  handler := cors.Default().Handler(router)

  http.ListenAndServe(":80", handler)
}

func GetHome(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to Golang")
}
