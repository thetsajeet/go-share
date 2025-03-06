package rooms

import "github.com/gorilla/websocket"

func (r *Room) BroadcastToRoom(msg []byte) {
	for _, conn := range r.Connections {
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}
