package output

import (
	"fmt"
	"strings"

	"myls/internal/types"
)

// PrintSimple prints file entries in simple format.
func PrintSimple(entries []types.FileEntry) {
	if len(entries) == 0 {
		return
	}
	names := make([]string, len(entries))
	for i, e := range entries {
		names[i] = e.Name
	}
	fmt.Println(strings.Join(names, "  "))
}
