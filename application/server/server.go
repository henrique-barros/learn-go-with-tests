package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var ErrorNotFound = errors.New("player score not found")

const jsonContentType = "application/json"

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type InMemoryPlayerStore struct {
	score map[string]int
}

type Player struct {
	Name string
	Wins int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	return i.score[name], nil
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.score[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.score {
		league = append(league, Player{
			Name: name,
			Wins: wins,
		})
	}
	return league
}

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	RecordWin(name string)
	GetLeague() []Player
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(response http.ResponseWriter, player string) {
	playerScore, err := p.store.GetPlayerScore(player)

	if playerScore == 0 || err == ErrorNotFound {
		response.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(response, playerScore)
}

func GetPlayerScore(player string) (int, error) {
	if player == "Floyd" {
		return 10, nil
	}

	if player == "Pepper" {
		return 20, nil
	}

	return 0, ErrorNotFound
}

func AddWin(player string) {

}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.leagueHandler))

	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router
	return p
}

func NewInMemoryStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		score: make(map[string]int),
	}
}
