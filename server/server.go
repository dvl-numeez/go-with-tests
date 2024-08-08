package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface{
	GetPlayerScore(name string)int
}
type PlayerServer struct{
	Store PlayerStore
}

type StubPlayerStore struct{
	scores map[string]int
}

func(store *StubPlayerStore)GetPlayerScore(name string)int{
	return store.scores[name]
}

func (s *PlayerServer)ServeHTTP(w http.ResponseWriter, r *http.Request){
	method:=r.Method
	switch method{
	case http.MethodGet:
		name:=strings.TrimPrefix(r.URL.Path,"/players/")
		score:=s.Store.GetPlayerScore(name)
		if score==0{
		w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprint(w,score)
	case http.MethodPost:
		w.WriteHeader(http.StatusAccepted)

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