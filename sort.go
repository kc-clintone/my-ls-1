package main

import "sort"

func SortEntries(entries []FileEntry, flags map[rune]bool) {
    sort.Slice(entries, func(i, j int) bool {
        if flags['t'] {
            return entries[i].ModTime.After(entries[j].ModTime)
        }
        return entries[i].Name < entries[j].Name
    })

    if flags['r'] {
        reverse(entries)
    }
}

func reverse(entries []FileEntry) {
    for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
        entries[i], entries[j] = entries[j], entries[i]
    }
}