package main

import (
	"log"
	"net/http"
	"os"

	poker "github.com/maddawik/learn-go-with-tests/command-line"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s, %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem initializing player store, %v", err)
	}
	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
