package main

import "io"

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() (league []Player) {
	return league
}
