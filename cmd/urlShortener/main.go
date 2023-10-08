package main

import (
	"urlShortener/internal/server"
	"urlShortener/internal/storage"
)

func main() {
	storage.InitDB("database/db.sqlite")
	server.Start(80)
}
