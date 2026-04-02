package main

import (
	"os"
	"time"
)

func formatMode(mode os.FileMode) string {
    var str string

    if mode.IsDir() {
        str += "d"
    } else {
        str += "-"
    }

    perms := []struct {
        r, w, x os.FileMode
    }{
        {0400, 0200, 0100},
        {0040, 0020, 0010},
        {0004, 0002, 0001},
    }

    for _, p := range perms {
        if mode&p.r != 0 {
            str += "r"
        } else {
            str += "-"
        }
        if mode&p.w != 0 {
            str += "w"
        } else {
            str += "-"
        }
        if mode&p.x != 0 {
            str += "x"
        } else {
            str += "-"
        }
    }

    return str
}

func formatTime(t time.Time) string {
    return t.Format("Jan _2 15:04")
}

func FormatName(e FileEntry) string {
	if e.SymlinkTo != "" {
		return e.Name + " -> " + e.SymlinkTo
	}
	return e.Name
}