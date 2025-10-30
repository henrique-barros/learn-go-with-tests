package application

import (
	"encoding/json"
	"io"
)

type League []Player

func NewLeague(reader io.Reader) (League, error) {
	var players []Player
	err := json.NewDecoder(reader).Decode(&players)
	if err != nil {
		return nil, err
	}
	return League(players), nil
}

func (l League) find(player string) *Player {
	for i, val := range l {
		if val.Name == player {
			return &l[i]
		}
	}

	return nil
}
