package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thetsajeet/go-drop/internal/config"
	"github.com/thetsajeet/go-drop/internal/model/rooms"
)

func HandleWebSocket(cfg *config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		roomID := vars["roomID"]
		if roomID == "" {
			http.Error(w, "Room ID is required", http.StatusBadRequest)
			return
		}

		conn, err := cfg.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade failed:", err)
			http.Error(w, "WebSocket upgrade failed", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		cfg.RoomsLock.Lock()
		room, exists := cfg.Rooms[roomID]
		if !exists {
			room = rooms.CreateRoom(roomID)
			cfg.Rooms[roomID] = room
		}
		cfg.RoomsLock.Unlock()

		cfg.Rooms[roomID] = room.AddConnection(conn)

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			room.BroadcastToRoom(msg)
		}
	}
}
