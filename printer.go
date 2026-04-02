package main

import (
	"fmt"
	"os"
)

func PrintEntries(entries []FileEntry, flags map[rune]bool) {
	if flags['l'] {
		PrintLong(entries)
	} else {
		PrintSimple(entries)
	}
}

func PrintSimple(entries []FileEntry) {
	for _, e := range entries {
		fmt.Println(e.Name)
	}
}

func FormatPermissions(mode os.FileMode) string {
	return mode.String()
}