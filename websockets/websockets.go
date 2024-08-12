package websockets

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

const templatePath = "game.html"
type PlayerStore interface{
	GetPlayerScore(name string)int
	RecordWin(name string)
}
type PlayerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
	playGame Game
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

func NewPlayerServer(store PlayerStore,game Game) (*PlayerServer,error) {
	p := new(PlayerServer)
	tmpl,err:= template.ParseFiles(templatePath)
	if err !=nil{
		return nil, fmt.Errorf("problem opening %s %v", templatePath, err)
	}
	p.store = store
	p.template = tmpl
	p.playGame = game

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game",http.HandlerFunc(p.game))
	router.Handle("/ws",http.HandlerFunc(p.webSocket))

	p.Handler = router

	return p,nil
}


func(p *PlayerServer)leagueHandler(w http.ResponseWriter,r *http.Request){

}


func(p *PlayerServer)playersHandler(w http.ResponseWriter,r *http.Request){
	
}

func(p *PlayerServer)game(w http.ResponseWriter,r *http.Request){
	p.template.Execute(w,nil)
}
var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func(p *PlayerServer)webSocket(w http.ResponseWriter, r *http.Request){
	conn,_:=wsUpgrader.Upgrade(w,r,nil)
	_, numberOfPlayersMsg, _ := conn.ReadMessage()
	numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayersMsg))
	p.playGame.Start(numberOfPlayers,io.Discard)
	_,winnerMsg,_:=conn.ReadMessage()
	p.playGame.Finish(string(winnerMsg))

}