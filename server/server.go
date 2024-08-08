package server

import (
	"fmt"
	"net/http"
	"strings"
)


func PlayerServer(w http.ResponseWriter,r *http.Request){
	name:=strings.TrimPrefix(r.URL.Path,"/players/")
	fmt.Fprint(w,GetPlayerScore(name))
	
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