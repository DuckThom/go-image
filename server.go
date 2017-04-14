package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	fmt.Println("Starting server...")

	http.HandleFunc("/", handleConnection)

	fmt.Println("Server is running on " + ip + ":" + port + "\nExit with Ctrl-C")

	http.ListenAndServe(ip+":"+port, nil)
}
