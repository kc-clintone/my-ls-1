package main

import (
	"os"
	"sort"
	"strings"

	"myls/internal/cli"
	"myls/internal/filesystem"
	"myls/internal/output"
	"myls/internal/types"
)

func main() {
	// Parse command-line arguments
	flags, path := cli.ParseFlags(os.Args[1:])

	// Read directory
	entries, err := filesystem.ListDirectory(path)
	if err != nil {
		println("Error:", err.Error())
		return
	}

	// Filter hidden files
	if !flags.All {
		var filtered []types.FileEntry
		for _, e := range entries {
			if !strings.HasPrefix(e.Name, ".") {
				filtered = append(filtered, e)
			}
		}
		entries = filtered
	}

	// Add . and ..
	if flags.All {
		dot := filesystem.CreateSpecialEntry(path, ".")
		dotdot := filesystem.CreateSpecialEntry(path, "..")
		entries = append([]types.FileEntry{dot, dotdot}, entries...)
	}

	// Sort (skip . and .. if present)
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

	// Reverse if requested
	if flags.Reverse {
		for i, j := start, len(entries)-1; i < j; i, j = i+1, j-1 {
			entries[i], entries[j] = entries[j], entries[i]
		}
	}

	// Output
	if flags.Long {
		output.PrintLong(entries)
	} else {
		output.PrintSimple(entries)
	}
}
