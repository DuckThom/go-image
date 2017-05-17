package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func response(w http.ResponseWriter, message string, code int) {
	var data []byte
	var err error
	var link string

	if code == 200 {
		link = "http://" + host + ":" + port + "/?id=" + message
	}

	if returnJson {
		data, err = json.Marshal(map[string]string{
			"payload": message,
			"link": link,
		})
	}

	if returnText {
		data = []byte(message + "\n" + link)
	}

	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(code)
	w.Write(data)
}