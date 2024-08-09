package poker

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct{
	playerStore PlayerStore
	in *bufio.Scanner
	out io.Writer
	alerter BlindAlerter
}
type BlindAlerterFunc func(duration time.Duration, amount int)

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
type BlindAlert struct {
	Alerts []ScheduledAlert
}
func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

func (s *BlindAlert) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts,ScheduledAlert{duration,amount})
}

func(cli *CLI)PlayPoker(){
	fmt.Fprint(cli.out,PlayerPrompt)
	cli.scheduleBlindAlerts()
	input := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(input))
}
func(cli *CLI)scheduleBlindAlerts(){
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _,blind:=range blinds{
		cli.alerter.ScheduleAlertAt(blindTime,blind)
		blindTime = blindTime+10*time.Minute
	}
}
func NewCLI(store PlayerStore, in io.Reader,out io.Writer ,alerter BlindAlerter ) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
		out:out,
		alerter: alerter,

	}
}
func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
