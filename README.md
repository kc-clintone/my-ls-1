# my-ls

**Rebuild the classic `ls` command in Go — from scratch.**

---

## Overview

`my-ls` is a Go implementation of the Unix `ls` command.

It lists files and directories, just like the original — with support for common flags and combinations.

---

## Features

### Supported Flags

- `-l` → Long format (permissions, size, owner, modification time, etc.)
- `-a` → Show all files (including hidden files starting with `.`)
- `-r` → Reverse the order of sorting
- `-t` → Sort by modification time (newest first)

Flags can be combined:

```bash
./my-ls -la
./my-ls -ltr
./my-ls -lat
```

---

## Usage

Building the project:

```bash
go build -o my-ls
```

Running the command:

```bash
./my-ls [flags] [directory]
```

### Examples

```bash
# List current directory (simple format)
./my-ls

# List all files including hidden
./my-ls -a

# Long format with details
./my-ls -l

# Long format, all files, reverse order
./my-ls -lar

# Sort by modification time
./my-ls -lt

# Combine multiple flags
./my-ls -la
./my-ls -ltr

# List a specific directory
./my-ls -la /path/to/dir
```

If no directory is provided, it defaults to the current directory (`.`).

---

## Project Structure

This project is organized with clean separation of concerns:

```
my-ls/
├── main.go                          # Entry point and orchestration
├── myls_test.go                     # Unit tests
├── go.mod                           # Go module definition
├── README.md                        # This file
│
└── internal/                        # Private implementation packages
    ├── cli/
    │   └── flags.go                 # Flag parsing and command-line logic
    │
    ├── filesystem/
    │   └── filesystem.go            # Directory reading and file metadata
    │
    ├── output/
    │   ├── long.go                  # Long format output (-l)
    │   └── simple.go                # Simple format output (default)
    │
    └── types/
        └── types.go                 # FileEntry struct and shared types
```

### Package Responsibilities

- **`main.go`** — Application entry point; orchestrates the workflow
- **`cli/flags.go`** — Parses command-line arguments and flags
- **`filesystem/filesystem.go`** — Reads directories, collects file metadata, handles special entries (`.` and `..`)
- **`output/long.go`** — Formats and prints files in long format
- **`output/simple.go`** — Formats and prints files in simple format
- **`types/types.go`** — Defines the `FileEntry` struct used throughout the application

The `internal/` directory structure prevents external packages from importing implementation details, promoting encapsulation.

---

## Building & Testing

### Build the executable:

```bash
go build -o my-ls
```

### Run all tests:

```bash
go test ./...
```

### Test a specific package:

```bash
go test ./internal/cli -v
go test ./internal/filesystem -v
```

---

## What You'll Learn

Working with this codebase will teach you:

- **Filesystem operations** — Working with `os.ReadDir()`, `os.Stat()`, and file metadata
- **CLI design** — Parsing flags and arguments cleanly
- **Code organization** — Using the `internal/` package pattern for clean architecture
- **Sorting & filtering** — Implementing different sort orders and filtering logic
- **Output formatting** — Matching real-world command output
- **Testing** — Writing and running unit tests in Go

---

## Implementation Details

### FileEntry Structure

Each file is represented as a `FileEntry` containing:

- Name, IsDir, Mode (permissions)
- Size, ModTime, Links (hard links)
- Owner, Group (from uid/gid lookup)
- SymlinkTo (for symbolic links)
- Blocks (disk usage in 512-byte blocks)

### Long Format Output

The long format matches `ls -l` output:

```
-rw-r--r-- 1 user group 1234 Apr  4 13:50 filename
^          ^ ^    ^     ^     ^   ^     ^     ^
├─ Mode    │ │    │     │     │   │     │     └─ Name
├─ Links ──┘ │    │     │     │   │     └─ Modification time
├─ Owner ────┘    │     │     │   └─ Size
├─ Group ─────────┘     │     └──── Month
                        └─ Total blocks
```

### Sorting

- **Alphabetical** (default, case-insensitive with uppercase as tiebreaker)
- **By modification time** (with `-t` flag)
- **Reversed** (with `-r` flag)

---

## Future Enhancements

- `-R` — Recursive directory listing
- `-S` — Sort by file size
- `-h` — Human-readable file sizes (e.g., "1.2K", "3.4M")
- `-1` — One entry per line
- Colored output for different file types

---

## Testing

The included test suite covers:

- Reading non-empty directories
- Handling empty directories
- File and directory detection

Run tests with:

```bash
go test ./... -v
```

---

## License

Educational project. Free to use and modify.
