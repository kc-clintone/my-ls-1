package cli

import (
	"reflect"
	"testing"
)

func TestParseFlags_NoArgs(t *testing.T) {
	flags, paths := ParseFlags([]string{})
	expectedFlags := Flags{Long: false, All: false, Reverse: false, TimeSort: false, Recursive: false}
	expectedPaths := []string{"."}

	if !reflect.DeepEqual(flags, expectedFlags) {
		t.Errorf("Expected flags %v, got %v", expectedFlags, flags)
	}
	if !reflect.DeepEqual(paths, expectedPaths) {
		t.Errorf("Expected paths %v, got %v", expectedPaths, paths)
	}
}

func TestParseFlags_WithLongFlag(t *testing.T) {
	flags, paths := ParseFlags([]string{"-l"})
	expectedFlags := Flags{Long: true, All: false, Reverse: false, TimeSort: false, Recursive: false}
	expectedPaths := []string{"."}

	if !reflect.DeepEqual(flags, expectedFlags) {
		t.Errorf("Expected flags %v, got %v", expectedFlags, flags)
	}
	if !reflect.DeepEqual(paths, expectedPaths) {
		t.Errorf("Expected paths %v, got %v", expectedPaths, paths)
	}
}

func TestParseFlags_WithMultipleFlags(t *testing.T) {
	flags, paths := ParseFlags([]string{"-la"})
	expectedFlags := Flags{Long: true, All: true, Reverse: false, TimeSort: false, Recursive: false}
	expectedPaths := []string{"."}

	if !reflect.DeepEqual(flags, expectedFlags) {
		t.Errorf("Expected flags %v, got %v", expectedFlags, flags)
	}
	if !reflect.DeepEqual(paths, expectedPaths) {
		t.Errorf("Expected paths %v, got %v", expectedPaths, paths)
	}
}

func TestParseFlags_WithPaths(t *testing.T) {
	flags, paths := ParseFlags([]string{"-l", "/tmp", "/home"})
	expectedFlags := Flags{Long: true, All: false, Reverse: false, TimeSort: false, Recursive: false}
	expectedPaths := []string{"/tmp", "/home"}

	if !reflect.DeepEqual(flags, expectedFlags) {
		t.Errorf("Expected flags %v, got %v", expectedFlags, flags)
	}
	if !reflect.DeepEqual(paths, expectedPaths) {
		t.Errorf("Expected paths %v, got %v", expectedPaths, paths)
	}
}