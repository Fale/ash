package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use: "ash",
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ash/config.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "explain what is being done")
	rootCmd.PersistentFlags().StringP("workspace", "w", "default", "set the desired workspace")
	if err := viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("workspace", rootCmd.PersistentFlags().Lookup("workspace")); err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(fmt.Sprintf("%v/.ash", home))
		viper.SetConfigName("config")
	}

	viper.SetDefault("workspaces", []workspace{{Name: "default", Location: fmt.Sprintf("%v/.ash/default", home)}})
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("verbose") {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}
