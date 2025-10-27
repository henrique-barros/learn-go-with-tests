package main

import (
	"log"
	"net/http"

	server "example.com/hello/application/server"
)

func main() {
	server := server.NewPlayerServer(server.NewInMemoryStore())
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
