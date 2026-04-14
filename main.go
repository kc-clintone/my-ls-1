package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	"myls/internal/cli"
	"myls/internal/filesystem"
	"myls/internal/recursive"
	"myls/internal/output"
	"myls/internal/types"
)

func main() {
	flags, targets := cli.ParseFlags(os.Args[1:])

	files := make([]types.FileEntry, 0)
	dirs := make([]string, 0)

	for _, target := range targets {
		info, err := os.Lstat(target)
		if err != nil {
			fmt.Printf("ls: cannot access '%s': %s\n", target, formatLsError(err))
			continue
		}

		if info.IsDir() {
			dirs = append(dirs, target)
			continue
		}

		entry, err := filesystem.SingleEntry(target)
		if err != nil {
			fmt.Printf("ls: cannot access '%s': %s\n", target, formatLsError(err))
			continue
		}
		entry.Name = target
		files = append(files, entry)
	}

	if len(files) == 0 && len(dirs) == 0 {
		return
	}

	if len(files) > 0 {
		cli.SortEntries(flags, files, 0)
		if flags.Reverse {
			cli.ReverseEntries(files, 0)
		}
		printEntries(files, flags, false)
	}

	if flags.Recursive {
		for i, dir := range dirs {
			if len(files) > 0 || i > 0 {
				fmt.Println()
			}
			recursive.ListRecursive(dir, flags)
		}
		return
	}

	for i, dir := range dirs {
		entries, err := filesystem.ListDirectory(dir)
		if err != nil {
			fmt.Printf("ls: cannot open directory '%s': %v\n", dir, err)
			continue
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

		needHeader := len(dirs) > 1 || len(files) > 0
		if needHeader {
			if i > 0 || len(files) > 0 {
				fmt.Println()
			}
			fmt.Printf("%s:\n", dir)
		}

		printEntries(entries, flags, true)
	}
}

func printEntries(entries []types.FileEntry, flags cli.Flags, showTotal bool) {
	if flags.Long {
		if showTotal {
			output.PrintLong(entries)
		} else {
			output.PrintLongNoTotal(entries)
		}
		return
	}
	output.PrintSimple(entries)
}

func formatLsError(err error) string {
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		var errno syscall.Errno
		if errors.As(pathErr.Err, &errno) {
			return capitalize(errno.Error())
		}
	}
	return capitalize(err.Error())
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}