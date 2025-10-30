package application

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {

	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()

	store, err := NewFileSystemStore(database)

	AssertNoError(t, err)

	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		AssertResponseStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		got := getLeagueFromResponse(t, response.Body)

		want := []Player{
			{
				Name: "Pepper",
				Wins: 3,
			},
		}

		AssertResponseStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, want)
	})
}
