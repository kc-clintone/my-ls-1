package main

import (
	"errors"
	"strings"
)

func ParseArgs(args []string) (map[rune]bool, []string, error) {
	flags := map[rune]bool{
		'l': false,
		'a': false,
		'r': false,
		't': false,
		'R': false,
	}

	var paths []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			for _, ch := range arg[1:] {
				if _, ok := flags[ch]; !ok {
					return nil, nil, errors.New("invalid flag: " + string(ch))
				}
				flags[ch] = true
			}
		} else {
			paths = append(paths, arg)
		}
	}

	if len(paths) == 0 {
		paths = []string{"."}
	}

	return flags, paths, nil
}