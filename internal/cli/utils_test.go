package cli

import (
	"testing"

	"myls/internal/types"
)

func TestSortEntries(t *testing.T) {
	entries := []types.FileEntry{
		{Name: "b.txt"},
		{Name: "a.txt"},
		{Name: "c.txt"},
	}

	flags := Flags{}
	SortEntries(flags, entries, 0)

	// Should be sorted alphabetically
	expected := []string{"a.txt", "b.txt", "c.txt"}
	for i, e := range entries {
		if e.Name != expected[i] {
			t.Errorf("Expected %s at position %d, got %s", expected[i], i, e.Name)
		}
	}
}

func TestSortEntriesReverse(t *testing.T) {
	entries := []types.FileEntry{
		{Name: "a.txt"},
		{Name: "b.txt"},
		{Name: "c.txt"},
	}

	flags := Flags{Reverse: true}
	SortEntries(flags, entries, 0)
	ReverseEntries(flags, entries, 0)

	// Should be reverse sorted
	expected := []string{"c.txt", "b.txt", "a.txt"}
	for i, e := range entries {
		if e.Name != expected[i] {
			t.Errorf("Expected %s at position %d, got %s", expected[i], i, e.Name)
		}
	}
}