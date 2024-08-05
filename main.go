package main

import (
	"log"
	"net/http"

	di "github.com/dvl-numeez/go-with-tests/DI"
)


func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}


func main(){
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}