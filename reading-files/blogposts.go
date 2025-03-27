package blogposts

import "testing/fstest"

type Post struct{}

func NewPostFromFS(filesystem fstest.MapFS) []Post {
	return nil
}
