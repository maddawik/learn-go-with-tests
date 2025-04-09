package main

import (
	"log"
	"net/http"

	poker "github.com/maddawik/learn-go-with-tests/websockets"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem initializing store for webserver, %v", err)
	}
	defer close()

	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
