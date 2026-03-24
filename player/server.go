package player

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerStore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)

	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.Store.GetPlayerStore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
