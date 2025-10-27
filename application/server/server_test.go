package application

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {
	return s.scores[name], nil
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.scores[name] = s.scores[name] + 1
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	stubPlayerStore := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := NewPlayerServer(&stubPlayerStore)
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)

	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Chris")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusNotFound)

	})
}

func TestStoreWins(t *testing.T) {
	stubPlayerStore := &StubPlayerStore{
		scores: map[string]int{},
	}
	server := NewPlayerServer(stubPlayerStore)

	t.Run("should increment player store", func(t *testing.T) {
		request := newPostScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusAccepted)

		request = newGetScoreRequest("Pepper")
		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "2")
		if len(stubPlayerStore.winCalls) != 2 {
			t.Errorf("got %d calls to RecordWin want %d", len(stubPlayerStore.winCalls), 2)
		}

		if stubPlayerStore.winCalls[0] != "Pepper" {
			t.Errorf("did not score correct winner got %q, want %q", stubPlayerStore.winCalls[0], "Pepper")
		}
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	playerStore := StubPlayerStore{
		league: wantedLeague,
	}
	playerServer := NewPlayerServer(&playerStore)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newLeagueRequest()
		playerServer.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		assertContentType(t, response, jsonContentType)
		assertResponseStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
	})
}

func newGetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+player, nil)
	return request
}

func getLeagueFromResponse(t testing.TB, body io.Reader) []Player {
	t.Helper()
	var league []Player
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Errorf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return league
}

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %q, got %v", want, response.Result().Header)
	}
}

func newPostScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/players/"+player, nil)
	return request
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertResponseStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got response status %d, want %d", got, want)
	}
}
