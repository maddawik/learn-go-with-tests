package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/maddawik/learn-go-with-tests/blogrenderer"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "Metroid Prime 4",
		Body:        "A highly anticipated sequel for Nintendo Switch",
		Description: "Metroid Prime 4 releases in 2025.",
		Tags:        []string{"metroid", "nintendo"},
	}

	t.Run("it converts a post to a single HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		renderer, err := blogrenderer.NewPostRenderer()
		if err != nil {
			return
		}

		if err := renderer.Render(&buf, aPost); err != nil {
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

	renderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for b.Loop() {
		renderer.Render(io.Discard, aPost)
	}
}
