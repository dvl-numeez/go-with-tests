package websockets

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/dvl-numeez/go-with-tests/websockets/poker"
)

type TexasGame struct {
	Alerter poker.BlindAlerter
	Store   poker.PlayerStore
}


func(g *TexasGame)Start(numberOfPlayers int, alertDestination io.Writer){
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Second

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.Alerter.ScheduleAlertAt(blindTime, blind,alertDestination)
		blindTime = blindTime + blindIncrement
	}
}

func (g *TexasGame) Finish(winner string) {
	g.Store.RecordWin(winner)
}
func NewGame(alerter poker.BlindAlerter, store poker.PlayerStore) *TexasGame {
	return &TexasGame{
		Alerter: alerter,
		Store:   store,
	}
}

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

const PlayerPrompt = "Please enter the number of players: "

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))
	if err!=nil{
		fmt.Fprint(cli.out,poker.BadInputPrompt)
	}

	cli.game.Start(numberOfPlayers,cli.out)

	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)

	cli.game.Finish(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins\n", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
