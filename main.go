package main

import (
	"net/http"
	"log"
	s "github.com/dvl-numeez/go-with-tests/server"
)


type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}


func main(){
	server := s.PlayerServer{Store:&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", &server))
}