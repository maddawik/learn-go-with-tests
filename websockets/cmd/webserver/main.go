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

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatal("problem initializing player server", err)
	}

	log.Fatal(http.ListenAndServe(":5000", server))
}
