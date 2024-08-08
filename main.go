package main

import (
	"log"
	"net/http"

	s "github.com/dvl-numeez/go-with-tests/server"
)






func main(){
	server := &s.PlayerServer{Store : s.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}