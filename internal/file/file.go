package f

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/thetsajeet/go-drop/internal/config"
)

func HandleUploadFile(cfg *config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		roomID := vars["roomID"]
		if roomID == "" {
			http.Error(w, "Room ID is required", http.StatusBadRequest)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Failed to read file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		roomPath := filepath.Join(cfg.RStoragePath, roomID)
		if err := os.MkdirAll(roomPath, os.ModePerm); err != nil {
			http.Error(w, "Unable to create directory", http.StatusInternalServerError)
			return
		}

		dstPath := filepath.Join(roomPath, header.Filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, "Unable to create the file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Failed to save the file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "File uploaded successfull to room %s", roomID)
	}
}

func HandleDownloadFile(w http.ResponseWriter, r *http.Request) {}
