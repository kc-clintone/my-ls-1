package main

import (
	"io/fs"
	"os"
)

func ListDirectory(path string) ([]fs.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return entries, nil
}