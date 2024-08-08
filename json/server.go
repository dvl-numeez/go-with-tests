package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface{
	GetPlayerScore(name string)int
	RecordWin(name string)
}
type PlayerServer struct{
	Store PlayerStore
}

type StubPlayerStore struct{
	scores map[string]int
	winCalls []string
}

func(store *StubPlayerStore)GetPlayerScore(name string)int{
	return store.scores[name]
}
func (store *StubPlayerStore)RecordWin(name string){
	store.winCalls = append(store.winCalls, name)
}

func (s *PlayerServer)ServeHTTP(w http.ResponseWriter, r *http.Request){
	router:=http.NewServeMux()
	router.Handle("/league",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	}))
	router.Handle("/players/",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		method:=r.Method
		player := strings.TrimPrefix(r.URL.Path, "/players/")
		switch method{
		case http.MethodGet:
			s.showScore(w,player)
		case http.MethodPost:
			s.processWin(w,player)
			
	
		}
	}))

	router.ServeHTTP(w,r)
	
	
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

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}