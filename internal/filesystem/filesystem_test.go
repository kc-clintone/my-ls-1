package filesystem

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSingleEntry(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// Write some content
	_, err = tmpFile.WriteString("test content")
	if err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	// Test SingleEntry
	entry, err := SingleEntry(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if entry.Name != filepath.Base(tmpFile.Name()) {
		t.Errorf("Expected name %s, got %s", filepath.Base(tmpFile.Name()), entry.Name)
	}
	if entry.IsDir {
		t.Error("Expected file, got directory")
	}
	if entry.Size != 12 { // "test content" is 12 bytes
		t.Errorf("Expected size 12, got %d", entry.Size)
	}
}