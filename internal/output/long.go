package output

import (
	"fmt"
	"strconv"
	"time"

	"myls/internal/types"
)

// PrintLong prints file entries in long format.
func PrintLong(entries []types.FileEntry) {
	var total int64
	var maxLinks, maxOwner, maxGroup, maxSize int

	// Calculate totals and column widths
	for _, e := range entries {
		total += e.Blocks

		if l := len(strconv.FormatUint(e.Links, 10)); l > maxLinks {
			maxLinks = l
		}
		if len(e.Owner) > maxOwner {
			maxOwner = len(e.Owner)
		}
		if len(e.Group) > maxGroup {
			maxGroup = len(e.Group)
		}
		if s := len(strconv.FormatInt(e.Size, 10)); s > maxSize {
			maxSize = s
		}
	}

	// Print total (1K blocks)
	fmt.Printf("total %d\n", total/2)

	// Time threshold (6 months)
	sixMonthsAgo := time.Now().AddDate(0, -6, 0)

	// Print entries
	for _, e := range entries {
		perm := e.Mode.String()

		var dateStr string
		if e.ModTime.Before(sixMonthsAgo) {
			dateStr = e.ModTime.Format("Jan _2  2006")
		} else {
			dateStr = e.ModTime.Format("Jan _2 15:04")
		}

		fmt.Printf(
			"%s %*d %-*s %-*s %*d %s %s",
			perm,
			maxLinks, e.Links,
			maxOwner, e.Owner,
			maxGroup, e.Group,
			maxSize, e.Size,
			dateStr,
			e.Name,
		)

		if e.SymlinkTo != "" {
			fmt.Printf(" -> %s", e.SymlinkTo)
		}

		fmt.Println()
	}
}
