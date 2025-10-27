package application

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var ErrorNotFound = errors.New("player score not found")

type PlayerServer struct {
	store PlayerStore
}

type InMemoryPlayerStore struct {
	score map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	return i.score[name], nil
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.score[name]++
}

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	RecordWin(name string)
}

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")

	switch request.Method {
	case http.MethodPost:
		p.processWin(response, player)
	case http.MethodGet:
		p.showScore(response, player)
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
	return &PlayerServer{store}
}

func NewInMemoryStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{make(map[string]int)}
}
