package main

import "fmt"

func PrintLong(entries []FileEntry) {
    for _, e := range entries {
        fmt.Printf("%s %d %s %s %6d %s %s",
            formatMode(e.Mode),
            e.Links,
            e.Owner,
            e.Group,
            e.Size,
            formatTime(e.ModTime),
            e.Name,
        )

        if e.SymlinkTo != "" {
            fmt.Printf(" -> %s", e.SymlinkTo)
        }

        fmt.Println()
    }
}