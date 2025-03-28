package blogrenderer_test

import (
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "Metroid Prime 4",
		Description: "A highly anticipated sequel for Nintendo Switch",
		Body:        "Metroid Prime 4 releases in 2025.",
		Tags:        []string{"metroid", "nintendo"},
	}

	t.Run("it converts a post to a single HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Metroid Prime 4</h1>`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
