package filesystem

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"syscall"

	"myls/internal/types"
)

// ListDirectory reads a directory and returns information about all entries.
func ListDirectory(path string) ([]types.FileEntry, error) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := make([]types.FileEntry, 0, len(dirEntries))

	for _, entry := range dirEntries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		stat, ok := info.Sys().(*syscall.Stat_t)
		if !ok {
			continue
		}

		ownerObj, _ := user.LookupId(fmt.Sprint(stat.Uid))
		groupObj, _ := user.LookupGroupId(fmt.Sprint(stat.Gid))

		ownerName := fmt.Sprint(stat.Uid)
		if ownerObj != nil {
			ownerName = ownerObj.Username
		}

		groupName := fmt.Sprint(stat.Gid)
		if groupObj != nil {
			groupName = groupObj.Name
		}

		var symlinkTo string
		if info.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(filepath.Join(path, entry.Name()))
			if err == nil {
				symlinkTo = target
			}
		}

		fileEntry := types.FileEntry{
			Name:      entry.Name(),
			IsDir:     entry.IsDir(),
			Mode:      info.Mode(),
			Size:      info.Size(),
			ModTime:   info.ModTime(),
			Links:     stat.Nlink,
			Owner:     ownerName,
			Group:     groupName,
			SymlinkTo: symlinkTo,
			Blocks:    stat.Blocks,
		}

		result = append(result, fileEntry)
	}

	return result, nil
}

// CreateSpecialEntry creates a FileEntry for . or .. entries.
func CreateSpecialEntry(path, name string) types.FileEntry {
	fullPath := path
	if name == ".." {
		fullPath = filepath.Join(path, "..")
	}

	info, err := os.Stat(fullPath)
	if err != nil {
		return types.FileEntry{Name: name}
	}

	stat, _ := info.Sys().(*syscall.Stat_t)

	owner := fmt.Sprint(stat.Uid)
	group := fmt.Sprint(stat.Gid)

	if u, err := user.LookupId(owner); err == nil {
		owner = u.Username
	}
	if g, err := user.LookupGroupId(group); err == nil {
		group = g.Name
	}

	return types.FileEntry{
		Name:    name,
		IsDir:   info.IsDir(),
		Mode:    info.Mode(),
		Size:    info.Size(),
		ModTime: info.ModTime(),
		Links:   stat.Nlink,
		Owner:   owner,
		Group:   group,
		Blocks:  stat.Blocks,
	}
}

// SingleEntry returns FileEntry for a single file path
func SingleEntry(path string) (types.FileEntry, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return types.FileEntry{}, err
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return types.FileEntry{}, fmt.Errorf("stat error")
	}

	ownerObj, _ := user.LookupId(fmt.Sprint(stat.Uid))
	groupObj, _ := user.LookupGroupId(fmt.Sprint(stat.Gid))

	ownerName := fmt.Sprint(stat.Uid)
	if ownerObj != nil {
		ownerName = ownerObj.Username
	}

	groupName := fmt.Sprint(stat.Gid)
	if groupObj != nil {
		groupName = groupObj.Name
	}

	var symlinkTo string
	if info.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(path)
		if err == nil {
			symlinkTo = target
		}
	}

	return types.FileEntry{
		Name:      filepath.Base(path),
		IsDir:     info.IsDir(),
		Mode:      info.Mode(),
		Size:      info.Size(),
		ModTime:   info.ModTime(),
		Links:     stat.Nlink,
		Owner:     ownerName,
		Group:     groupName,
		SymlinkTo: symlinkTo,
		Blocks:    stat.Blocks,
	}, nil
}