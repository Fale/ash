package main

import (
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	repositoriesCmd.AddCommand(repositoriesAddCmd)
}

var repositoriesAddCmd = &cobra.Command{
	Use:   "add URL",
	Short: "add an ansible repository",
	Args:  cobra.ExactArgs(1),
	RunE:  repositoriesAdd,
}

func repositoriesAdd(cmd *cobra.Command, args []string) error {
	ws, err := getWorkspaceByName(viper.GetString("workspace"))
	if err != nil {
		return err
	}
	u, err := url.Parse(args[0])
	return ws.AddRepositories([]url.URL{*u})
}
