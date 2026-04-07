package main

import (
	"os"

	"myls/internal/cli"
	"myls/internal/filesystem"
	"myls/internal/output"
)

func main() {
	// Parse command-line arguments
	flags, path := cli.ParseFlags(os.Args[1:])

	if flags.Recursive {
		output.PrintRecursive(path, flags)
		return
	}

	entries, err := filesystem.ListDirectory(path)
	if err != nil {
		println("Error:", err.Error())
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
