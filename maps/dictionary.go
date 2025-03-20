package main

import "errors"

var ErrNotFound = errors.New("this isn't the word you're looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
