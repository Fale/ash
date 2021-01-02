package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	repositoriesCmd.AddCommand(repositoriesListCmd)
}

var repositoriesListCmd = &cobra.Command{
	Use:   "list",
	Short: "list known ansible repositories",
	RunE:  repositoriesList,
}

func repositoriesList(cmd *cobra.Command, args []string) error {
	ws, err := getWorkspaceByName(viper.GetString("workspace"))
	if err != nil {
		return err
	}
	for _, repo := range ws.ListRepositories() {
		fmt.Printf("* %v\n", repo.URL)
	}
	return nil
}
