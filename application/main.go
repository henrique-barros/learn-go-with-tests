package main

import (
	"log"
	"net/http"

	server "example.com/hello/application/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
