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
	name:=strings.TrimPrefix(r.URL.Path,"/players/")
	fmt.Fprint(w,s.Store.GetPlayerScore(name))
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