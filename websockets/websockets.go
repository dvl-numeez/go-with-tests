package websockets

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/websocket"
)


type PlayerStore interface{
	GetPlayerScore(name string)int
	RecordWin(name string)
}
type PlayerServer struct {
	store PlayerStore
	http.Handler
}
type StubPlayerStore struct{
	WinCalls []string
}

func(s *StubPlayerStore)GetPlayerScore(name string)int{
	return 0	
}
func(s *StubPlayerStore)RecordWin(name string){
	s.WinCalls = append(s.WinCalls, name)
}

const jsonContentType = "application/json"

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game",http.HandlerFunc(p.game))
	router.Handle("/ws",http.HandlerFunc(p.webSocket))

	p.Handler = router

	return p
}


func(p *PlayerServer)leagueHandler(w http.ResponseWriter,r *http.Request){

}


func(p *PlayerServer)playersHandler(w http.ResponseWriter,r *http.Request){
	
}

func(p *PlayerServer)game(w http.ResponseWriter,r *http.Request){
	tmpl,err:=template.ParseFiles("game.html")
	if err!=nil{
		http.Error(w, fmt.Sprintf("problem loading template %s", err.Error()), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w,tmpl)
}

func(p *PlayerServer)webSocket(w http.ResponseWriter, r *http.Request){
	upgrader:= websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
	conn,_:=upgrader.Upgrade(w,r,nil)
	_,winnerMsg,_:=conn.ReadMessage()
	p.store.RecordWin(string(winnerMsg))

}