package main

import (
	"net/http"

	s "github.com/dvl-numeez/go-with-tests/server"
)





func main(){
	handler:=http.HandlerFunc(s.PlayerServer)
	http.ListenAndServe(":3000",handler)

}