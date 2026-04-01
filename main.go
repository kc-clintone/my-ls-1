package main

import (
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	path := "."
	flags := map[rune]bool{
		'l': false,
		'a': false,
		'r': false,
		't': false,
		'R': false,
	}
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			for _, ch := range arg {
				flags[ch] = true
			}
		} else {
			path = arg
		}
	}
	ListDirectory(path, flags)
}
