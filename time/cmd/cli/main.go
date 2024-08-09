package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dvl-numeez/go-with-tests/time/poker"
)

const dbFileName = "game.db.json"

func main(){
	

	store,close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)


	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	defer close()
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store,os.Stdin,os.Stdout,poker.BlindAlerterFunc(poker.StdOutAlerter)).PlayPoker()
}