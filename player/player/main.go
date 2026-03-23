package main

import (
	"log"
	"net/http"

	"example.com/learn/player"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerStore(name string) int {
	return 123
}

func main() {
	playerServer := &player.PlayerServer{
		Store: &InMemoryPlayerStore{},
	}

	log.Fatal(http.ListenAndServe(":5000", playerServer))
}
