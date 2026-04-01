package main

func FilterEntries(entries []FileEntry, flags map[rune]bool) []FileEntry {
    if flags['a'] {
        return entries
    }

    var result []FileEntry
    for _, e := range entries {
        if len(e.Name) > 0 && e.Name[0] != '.' {
            result = append(result, e)
        }
    }
    return result
}