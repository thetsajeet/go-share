package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thetsajeet/go-drop/handlers"
)

func main() {
	r := mux.NewRouter()

	r.
		HandleFunc("/", handlers.HandleHelloWorld).
		Methods("GET")

	r.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	log.
		Default().
		Println("Starting server on port: 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
