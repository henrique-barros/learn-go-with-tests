package application

type InMemoryPlayerStore struct {
	score map[string]int
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

func NewInMemoryStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		score: make(map[string]int),
	}
}
