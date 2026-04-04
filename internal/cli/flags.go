package cli

import (
	"strings"
)

// Flags holds the command-line flag values.
type Flags struct {
	Long       bool
	All        bool
	Reverse    bool
	TimeSort   bool
	Recursive  bool
}

// ParseFlags parses command-line arguments and returns flags and the target path.
func ParseFlags(args []string) (Flags, string) {
	flags := Flags{
		Long:      false,
		All:       false,
		Reverse:   false,
		TimeSort:  false,
		Recursive: false,
	}
	path := "."

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			for _, ch := range arg[1:] {
				switch ch {
				case 'l':
					flags.Long = true
				case 'a':
					flags.All = true
				case 'r':
					flags.Reverse = true
				case 't':
					flags.TimeSort = true
				case 'R':
					flags.Recursive = true
				}
			}
		} else {
			path = arg
		}
	}

	return flags, path
}
