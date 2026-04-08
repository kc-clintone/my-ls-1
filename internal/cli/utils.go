package cli

import (
	"sort"
	"strings"

	"myls/internal/filesystem"
	"myls/internal/types"
)

// helper: filter out hidden files unless -a
func FilterHidden(flags Flags, entries []types.FileEntry) []types.FileEntry {
	if flags.All {
		return entries
	}
	var filtered []types.FileEntry
	for _, e := range entries {
		if !strings.HasPrefix(e.Name, ".") {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

// helper: add "." and ".." at front when -a
func AddSpecialEntries(dir string, entries []types.FileEntry) []types.FileEntry {
	dot := filesystem.CreateSpecialEntry(dir, ".")
	dotdot := filesystem.CreateSpecialEntry(dir, "..")
	return append([]types.FileEntry{dot, dotdot}, entries...)
}

func SpecialStart(flags Flags) int {
	if flags.All {
		return 2
	}
	return 0
}

func SortEntries(flags Flags, entries []types.FileEntry, start int) {
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

func ReverseEntries(entries []types.FileEntry, start int) {
	for i, j := start, len(entries)-1; i < j; i, j = i+1, j-1 {
		entries[i], entries[j] = entries[j], entries[i]
	}
}
