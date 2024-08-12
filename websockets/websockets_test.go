package websockets

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)


func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := NewPlayerServer(&StubPlayerStore{})

		request:=newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("when we get a message over a websocket it is a winner of a game",func(t *testing.T) {
		store:=&StubPlayerStore{}
		winner:="Numeez"
		server:=httptest.NewServer(NewPlayerServer(store))
		defer server.Close()
		wsUrl:="ws"+strings.TrimPrefix(server.URL,"http")+"/ws"
		ws,_,err:=websocket.DefaultDialer.Dial(wsUrl,nil)
		if err!=nil{
			t.Fatalf("could not open a ws connection on %s %v", wsUrl, err)
		}


		defer ws.Close()
		if err:= ws.WriteMessage(websocket.TextMessage,[]byte(winner));err!=nil{
			t.Fatalf("could not send message over ws connection %v", err)
		}
		time.Sleep(10 * time.Millisecond)
		assertPlayerWin(t,store,winner)

	})
}


func assertStatus(t testing.TB,got ,want int){
	t.Helper()
	if got!=want{
		t.Errorf("Actual : %d , Expected : %d",got ,want)
	}
}

func assertPlayerWin(t testing.TB, store *StubPlayerStore,winner string ){
	t.Helper()
	got:=store.WinCalls[0]
	if got!=winner{
		t.Errorf("Actual : %s , Expected : %s",got,winner)
	}
}

func newGameRequest()*http.Request{
	req,_:= http.NewRequest(http.MethodGet,"/game",nil)
	return req
}