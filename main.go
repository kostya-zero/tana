package main

import "os"

func main() {
	cmd := BuildCli()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
