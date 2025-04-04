package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record May win from user input", func(t *testing.T) {
		in := strings.NewReader("May wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "May")
	})

	t.Run("record Cody win from user input", func(t *testing.T) {
		in := strings.NewReader("Cody wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Cody")
	})
}
