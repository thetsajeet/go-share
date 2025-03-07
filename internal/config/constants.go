package config

import (
	"log"
	"os"
)

var (
	MaxRoomSize     int
	MaxFileSize     int
	MaxFilesPerRoom int
	RoomsStorage    string = "/rooms"
)

func (cfg *AppConfig) InitConfig() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("can't get current working directory: %v", err)
	}
	cfg.RStoragePath = wd + RoomsStorage
}
