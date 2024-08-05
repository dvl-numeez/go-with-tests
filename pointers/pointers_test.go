package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertError:=func(t testing.TB,got ,want error){
		t.Helper()
		if got==nil{
			t.Fatal("wanted an error but didn't get one")
		}
		if got!=want{
			t.Errorf("got %q, want %q", got, want)
		}

	}
	assertBalance:= func(t testing.TB,wallet Wallet,want Bitcoin){
		t.Helper()
		got:= wallet.Balance()
		if got != want {
			t.Errorf("Got : %s Wanted : %s", got, want)
		}

	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		amount := Bitcoin(10)
		wallet.Deposit(amount)
		want := Bitcoin(10)
		assertBalance(t,wallet,want)

		
	})

	t.Run("Withdarw",func(t *testing.T){
		wallet:=Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		want:=Bitcoin(10)
		assertBalance(t,wallet,want)

	})

	t.Run("Insufficient Funds",func(t *testing.T){
		startingBalance:=Bitcoin(20)
		wallet:=Wallet{balance: startingBalance}
		err:=wallet.Withdraw(Bitcoin(100))
		assertError(t,err,errInsufficientFunds)
		assertBalance(t,wallet,startingBalance)
		

	})

}
