package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/maddawik/learn-go-with-tests/reading-files"
)

func TestBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hey")},
		"hello-world2.md": {Data: []byte("bonjour")},
	}

	posts, err := blogposts.NewPostFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d, wanted %d posts", len(posts), len(fs))
	}
}
