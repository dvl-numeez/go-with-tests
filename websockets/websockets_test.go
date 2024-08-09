package websockets

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := NewPlayerServer(&StubPlayerStore{})

		request:=newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}


func assertStatus(t testing.TB,got ,want int){
	t.Helper()
	if got!=want{
		t.Errorf("Actual : %d , Expected : %d",got ,want)
	}
}

func newGameRequest()*http.Request{
	req,_:= http.NewRequest(http.MethodGet,"/game",nil)
	return req
}