package server_player

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Mina"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinsRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinsRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinsRequest(player))
	response := httptest.NewRecorder()

	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
