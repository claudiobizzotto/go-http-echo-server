package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func getServerPort() string {
	port := "8080"
	if len(os.Args) > 1 {
		port = strings.TrimSpace(os.Args[1])
	}

	return port
}

func handleRequest(responseWriter http.ResponseWriter, request *http.Request) {
	log.Println("Received request from " + request.RemoteAddr)
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")
	request.Write(responseWriter)
}

func main() {
	port := getServerPort()
	log.Println("Starting HTTP echo server. Listening on port " + port)
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":"+port, nil)
}
