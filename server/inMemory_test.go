package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
    server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), NewPostRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, NewGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	AssertResponseBody(t, response.Body.String(), "3")
}