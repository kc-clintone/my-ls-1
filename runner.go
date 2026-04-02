package main

import "fmt"

func Run(flags map[rune]bool, paths []string) {
	for i, path := range paths {

		// TODO: add recursive flag handling here

		entries, err := ListDirectory(path, flags)
		if err != nil {
			PrintError(err)
			continue
		}

		entries = FilterEntries(entries, flags)
		SortEntries(entries, flags)

		PrintHeader(path, len(paths) > 1)
		PrintEntries(entries, flags)

		if i < len(paths)-1 {
			fmt.Println()
		}
	}
}