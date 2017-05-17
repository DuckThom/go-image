package main

import (
	"flag"
)

var ip string
var port string
var host string
var uploadDir string
var verbose bool
var key string
var returnFormat string

func main() {
	flag.StringVar(&ip, "ip", "0.0.0.0", "IP Address to bind the server to")
	flag.StringVar(&port, "port", "8080", "Port to bind the server to")
	flag.StringVar(&key, "key", "random", "API Key to use, when set to 'random', a uuid will be generated and used")
	flag.StringVar(&host, "host", "example.com", "Hostname of the server")
	flag.StringVar(&uploadDir, "dir", "./uploads", "Location where the uploaded images are stored")
	flag.StringVar(&returnFormat, "format", "json", "Change the default return type ('plain' or 'json')")
	flag.BoolVar(&verbose, "v", false, "Verbose logging")
	flag.Parse()

	startServer()
}
