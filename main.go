package main

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/thetsajeet/go-share/cmd/server"
	"github.com/thetsajeet/go-share/internal/config"
	"github.com/thetsajeet/go-share/internal/model/rooms"
)

func main() {
	cfg := &config.AppConfig{
		Rooms:        make(map[string]*rooms.Room, 0),
		Upgrader:     websocket.Upgrader{},
		RoomsLock:    &sync.Mutex{},
		RStoragePath: "",
	}
	cfg.InitConfig()

	server.StartServer(cfg)
}
