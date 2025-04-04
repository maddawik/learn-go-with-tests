package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/maddawik/learn-go-with-tests/command-line"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem initializing store for webserver, %v", err)
	}
	defer close()

	fmt.Println("Let's play some poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
