package time


type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartedCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
	g.StartedCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}