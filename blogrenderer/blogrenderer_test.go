package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/maddawik/learn-go-with-tests/blogrenderer"
)

func TestRender(t *testing.T) {
	aBody := `# Metroid Prime 4: Beyond

This highly anticipated sequel for Nintendo Switch is coming out this year!

## Gameplay

Samus has psychic powers she can use to unravel the mysteries of a new planet!

She'll have to:

- Explore
- Solve Puzzles
- Fight Enemies

And more!

Find out more at [Nintendo](https://nintendo.com)`

	aPost := blogrenderer.Post{
		Title:       "Metroid Prime 4",
		Body:        aBody,
		Description: "Metroid Prime 4 releases in 2025.",
		Tags:        []string{"metroid", "nintendo"},
	}

	t.Run("it converts a post to a single HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		postRenderer, err := blogrenderer.NewPostRenderer()
		if err != nil {
			return
		}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

// FIX: `go test -bench=.` does not run this test?
func BencharkRender(b *testing.B) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Description: "the description",
		Body:        "the body",
		Tags:        []string{"go", "tdd"},
	}

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for b.Loop() {
		postRenderer.Render(io.Discard, aPost)
	}
}
