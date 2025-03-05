package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var rooms = make(map[string][]*websocket.Conn)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["roomID"]
	if roomID == "" {
		http.Error(w, "Room ID is required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		http.Error(w, "WebSocket upgrade failed", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	rooms[roomID] = append(rooms[roomID], conn)
	log.Default().Println("New client joined the room: ", roomID)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		broadcastToRoom(roomID, msg)
	}
}

func broadcastToRoom(roomID string, msg []byte) {
	for _, conn := range rooms[roomID] {
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}
