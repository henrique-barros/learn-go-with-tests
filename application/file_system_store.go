package application

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) (int, error) {
	player := f.league.find(name)

	if player != nil {
		return player.Wins, nil
	}

	return 0, nil
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.league
	player := league.find(name)

	if player != nil {
		player.Wins++
	} else {
		newPlayer := Player{
			Wins: 1,
			Name: name,
		}
		f.league = append(league, newPlayer)
	}

	f.database.Encode(f.league)
}

func NewFileSystemStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	fileSystemPlayerStore, err := NewFileSystemStore(db)

	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store for file %s %v", path, err)
	}

	return fileSystemPlayerStore, closeFunc, nil
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}

	return nil
}
