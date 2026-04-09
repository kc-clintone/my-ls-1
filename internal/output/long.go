package output

import (
	"fmt"
	"os"
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

		// for device nodes, size column shows "major, minor"
		var sizeStr string
		isDevice := e.Mode&os.ModeDevice != 0 || e.Mode&os.ModeCharDevice != 0
		if isDevice {
			// match ls style spacing for "major, minor"
			sizeStr = fmt.Sprintf("%d, %3d", e.DeviceMajor, e.DeviceMinor)
		} else {
			sizeStr = strconv.FormatInt(e.Size, 10)
		}
		if s := len(sizeStr); s > maxSize {
			maxSize = s
		}
	}

	// Print total (1K blocks)
	fmt.Printf("total %d\n", total/2)

	// Time threshold (6 months)
	sixMonthsAgo := time.Now().AddDate(0, -6, 0)

	// Print entries
	for _, e := range entries {
		// compute leading type char like ls does (d, l, p, s, c, b or -)
		var typeChar byte = '-'
		switch {
		case e.Mode&os.ModeDir != 0:
			typeChar = 'd'
		case e.Mode&os.ModeSymlink != 0:
			typeChar = 'l'
		case e.Mode&os.ModeNamedPipe != 0:
			typeChar = 'p'
		case e.Mode&os.ModeSocket != 0:
			typeChar = 's'
		case (e.Mode & os.ModeDevice) != 0:
			if e.Mode&os.ModeCharDevice != 0 {
				typeChar = 'c'
			} else {
				typeChar = 'b'
			}
		default:
			typeChar = '-'
		}

		permFull := e.Mode.String()
		// Go's String() returns something like "Dcrw-------" for devices
		// We need to replace the first char with our computed typeChar
		if len(permFull) > 0 {
			// Skip the first character (which might be 'D' for devices)
			// and use only the permission bits
			if len(permFull) >= 10 {
				permFull = string(typeChar) + permFull[len(permFull)-9:]
			} else {
				permFull = string(typeChar) + permFull[1:]
			}
		} else {
			permFull = string(typeChar)
		}

		perm := permFull
		if e.HasXattr {
			perm += "+"
		}

		var dateStr string
		if e.ModTime.Before(sixMonthsAgo) {
			dateStr = e.ModTime.Format("Jan _2  2006")
		} else {
			dateStr = e.ModTime.Format("Jan _2 15:04")
		}

		// prepare size field as string (either "major, minor" or plain size)
		var sizeField string
		isDevice := e.Mode&os.ModeDevice != 0 || e.Mode&os.ModeCharDevice != 0
		if isDevice {
			sizeField = fmt.Sprintf("%d, %3d", e.DeviceMajor, e.DeviceMinor)
		} else {
			sizeField = strconv.FormatInt(e.Size, 10)
		}

		fmt.Printf(
			"%s %*d %-*s %-*s %*s %s %s",
			perm,
			maxLinks, e.Links,
			maxOwner, e.Owner,
			maxGroup, e.Group,
			maxSize, sizeField,
			dateStr,
			e.Name,
		)

		if e.SymlinkTo != "" {
			fmt.Printf(" -> %s", e.SymlinkTo)
		}

		fmt.Println()
	}
}
