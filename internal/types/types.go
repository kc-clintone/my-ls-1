package types

import (
	"os"
	"time"
)

// FileEntry represents metadata about a file or directory.
type FileEntry struct {
	Name      string
	IsDir     bool
	Mode      os.FileMode
	Size      int64
	ModTime   time.Time
	Links     uint64
	Owner     string
	Group     string
	SymlinkTo string
	Blocks    int64
}
