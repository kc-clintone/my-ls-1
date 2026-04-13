package output

import (
	"fmt"
	"strconv"
	"time"
	"os"

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
		perm := formatMode(e.Mode)

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

func formatMode(mode os.FileMode) string {
	var fileType byte = '-'
	switch {
	case mode&os.ModeDir != 0:
		fileType = 'd'
	case mode&os.ModeSymlink != 0:
		fileType = 'l'
	case mode&os.ModeNamedPipe != 0:
		fileType = 'p'
	case mode&os.ModeSocket != 0:
		fileType = 's'
	case mode&os.ModeDevice != 0 && mode&os.ModeCharDevice != 0:
		fileType = 'c'
	case mode&os.ModeDevice != 0:
		fileType = 'b'
	}

	out := []byte{
		fileType,
		'-', '-', '-',
		'-', '-', '-',
		'-', '-', '-',
	}

	if mode&0400 != 0 {
		out[1] = 'r'
	}
	if mode&0200 != 0 {
		out[2] = 'w'
	}
	if mode&0100 != 0 {
		out[3] = 'x'
	}
	if mode&0040 != 0 {
		out[4] = 'r'
	}
	if mode&0020 != 0 {
		out[5] = 'w'
	}
	if mode&0010 != 0 {
		out[6] = 'x'
	}
	if mode&0004 != 0 {
		out[7] = 'r'
	}
	if mode&0002 != 0 {
		out[8] = 'w'
	}
	if mode&0001 != 0 {
		out[9] = 'x'
	}

	// setuid/setgid/sticky bits
	if mode&os.ModeSetuid != 0 {
		if out[3] == 'x' {
			out[3] = 's'
		} else {
			out[3] = 'S'
		}
	}
	if mode&os.ModeSetgid != 0 {
		if out[6] == 'x' {
			out[6] = 's'
		} else {
			out[6] = 'S'
		}
	}
	if mode&os.ModeSticky != 0 {
		if out[9] == 'x' {
			out[9] = 't'
		} else {
			out[9] = 'T'
		}
	}

	return string(out)
}
