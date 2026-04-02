package main

import "os"

func main() {
	flags, paths, err := ParseArgs(os.Args[1:])
	if err != nil {
		PrintError(err)
		return
	}

	Run(flags, paths)
}
