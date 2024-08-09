package main

import (
	"fmt"
	"log"
	"os"

	t "github.com/dvl-numeez/go-with-tests/time"
	"github.com/dvl-numeez/go-with-tests/time/poker"
)

const dbFileName = "game.db.json"

func main() {

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	defer close()
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := t.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	cli := t.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
