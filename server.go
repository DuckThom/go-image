package main

import (
	"net/http"
	"github.com/google/uuid"
	"log"
)

func startServer() {
	log.Println("Starting server...")

	http.HandleFunc("/", handleConnection)

	if key == "random" {
		key = uuid.New().String()
	}

	log.Println("API Key set to: " + key)
	log.Println("Server is running on " + ip + ":" + port + "\nExit with Ctrl-C")

	panic(http.ListenAndServe(ip+":"+port, nil))
}
