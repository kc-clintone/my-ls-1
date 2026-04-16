package recursive

import (
	"fmt"
	"path/filepath"
	"strings"

	"myls/internal/cli"
	"myls/internal/filesystem"
	"myls/internal/output"
)

func ListRecursive(root string, flags cli.Flags) {
	first := true
	var walk func(dir string)
	walk = func(dir string) {
		entries, err := filesystem.ListDirectory(dir)
		if err != nil {
			fmt.Printf("ls: cannot open directory '%s': %v\n", dir, err)
			return
		}

		entries = cli.FilterHidden(flags, entries)
		if flags.All {
			entries = cli.AddSpecialEntries(dir, entries)
		}

		start := cli.SpecialStart(flags)
		cli.SortEntries(flags, entries, start)
		if flags.Reverse {
			cli.ReverseEntries(flags, entries, start)
		}

		// header and output
		if !first {
			fmt.Println()
		}
		first = false
		fmt.Printf("%s:\n", dir)
		if flags.Long {
			output.PrintLong(entries)
		} else {
			output.PrintSimple(entries)
		}

		// recurse into subdirectories
		for _, e := range entries {
			if e.Name == "." || e.Name == ".." {
				continue
			}
			isDir := e.IsDir
			// if types.FileEntry has no IsDir, uncomment the next line and remove the e.IsDir line above:
			// isDir = filesystem.IsDir(filepath.Join(dir, e.Name))
			if isDir {
				child := filepath.Join(dir, e.Name)
				if dir == "." || strings.HasPrefix(dir, "./") {
					child = "./" + strings.TrimPrefix(child, "./")
				}
				walk(child)
			}
		}
	}

	walk(root)
}
