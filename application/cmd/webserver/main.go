package main

import (
	"log"
	"net/http"

	server "example.com/hello/application"
)

const dbFileName = "game.db.json"

func main() {
	fileSystemStore, cleanDatabase, err := server.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	defer cleanDatabase()

	server := server.NewPlayerServer(fileSystemStore)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
