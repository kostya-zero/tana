package main

import (
	"fmt"
	"os"
)

func PrintError(msg string) {
	fmt.Fprintln(os.Stderr, " \x1b[91mError\x1b[0m:", msg)
}
