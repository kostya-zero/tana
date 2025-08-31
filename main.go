package main

func main() {
	cmd := BuildCli()
	if err := cmd.Execute(); err != nil {
		println("Failed to prepare CLI interactions")
	}
}
