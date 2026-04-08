package output

import (
	"fmt"
	"path/filepath"

	"myls/internal/cli"
	"myls/internal/filesystem"
)

func PrintRecursive(root string, flags cli.Flags) {
	var walk func(dir string)
	walk = func(dir string) {
		entries, err := filesystem.ListDirectory(dir)
		if err != nil {
			println("Error:", err.Error())
			return
		}

		entries = cli.FilterHidden(flags, entries)
		if flags.All {
			entries = cli.AddSpecialEntries(dir, entries)
		}

		start := cli.SpecialStart(flags)
		cli.SortEntries(flags, entries, start)
		if flags.Reverse {
			cli.ReverseEntries(entries, start)
		}

		// header and output
		println(dir + ":")
		if flags.Long {
			PrintLong(entries)
			fmt.Println()
		} else {
			PrintSimple(entries)
			fmt.Println()
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
				walk(child)
			}
		}
	}

	walk(root)
}
