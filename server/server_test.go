package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestGetPlayer(t *testing.T) {
	t.Run("Checking the GET for players endpoint",func(t *testing.T){
		request:=httptest.NewRequest(http.MethodGet,"/players/Pepper",nil)
		response:=httptest.NewRecorder()
		PlayerServer(response,request)
		got:=response.Body.String()
		want:="20"
		if got!=want{
			t.Errorf("Actual : %s , Expected : %s",got ,want)
		}

	})
}