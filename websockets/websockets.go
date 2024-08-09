package websockets

import "net/http"


type PlayerStore interface{
	GetPlayerScore(name string)int
	RecordWin(name string)
}
type PlayerServer struct {
	store PlayerStore
	http.Handler
}
type StubPlayerStore struct{

}

func(s *StubPlayerStore)GetPlayerScore(name string)int{
	return 0	
}
func(s *StubPlayerStore)RecordWin(name string){

}

const jsonContentType = "application/json"

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game",http.HandlerFunc(p.game))

	p.Handler = router

	return p
}


func(p *PlayerServer)leagueHandler(w http.ResponseWriter,r *http.Request){

}


func(p *PlayerServer)playersHandler(w http.ResponseWriter,r *http.Request){
	
}

func(p *PlayerServer)game(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusOK)
}