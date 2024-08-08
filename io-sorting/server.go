package iosorting

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

)

const jsonContentType = "application/json"

type Player struct{
	Name string
	Wins int
}

type PlayerStore interface{
	GetPlayerScore(name string)int
	RecordWin(name string)
	GetLeague() League
}
type PlayerServer struct{
	Store PlayerStore
	http.Handler
}

type StubPlayerStore struct{
	scores   map[string]int
	winCalls []string
	league   []Player
}

type FileSystemPlayerStore struct{
	Database io.ReadWriteSeeker
	
}

func(s *FileSystemPlayerStore)GetLeague() League{
	s.Database.Seek(0,io.SeekStart)
	player,_:=NewLeague(s.Database)
	return player
}

func(s *FileSystemPlayerStore)GetPlayerScore(name string)int{
	league:=s.GetLeague().Find(name)
	if league != nil {
		return league.Wins
	}

	return 0
}
func(s *FileSystemPlayerStore)RecordWin(name string){
	league:= s.GetLeague()
	player:=league.Find(name)
	if player != nil {
		player.Wins++
	}else{
		league = append(league, Player{name, 1})
	}
	s.Database.Seek(0, io.SeekStart)
	json.NewEncoder(s.Database).Encode(league)

}

func(store *StubPlayerStore)GetLeague()[]Player{
	return store.league
}

func(store *StubPlayerStore)GetPlayerScore(name string)int{
	return store.scores[name]
}
func (store *StubPlayerStore)RecordWin(name string){
	store.winCalls = append(store.winCalls, name)
}

func NewPlayerServer(store PlayerStore)*PlayerServer{
	p := new(PlayerServer)

	p.Store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
	
}

func (s *PlayerServer)ServeHTTP(w http.ResponseWriter, r *http.Request){
	
	s.Handler.ServeHTTP(w,r)
	
	
}

func (p *PlayerServer)leagueHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.Store.GetLeague())
	w.WriteHeader(http.StatusOK)
}
func (p *PlayerServer)playersHandler(w http.ResponseWriter, r *http.Request){
	method:=r.Method
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch method{
	case http.MethodGet:
		p.showScore(w,player)
	case http.MethodPost:
		p.processWin(w,player)
		

	}
}
func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter,player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
func (p *PlayerServer)getLeagueTable()[]Player{
	return []Player{
		{"Chris", 20},
	}
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}

	return league, err
}