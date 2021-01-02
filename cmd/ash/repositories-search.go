package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v33/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	repositoriesCmd.AddCommand(repositoriesSearchCmd)
	repositoriesSearchCmd.Flags().Bool("add-all", false, "add all found repositories to the list")
}

var repositoriesSearchCmd = &cobra.Command{
	Use:   "search PLATFORM ORG_NAME",
	Short: "search new ansible repositories",
	Args:  cobra.ExactArgs(2),
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("add-all", cmd.Flags().Lookup("add-all"))
	},
	RunE: repositoriesSearch,
}

func repositoriesSearch(cmd *cobra.Command, args []string) error {
	var repositories []string
	switch args[0] {
	case "github":
		ctx := context.Background()
		client := github.NewClient(nil)
		opt := &github.RepositoryListByOrgOptions{
			ListOptions: github.ListOptions{PerPage: 10},
		}
		var allRepos []*github.Repository
		for {
			repos, resp, err := client.Repositories.ListByOrg(ctx, args[1], opt)
			if err != nil {
				return err
			}
			allRepos = append(allRepos, repos...)
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
		for _, repo := range allRepos {
			repositories = append(repositories, *repo.HTMLURL)
		}
	default:
		return fmt.Errorf("provider %v not recognised. Available providers are: github", args[0])
	}
	for _, repository := range repositories {
		fmt.Printf("* %v\n", repository)
		if viper.GetBool("add-all") {
			if err := repositoriesAdd(cmd, []string{repository}); err != nil {
				fmt.Printf("    skipped: %v\n", err)
			}
		}
	}
	return nil
}
