package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	poker "github.com/maddawik/learn-go-with-tests/websockets"
)

func TestCLI(t *testing.T) {
	dummyStdOut := &bytes.Buffer{}

	t.Run("record Cody win from user input", func(t *testing.T) {
		playerName := "Cody"
		in := strings.NewReader(fmt.Sprintf("5\n%s wins\n", playerName))

		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		poker.AssertGameStartedWith(t, 5, game.StartedWith)
		poker.AssertGameFinishedWith(t, playerName, game.FinishedWith)
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		playerName := "May"
		in := strings.NewReader(fmt.Sprintf("7\n%s wins\n", playerName))

		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		want := poker.PlayerPrompt

		poker.AssertMessagesSentToUser(t, stdout, want)
		poker.AssertGameStartedWith(t, 7, game.StartedWith)
	})
}
