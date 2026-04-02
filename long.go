package main

import "fmt"

func PrintLong(entries []FileEntry) {
	var maxLinks, maxOwner, maxGroup, maxSize int

	for _, e := range entries {
		maxLinks = max(maxLinks, len(fmt.Sprint(e.Links)))
		maxOwner = max(maxOwner, len(e.Owner))
		maxGroup = max(maxGroup, len(e.Group))
		maxSize = max(maxSize, len(fmt.Sprint(e.Size)))
	}

	for _, e := range entries {
		fmt.Printf(
			"%s %*d %-*s %-*s %*d %s %s\n",
			FormatPermissions(e.Mode),
			maxLinks, e.Links,
			maxOwner, e.Owner,
			maxGroup, e.Group,
			maxSize, e.Size,
			formatTime(e.ModTime),
			FormatName(e),
		)
	}
}