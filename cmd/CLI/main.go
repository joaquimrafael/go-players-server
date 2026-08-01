package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joaquimrafael/go-players-server/domain"
	"github.com/joaquimrafael/go-players-server/filesystem"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := domain.NewGame(domain.BlindAlerterFunc(domain.StdOutAlerter), store)
	cli := domain.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
