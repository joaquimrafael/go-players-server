package main

import (
	"log"
	"net/http"

	"github.com/joaquimrafael/go-players-server/domain"
	"github.com/joaquimrafael/go-players-server/filesystem"
	"github.com/joaquimrafael/go-players-server/server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := filesystem.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := domain.NewTexasHoldem(domain.BlindAlerterFunc(domain.Alerter), store)

	playerServer, err := server.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("could not create player server: %v", err)
	}

	if err := http.ListenAndServe(":5500", playerServer); err != nil {
		log.Fatalf("could not listen on port 5500 %v", err)
	}
}
