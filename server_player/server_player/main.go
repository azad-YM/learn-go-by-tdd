package main

import (
	"log"
	"net/http"

	game "example.com/learn/server_player"
)

func main() {
	server := game.NewPlayerServer(game.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5001", server))
}
