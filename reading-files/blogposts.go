package blogposts

import (
	"errors"
	"io/fs"
	"strings"
)

var ErrNoMarkdownFiles = errors.New("no markdown files in directory")

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	mdFiles, err := getMarkdownFiles(dir)
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range mdFiles {
		post, err := getPost(filesystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(filesystem fs.FS, filename string) (Post, error) {
	postFile, err := filesystem.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func getMarkdownFiles(dir []fs.DirEntry) ([]fs.DirEntry, error) {
	result := make([]fs.DirEntry, 0)
	for _, f := range dir {
		if strings.HasSuffix(f.Name(), ".md") {
			result = append(result, f)
		}
	}

	if len(result) == 0 {
		return nil, ErrNoMarkdownFiles
	}

	return result, nil
}
