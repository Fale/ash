package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	workspacesCmd.AddCommand(workspacesListCmd)
}

var workspacesListCmd = &cobra.Command{
	Use:   "list",
	Short: "list known ash workspaces",
	RunE:  workspacesList,
}

func workspacesList(cmd *cobra.Command, args []string) error {
	for _, w := range viper.Get("workspaces").([]workspace) {
		fmt.Printf("* %v (%v)\n", w.Name, w.Location)
	}
	return nil
}
