package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/maddawik/learn-go-with-tests/reading-files"
)

func TestBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: go, tdd
---
Hello
World`

		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
Hola
Enrique`
	)

	t.Run("fs with 2 valid files", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		assertPosts(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"go", "tdd"},
			Body: `Hello
World`,
		})
		assertPosts(t, posts[1], blogposts.Post{
			Title:       "Post 2",
			Description: "Description 2",
			Tags:        []string{"rust", "borrow-checker"},
			Body: `Hola
Enrique`,
		})
	})

	t.Run("fs with no files returns error", func(t *testing.T) {
		fs := fstest.MapFS{}

		_, err := blogposts.NewPostsFromFS(fs)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

	t.Run("fs with no markdown files", func(t *testing.T) {
		fs := fstest.MapFS{
			"txt-is-not-md.txt": {Data: []byte("hello")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if len(posts) != 0 {
			t.Errorf("got %d posts, expected none", len(posts))
		}

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

	t.Run("fs with some markdown files", func(t *testing.T) {
		file := `Title: GNUs Not UNIX
Description: Seriously
Tags: gnu, linux
---
Bonjour
GNU`

		fs := fstest.MapFS{
			"txt-is-not-md.txt": {Data: []byte("hello")},
			"gnus-not-unix.md":  {Data: []byte(file)},
		}

		got, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		if len(got) != 1 {
			t.Errorf("got %d posts, expected ", len(got))
		}
	})
}

func assertPosts(t testing.TB, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}
