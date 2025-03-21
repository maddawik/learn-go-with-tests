package main

import "errors"

var (
	ErrNotFound   = errors.New("this isn't the word you're looking for")
	ErrWordExists = errors.New("you know this word... think")
)

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) error {
	d[word] = definition
	return nil
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
