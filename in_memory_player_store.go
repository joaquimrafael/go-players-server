package main

import (
	"sync"

	Server "github.com/joaquimrafael/go-players-server/server"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}, mu: sync.Mutex{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []Server.Player {
	var league []Server.Player
	for name, wins := range i.store {
		league = append(league, Server.Player{Name: name, Wins: wins})
	}
	return league
}
