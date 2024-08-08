package cli

import p "github.com/dvl-numeez/go-with-tests/cli/poker"



func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}


func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []p.Player {
	var league []p.Player
	for name, wins := range i.store {
		league = append(league, p.Player{name, wins})
	}
	return league
}
