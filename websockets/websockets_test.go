package websockets

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dvl-numeez/go-with-tests/websockets/poker"
	"github.com/gorilla/websocket"
)


func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		game := &GameSpy{}
		server := mustMakePlayerServer(t,&StubPlayerStore{},game)
		request:=newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("when we get a message over a websocket it is a winner of a game",func(t *testing.T) {
		store:=&StubPlayerStore{}
		winner:="Numeez"
		game := &GameSpy{}
		server:=httptest.NewServer(mustMakePlayerServer(t,store,game))
		defer server.Close()
		wsUrl:="ws"+strings.TrimPrefix(server.URL,"http")+"/ws"
		ws:=mustDialWS(t,wsUrl)
		writeWSMessage(t,ws,winner)
		defer ws.Close()
		if err:= ws.WriteMessage(websocket.TextMessage,[]byte(winner));err!=nil{
			t.Fatalf("could not send message over ws connection %v", err)
		}
		time.Sleep(10 * time.Millisecond)
		assertFinishCalledWith(t,game,winner)

	})
	t.Run("start game with 3 players and finish game with 'Numeez'as winner",func(t *testing.T){
		game := &GameSpy{}
		out:= &bytes.Buffer{}
		in:=userSends(3,"Numeez")
		cli:=NewCLI(in,out,game)
		cli.PlayPoker()
		assertMessagesSentToUser(t, out, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Numeez")
	
		

	})
	t.Run("start a game with 3 players and declare Ruth the winner", func(t *testing.T) {
		game := &GameSpy{}
		dummyPlayerStore:=&poker.StubPlayerStore{}
		winner := "Ruth"
		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")
	
		defer server.Close()
		defer ws.Close()
	
		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)
	
		time.Sleep(10 * time.Millisecond)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, winner)
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

func mustMakePlayerServer(t *testing.T, store PlayerStore, game *GameSpy) *PlayerServer {
	server, err := NewPlayerServer(store,game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}

func userSends(numberOfPlayer int , nameOfPlayer string) io.Reader{
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%d\n%s\n", numberOfPlayer, nameOfPlayer)
	ioReader := io.Reader(&buf)
	return ioReader
}

func assertMessagesSentToUser(t testing.TB,out *bytes.Buffer, message string){
	t.Helper()
	if out.String()!=message{
		t.Errorf("Actual :%s Expected %s ",out.String(),message)
	}
}

func assertGameStartedWith(t testing.TB,game *GameSpy,num int){
	t.Helper()
	if game.StartedWith!=num{
		t.Errorf("Actual : %d Expected : %d",game.StartedWith,num)
	}
}

func assertFinishCalledWith(t testing.TB,game *GameSpy,name string){
	t.Helper()

	if game.FinishedWith!=name{
		t.Errorf("Actual : %s Expected : %s",game.FinishedWith,name)
	}
}