package main

import (
	"log"
	"net/http"

	"example.com/learn/player"
)

func main() {
	server := &player.PlayerServer{
		Store: player.NewInMemoryPlayerStore(),
	}

	log.Fatal(http.ListenAndServe(":5001", server))
}
