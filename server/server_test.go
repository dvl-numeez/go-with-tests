package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestGetPlayer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server:=&PlayerServer{&store}
	t.Run("Checking the GET for players endpoint",func(t *testing.T){
		request:=NewGetScoreRequest("Pepper")
		response:=httptest.NewRecorder()
		server.ServeHTTP(response,request)
		assertStatus(t,response.Code,http.StatusOK)
		AssertResponseBody(t,response.Body.String(),"20")

	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request:=NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
	
		server.ServeHTTP(response, request)
		assertStatus(t,response.Code,http.StatusOK)
		AssertResponseBody(t,response.Body.String(),"10")
	
		
})
t.Run("returns 404 on missing players", func(t *testing.T) {
	request := NewGetScoreRequest("Apollo")
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	assertStatus(t,response.Code,http.StatusNotFound)
})

}

func TestStoreWins(t *testing.T){
	store:=StubPlayerStore{
		map[string]int{},
	}
	server:=&PlayerServer{&store}
	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
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

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
