package websockets

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

const templatePath = "game.html"
type playerServerWs struct{
	*websocket.Conn
}
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

func newPlayerServerWs(w http.ResponseWriter ,r *http.Request) *playerServerWs{
	conn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("problem upgrading connection to WebSockets %v\n", err)
	}

	return &playerServerWs{conn}
}
func (w *playerServerWs) WaitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("error reading from websocket %v\n", err)
	}
	return string(msg)
}
func (w *playerServerWs) Write(p []byte) (n int, err error) {
	err = w.WriteMessage(websocket.TextMessage, p)

	if err != nil {
		return 0, err
	}

	return len(p), nil
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
	ws:=newPlayerServerWs(w,r)
	numberOfPlayersMsg:= ws.WaitForMsg()
	numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayersMsg))
	p.playGame.Start(numberOfPlayers,ws)
	winnerMsg:=ws.WaitForMsg()
	p.playGame.Finish(winnerMsg)

}