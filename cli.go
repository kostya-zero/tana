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
		Use:   "set [key] [value]",
		Short: "Add a key to the database",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			SetCommand(args[0], args[1])
		},
	}

	getCmd := &cobra.Command{
		Use:   "get [key]",
		Short: "Get the value of the key",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			GetCommand(args[0])
		},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Get all keys with their values.",
		Run: func(cmd *cobra.Command, args []string) {
			ListCommand()
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete [key]",
		Short: "Delete key",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			DeleteCommand(args[0])
		},
	}

	updateCmd := &cobra.Command{
		Use:   "update [key] [newValue]",
		Short: "Update value of the key",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			UpdateCommand(args[0], args[1])
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show version of Tana",
		Run: func(cmd *cobra.Command, args []string) {
			println(VERSION)
		},
	}

	rootCmd.AddCommand(setCmd, getCmd, listCmd, deleteCmd, updateCmd, versionCmd)

	return rootCmd
}
