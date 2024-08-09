package time

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/dvl-numeez/go-with-tests/time/poker"
)

func TestTime(t *testing.T) {

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("5\n")
		playerStore := &poker.StubPlayerStore{}
		var dummyStdOut = &bytes.Buffer{}
		blindAlerter := &SpyBlindAlerter{}
		game := NewGame(blindAlerter, playerStore)
		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}
				got := blindAlerter.alerts[i]
				assertScheduleAlert(t, got, want)
			})
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		var dummyPlayerStore = &poker.StubPlayerStore{}
		stdout := &bytes.Buffer{}
		blindAlerter := &SpyBlindAlerter{}
		input := strings.NewReader("7\n")
		game := NewGame(blindAlerter, dummyPlayerStore)
		cli := NewCLI(input, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := "Please enter the number of players: "

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduleAlert(t, got, want)
			})
		}

	})

}

func TestGame_Start(t *testing.T) {
	var dummyPlayerStore = &poker.StubPlayerStore{}

	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {

		blindAlerter := &poker.BlindAlert{}
		game := NewGame(blindAlerter, dummyPlayerStore)

		game.Start(5)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.BlindAlert{}
		game := NewGame(blindAlerter, dummyPlayerStore)

		game.Start(7)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()
		wantPrompt := poker.PlayerPrompt + poker.BadInputPrompt
		assertMessagesSentToUser(t, stdout, wantPrompt)

	})

}

func TestGame_Finish(t *testing.T) {
	var dummyBlindAlerter = &SpyBlindAlerter{}
	store := &poker.StubPlayerStore{}
	game := NewGame(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	AssertWin(t, store, winner)
}

func assertScheduleAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("Actual : %+v Expected : %+v", got, want)
	}
}

func checkSchedulingCases(t testing.TB, cases []poker.ScheduledAlert, alerter *poker.BlindAlert) {
	t.Helper()
	for i, want := range cases {
		if len(alerter.Alerts) <= i {
			t.Fatalf("alert %d was not scheduled %v", i, alerter.Alerts)
		}

		got := alerter.Alerts[i]
		assertScheduleAlert(t, got, want)

	}
}

func AssertWin(t testing.TB, store *poker.StubPlayerStore, winner string) {

	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store the correct winner got %q want %q", store.WinCalls[0], winner)
	}

}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
