package pointers

import (
	"fmt"
	"testing"
)


func TestWall(t *testing.T){
	wallet:=Wallet{}
	wallet.Deposit(10)
	got:=wallet.Balance()
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)
	want:=10
	if got!=want{
		t.Errorf("Got : %d Wanted : %d",got,want)
	}
}