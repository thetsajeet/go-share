package rooms

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

func (r *Room) BroadcastToRoom(msg []byte) {
	for _, conn := range r.Connections {
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func CreateRoom(roomID string) *Room {
	return &Room{
		ID:          roomID,
		UploadLock:  &sync.Mutex{},
		Connections: []*websocket.Conn{},
	}
}

func (r *Room) AddConnection(conn *websocket.Conn) *Room {
	r.UploadLock.Lock()
	defer r.UploadLock.Unlock()

	r.Connections = append(r.Connections, conn)
	log.Default().Printf("New client added to room: %v", r.ID)
	return r
}
