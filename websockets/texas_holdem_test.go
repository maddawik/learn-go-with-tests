package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/maddawik/learn-go-with-tests/websockets"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Play(5)

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

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("schedules alerts for game start on 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		dummyPlayerStore := &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Play(7)

		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		checkSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := strings.NewReader("Berries\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		poker.AssertGameNotStarted(t, game.StartCalled)
		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func TestGame_Finish(t *testing.T) {
	t.Run("saves a win by Kayla for valid input", func(t *testing.T) {
		dummyBlindAlerter := &poker.SpyBlindAlerter{}
		store := &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(dummyBlindAlerter, store)
		winner := "Kayla"
		game.Finish(winner)

		poker.AssertPlayerWin(t, store, winner)
	})

	t.Run("it throws an error for not sending exactly 2 words in input", func(t *testing.T) {
		// 1 token in input
		out := &bytes.Buffer{}
		in := strings.NewReader("5\nBananas\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)

		// 3 tokens in input
		out.Reset()
		in = strings.NewReader("5\nApples and Grapes\n")
		cli = poker.NewCLI(in, out, game)
		cli.PlayPoker()

		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
	})

	t.Run("it throws an error when input is not in the form `playername wins`", func(t *testing.T) {
		out := &bytes.Buffer{}
		in := strings.NewReader("5\nBananas loses\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
	})
}

func checkSchedulingCases(t *testing.T, cases []poker.ScheduledAlert, blindAlerter *poker.SpyBlindAlerter) {
	t.Helper()

	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()

	amountGot := got.Amount
	if amountGot != want.Amount {
		t.Errorf("got amount %d, want %d", amountGot, want.Amount)
	}

	gotScheduledTime := got.At
	if gotScheduledTime != want.At {
		t.Errorf("got time %v, want %v", gotScheduledTime, want.At)
	}
}
