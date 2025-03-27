package blogposts

import (
	"bufio"
	"io"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

type Post struct {
	Title       string
	Description string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}
	titleLine := readLine()[len(titleSeparator):]
	descriptionLine := readLine()[len(descriptionSeparator):]

	return Post{Title: titleLine, Description: descriptionLine}, nil
}
