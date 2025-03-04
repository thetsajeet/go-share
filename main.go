package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}).Methods("GET")

	log.Default().Println("Starting server on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
