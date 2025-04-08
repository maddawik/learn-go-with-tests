package poker_test

import (
	"bytes"
	"strings"
	"testing"

	poker "github.com/maddawik/learn-go-with-tests/time"
)

func TestCLI(t *testing.T) {
	dummyStdOut := &bytes.Buffer{}

	t.Run("record May win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nMay wins\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.StartedWith != 5 {
			t.Errorf("wanted Play called with 5 but got %d", game.StartedWith)
		}

		if game.FinishedWith != "May" {
			t.Errorf("wanted winner to be May but got %q", game.FinishedWith)
		}
	})

	t.Run("record Cody win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nCody wins\n")

		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.StartedWith != 5 {
			t.Errorf("wanted Play called with 5 but got %d", game.StartedWith)
		}

		if game.FinishedWith != "Cody" {
			t.Errorf("wanted winner to be Cody but got %q", game.FinishedWith)
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")

		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Play called with 7 but got %d", game.StartedWith)
		}
	})
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
