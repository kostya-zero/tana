package main

import "os"

func main() {
	cmd := BuildCli()
	if err := cmd.Execute(); err != nil {
		// PrintError("failed to prepare CLI interactions.")
		os.Exit(1)
	}
}
