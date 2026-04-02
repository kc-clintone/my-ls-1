package main

import "fmt"

func PrintHeader(path string, show bool) {
	if show {
		fmt.Println(path + ":")
	}
}

func PrintError(err error) {
	fmt.Println("Error:", err.Error())
}