package sort

import (
	"sort"
	"strings"

	"myls/internal/cli"
	"myls/internal/types"
)

func SortEntries(entries []types.FileEntry, flags cli.Flags) {
	sortStart := 0
	if !flags.TimeSort {
		sortStart = 2
	}

	sort.Slice(entries[sortStart:], func(i, j int) bool {
		if flags.TimeSort {
			return entries[sortStart+i].ModTime.After(entries[sortStart+j].ModTime)
		}

		a := entries[sortStart+i].Name
		b := entries[sortStart+j].Name

		la := strings.ToLower(a)
		lb := strings.ToLower(b)

		if la == lb {
			return a < b
		}
		return la < lb
	})
}
