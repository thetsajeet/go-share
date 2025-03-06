package config

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/thetsajeet/go-drop/internal/model/rooms"
)

type AppConfig struct {
	Rooms     map[string]*rooms.Room
	Upgrader  websocket.Upgrader
	RoomsLock *sync.Mutex
}
