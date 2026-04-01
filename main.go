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
			for _, ch := range arg[1:] {
				flags[ch] = true
			}
		} else {
			path = arg
		}
	}
	entries, err := ListDirectory(path, flags)
	if err != nil {
		println("Error:", err.Error())
		return
	}

	entries = FilterEntries(entries, flags)

	for _, e := range entries {
		println(e.Name)
	}
}
