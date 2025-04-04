package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Jeff wins\n")
	playerStore := &StubPlayerStore{}

	cli := &CLI{playerStore, in}
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Error("expected a win call but didn't get any")
	}

	got := playerStore.winCalls[0]
	want := "Jeff"

	if got != want {
		t.Errorf("didn't record correct winner, got %q want %q", got, want)
	}
}
