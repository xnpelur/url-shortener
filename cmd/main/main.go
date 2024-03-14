package main

import (
	"fmt"
	"os"
	"strconv"
	"url-shortener/internal/server"
	"url-shortener/internal/storage"
)

func main() {
	port := 80
	if len(os.Args) > 1 {
		var err error
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error while parsing command line argument, using default port value (80)")
			port = 80
		}
	}
	storage.InitDB("database/db.sqlite")
	server.Start(port)
}
