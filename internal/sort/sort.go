package sort

import (
	"sort"
	"strings"

	"myls/internal/cli"
	"myls/internal/types"
)

func SortEntries(entries []types.FileEntry, flags cli.Flags) {
	start := 0
	if flags.All {
		start = 2
	}

	sort.Slice(entries[start:], func(i, j int) bool {
		if flags.TimeSort {
			return entries[start+i].ModTime.After(entries[start+j].ModTime)
		}

		a := entries[start+i].Name
		b := entries[start+j].Name

		la := strings.ToLower(a)
		lb := strings.ToLower(b)

		if la == lb {
			return a < b
		}
		return la < lb
	})
}

