package main

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(repositoriesCmd)
}

var repositoriesCmd = &cobra.Command{
	Use:   "repositories",
	Short: "manage tracked ansible repositories",
}
