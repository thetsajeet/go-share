package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thetsajeet/go-drop/internal/config"
	f "github.com/thetsajeet/go-drop/internal/file"
	"github.com/thetsajeet/go-drop/internal/hello.go"
	ws "github.com/thetsajeet/go-drop/internal/websocket"
)

func StartServer(cfg *config.AppConfig) {
	r := mux.NewRouter()

	r.
		HandleFunc("/", hello.HandleHelloWorld).
		Methods("GET")

	r.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	r.
		HandleFunc("/rooms/{roomID}", ws.HandleWebSocket(cfg)).
		Methods("GET")

	r.
		HandleFunc("/rooms/{roomID}/upload", f.HandleUploadFile(cfg)).
		Methods("POST")

	r.
		HandleFunc("/rooms/{roomID}/download/{fileName}", f.HandleDownloadFile(cfg)).
		Methods("GET")

	log.
		Default().
		Println("Starting server on port: 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
