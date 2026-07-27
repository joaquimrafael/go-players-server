package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joaquimrafael/go-players-server/filesystem"
	"github.com/joaquimrafael/go-players-server/server"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := filesystem.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := server.NewPlayerServer(store)

	if err := http.ListenAndServe(":5500", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
