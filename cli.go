package main

import (
	"github.com/spf13/cobra"
)

func BuildCli() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "tana",
		Short: "No-nonsense key-value store for CLI",
		Run: func(cmd *cobra.Command, args []string) {
			println("Nothing to do. Use `tana --help` for available commands.")
		},
	}

	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Add a key to the database",
		Run:   SetCommand,
	}

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get the value of the key",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete key",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update value of the key",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.AddCommand(setCmd, getCmd, deleteCmd, updateCmd)

	return rootCmd
}
