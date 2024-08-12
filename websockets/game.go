package websockets

import "io"


type Game interface {
	Start(numberOfPlayers int, alertDestination io.Writer)
	Finish(winner string)
}

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartedCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int,alertDestination io.Writer) {
	g.StartedWith = numberOfPlayers
	g.StartedCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}