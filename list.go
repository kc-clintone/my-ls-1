package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
	"time"
)

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
}

func ListDirectory(path string, flags map[rune]bool) ([]FileEntry, error) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := make([]FileEntry, 0, len(dirEntries))

	for _, entry := range dirEntries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		// stat := info.Sys().(*syscall.Stat_t)
		stat, ok := info.Sys().(*syscall.Stat_t)
		if !ok {
			continue
		}

		ownerObj, _ := user.LookupId(fmt.Sprint(stat.Uid))
		groupObj, _ := user.LookupGroupId(fmt.Sprint(stat.Gid))

		var symlinkTo string
		if info.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(filepath.Join(path, entry.Name()))
			if err == nil {
				symlinkTo = target
			}
		}

		fileEntry := FileEntry{
			Name:      entry.Name(),
			IsDir:     entry.IsDir(),
			Mode:      info.Mode(),
			Size:      info.Size(),
			ModTime:   info.ModTime(),
			Links:     stat.Nlink,
			Owner:     ownerObj.Username,
			Group:     groupObj.Name,
			SymlinkTo: symlinkTo,
		}

		result = append(result, fileEntry)
	}

	return result, nil
}
