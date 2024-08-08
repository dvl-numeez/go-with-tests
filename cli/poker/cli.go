package poker

import (
	"bufio"
	"io"
	"strings"
)


type CLI struct{
	playerStore PlayerStore
	in *bufio.Scanner
}

func(cli *CLI)PlayPoker(){
	reader:=cli.in
	reader.Scan()
	name:=extractWinner(reader.Text())
	cli.playerStore.RecordWin(name)
}
func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
