package blogrenderer_test

import (
	"bytes"
	"testing"

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
		err := blogrenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Metroid Prime 4</h1><p>Metroid Prime 4 releases in 2025.</p>Tags: <ul><li>metroid</li><li>nintendo</li></ul>`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
