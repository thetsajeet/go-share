package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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
			room = &rooms.Room{
				ID:          roomID,
				Connections: []*websocket.Conn{},
				UploadLock:  &sync.Mutex{},
			}
			cfg.Rooms[roomID] = room
		}
		cfg.RoomsLock.Unlock()

		room.UploadLock.Lock()
		room.Connections = append(room.Connections, conn)
		cfg.Rooms[roomID] = room
		room.UploadLock.Unlock()

		log.Default().Println("New client joined the room: ", roomID)

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			room.BroadcastToRoom(msg)
		}
	}
}
