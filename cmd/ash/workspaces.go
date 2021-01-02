package main

import (
	"fmt"

	"github.com/fale/ash"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(workspacesCmd)
}

var workspacesCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "manage tracked ash workspaces",
}

type workspace struct {
	Name     string
	Location string
}

func getWorkspaceByName(name string) (*ash.Workspace, error) {
	for _, w := range viper.Get("workspaces").([]workspace) {
		if w.Name == name {
			return ash.NewWorkspace(w.Location)
		}
	}
	return nil, fmt.Errorf("the workspace %v does not exists", name)
}
