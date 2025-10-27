package application

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {
	return s.scores[name], nil
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.scores[name] = s.scores[name] + 1
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	stubPlayerStore := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := &PlayerServer{&stubPlayerStore}
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
	server := &PlayerServer{stubPlayerStore}

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

func newGetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/players/"+player, nil)
	return request
}

func newPostScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, "/players/"+player, nil)
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
