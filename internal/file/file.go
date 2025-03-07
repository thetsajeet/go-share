package f

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/thetsajeet/go-share/internal/config"
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

func HandleDownloadFile(cfg *config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		roomID := vars["roomID"]
		fileName := vars["fileName"]
		if roomID == "" || fileName == "" {
			http.Error(w, "room id and filename are required", http.StatusBadRequest)
			return
		}

		filePath := filepath.Join(cfg.RStoragePath, roomID, fileName)
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "file not found", http.StatusBadRequest)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		w.Header().Set("Content-Type", "application/octet-stream")

		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "error while downloading the file", http.StatusInternalServerError)
			return
		}
	}
}
