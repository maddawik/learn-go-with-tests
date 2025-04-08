package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/maddawik/learn-go-with-tests/time"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem initializing store for webserver, %v", err)
	}
	defer close()

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.StdOutAlerter), store)

	fmt.Println("Let's play some poker")
	fmt.Println("Type {Name} wins to record a win")
	cli := poker.NewCLI(os.Stdin, os.Stdin, game)
	cli.PlayPoker()
}
