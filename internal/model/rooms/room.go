package rooms

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	Connections []*websocket.Conn
	ID          string
	UploadLock  *sync.Mutex
}
