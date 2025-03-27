package blogposts

import (
	"io/fs"
)

type Post struct{}

func NewPostFromFS(filesystem fs.FS) []Post {
	dir, _ := fs.ReadDir(filesystem, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts
}
