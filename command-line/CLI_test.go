package poker_test

import (
	"strings"
	"testing"

	poker "github.com/maddawik/learn-go-with-tests/command-line"
)

func TestCLI(t *testing.T) {
	t.Run("record May win from user input", func(t *testing.T) {
		in := strings.NewReader("May wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "May")
	})

	t.Run("record Cody win from user input", func(t *testing.T) {
		in := strings.NewReader("Cody wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cody")
	})
}
