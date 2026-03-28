# my-ls

**Rebuild the classic `ls` command in Go — from scratch.**

No shortcuts. No cheating with `os/exec`. Just you, the filesystem, and a bit of curiosity.

---

## Overview

`my-ls` is a Go implementation of the Unix `ls` command.

It lists files and directories, just like the original — with support for common flags and combinations.


---

## Features

### Required Flags

* `-l` → Long format (permissions, size, owner, etc.)
* `-a` → Show hidden files
* `-r` → Reverse order
* `-t` → Sort by modification time

Flags can be combined:

```bash
./my-ls -la
./my-ls -ltr
```

---

## Usage

```bash
go run . [flags] [directory]
```

### Examples

```bash
# List current directory
go run .

# List all files including hidden
go run . -a

# Long format
go run . -l

# Combine flags
go run . -ltr /path/to/dir
```

If no directory is provided, it defaults to the current directory.

---

## Structure

```
my-ls/
├── main.go
├── flags.go
├── list.go
├── filter.go
├── sort.go
├── format.go
├── printer.go
├── types.go
├── utils.go
└── tests/
```

---

## What You’ll Learn

* Working with the filesystem in Go
* Parsing CLI arguments
* Sorting and filtering data
* Formatting output like a Unix tool
* Writing clean, testable code

---

## Testing

Create test cases for:

* Flag combinations
* Empty directories
* Hidden files
* Invalid paths

Run tests with:

```bash
go test ./...
```

---

## Bonus

* `-R` → Recursive listing
* Colored output (directories, executables, etc.)
* Add extra flags (`-S`, `-h`, etc.)

---

## Goal

By the end, you should be able to run:

```bash
./my-ls -lart
```

…and have it behave like the real thing.

If it doesn’t… well, the terminal will judge you silently.
