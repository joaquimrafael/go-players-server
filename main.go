package main

import (
	"log"
	"net/http"

	Server "github.com/joaquimrafael/go-players-server/server"
)

func main() {
	server := Server.NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5500", server))
}
