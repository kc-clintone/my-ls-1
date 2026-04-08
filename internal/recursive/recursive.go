package recursive

import (
	"fmt"
	"os"
	"strings"

	"myls/internal/cli"
	"myls/internal/filesystem"
	"myls/internal/output"
	"myls/internal/types"
	"myls/internal/sort"
)

func ListRecursive(path string, flags cli.Flags) {
	info, err := os.Lstat(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if !info.IsDir() {
		entry, err := filesystem.SingleEntry(path)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if flags.Long {
			output.PrintLong([]types.FileEntry{entry})
		} else {
			output.PrintSimple([]types.FileEntry{entry})
		}
		return
	}

	entries, err := filesystem.ListDirectory(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// filter hidden
	if !flags.All {
		var filtered []types.FileEntry
		for _, e := range entries {
			if !strings.HasPrefix(e.Name, ".") {
				filtered = append(filtered, e)
			}
		}
		entries = filtered
	}

	// add . and ..
	if flags.All {
		dot := filesystem.CreateSpecialEntry(path, ".")
		dotdot := filesystem.CreateSpecialEntry(path, "..")
		entries = append([]types.FileEntry{dot, dotdot}, entries...)
	}

	sort.SortEntries(entries, flags)

	// print header
	fmt.Println(path + ":")

	if flags.Long {
		output.PrintLong(entries)
	} else {
		output.PrintSimple(entries)
	}

	fmt.Println()

	// recurse
	for _, e := range entries {
		if e.IsDir && e.Name != "." && e.Name != ".." {
			sub := path + "/" + e.Name
			ListRecursive(sub, flags)
		}
	}
}