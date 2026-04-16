package main

import (
	"errors"
	"os"
	"syscall"
	"testing"
)

func TestFormatLsError(t *testing.T) {
	// Test with a PathError
	pathErr := &os.PathError{Op: "stat", Path: "/nonexistent", Err: syscall.ENOENT}
	result := formatLsError(pathErr)
	expected := "No such file or directory"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test with a generic error
	genericErr := errors.New("some error")
	result = formatLsError(genericErr)
	expected = "Some error"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}