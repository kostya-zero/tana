package main

import "fmt"

func PrintError(msg string) {
	fmt.Printf(" \x1b[91mError\x1b[0m: %s\n", msg)
}
