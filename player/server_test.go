package player

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	score map[string]int
}

func (s *StubPlayerStore) GetPlayerStore(name string) int {
	return s.score[name]
}

func TestGetPlayers(t *testing.T) {
	server := PlayerServer{
		Store: &StubPlayerStore{
			score: map[string]int{
				"Pepper": 20,
				"Azad":   10,
			},
		},
	}

	t.Run("Get Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("Get Azad's score", func(t *testing.T) {
		request := newGetScoreRequest("Azad")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
