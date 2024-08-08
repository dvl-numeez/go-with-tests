package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		if err!=nil{
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func AssertResponseBody(t testing.TB,got,want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	
}

func NewGetScoreRequest(name string)*http.Request{
	url:=fmt.Sprintf("/players/%s",name)
	req,_:= http.NewRequest(http.MethodGet, url, nil)
	return req
}
func NewPostRequest(name string)*http.Request{
	url:=fmt.Sprintf("/players/%s",name)
	req,_:=http.NewRequest(http.MethodPost, url, nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
