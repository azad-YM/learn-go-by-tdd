package player

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerStore(name string) int
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.Store.GetPlayerStore(player))
}
