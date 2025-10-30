package main

import (
	"fmt"
	"log"
	"os"

	app "example.com/hello/application"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := app.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating new file system store, %v ", err)
	}
	defer close()
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := app.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
