package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/maddawik/learn-go-with-tests/command-line"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play some poker")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s, %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem initializing player store, %v", err)
	}

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
