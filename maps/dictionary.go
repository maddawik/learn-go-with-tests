package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("this isn't the word you're looking for")
	}
	return definition, nil
}
