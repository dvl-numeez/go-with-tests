package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestGetPlayer(t *testing.T) {
	t.Run("Checking the GET for players endpoint",func(t *testing.T){
		request:=NewGetScoreRequest("Pepper")
		response:=httptest.NewRecorder()
		PlayerServer(response,request)
		AssertResponseBody(t,response.Body.String(),"20")

	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request:=NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
	
		PlayerServer(response, request)
	
		AssertResponseBody(t,response.Body.String(),"10")
	
		
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