package poker

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

)

const PlayerPrompt = "Please enter the number of players: "
const BadInputPrompt = "Enter a number please ........."
type CLI struct{
	playerStore PlayerStore
	in *bufio.Scanner
	out io.Writer
	alerter BlindAlerter
	
}
type BlindAlerterFunc func(duration time.Duration, amount int,to io.Writer)

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int,to io.Writer)
}
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int,to io.Writer) {
	a(duration, amount,to)
}
func Alerter(duration time.Duration, amount int) {
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
	
	numOfPlayers, _ := strconv.Atoi(cli.readLine())

	cli.scheduleBlindAlerts(numOfPlayers)
	input := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(input))
}
func(cli *CLI)scheduleBlindAlerts(num int){
	blindIncrement := time.Duration(5+num) * time.Minute
	blindTime := 0 * time.Second
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	for _,blind:=range blinds{
		cli.alerter.ScheduleAlertAt(blindTime,blind,cli.out)
		blindTime = blindTime+blindIncrement
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
