package main

import (
	"fmt"
	"os"

	"myls/internal/cli"
	"myls/internal/filesystem"
	"myls/internal/output"
)

func main() {
	flags, path := cli.ParseFlags(os.Args[1:])

	if flags.Recursive {
		output.PrintRecursive(path, flags)
		return
	}

	entries, err := filesystem.ListDirectory(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	entries = cli.FilterHidden(flags, entries)
	if flags.All {
		entries = cli.AddSpecialEntries(path, entries)
	}

	start := cli.SpecialStart(flags)
	cli.SortEntries(flags, entries, start)
	if flags.Reverse {
		cli.ReverseEntries(entries, start)
	}

	if flags.Long {
		output.PrintLong(entries)
	} else {
		output.PrintSimple(entries)
	}
}