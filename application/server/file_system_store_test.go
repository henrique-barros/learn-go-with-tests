package application

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	var initialData = `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}
	]`

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)

		assertNoError(t, err)

		got := store.GetLeague()

		want := []Player{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}

		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)

		assertNoError(t, err)

		got, _ := store.GetPlayerScore("Chris")

		want := 33
		assertScoreEquals(t, want, got)
	})

	t.Run("record wins for an existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)

		assertNoError(t, err)

		store.RecordWin("Chris")
		got, _ := store.GetPlayerScore("Chris")
		want := 34

		assertScoreEquals(t, want, got)
	})

	t.Run("record wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

		store, err := NewFileSystemStore(database)

		assertNoError(t, err)

		store.RecordWin("John")
		got, _ := store.GetPlayerScore("John")
		want := 1

		assertScoreEquals(t, want, got)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemStore(database)

		assertNoError(t, err)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted score to be %d, got %d", want, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("wasn't expecting error, got %v", got)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Errorf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
