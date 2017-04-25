package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func response(w http.ResponseWriter, message string, code int) {
	var data []byte
	var err error

	if returnJson {
		data, err = json.Marshal(map[string]string{
			"payload": message,
		})
	}

	if returnText {
		data = []byte(message)
	}

	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(code)
	w.Write(data)
}