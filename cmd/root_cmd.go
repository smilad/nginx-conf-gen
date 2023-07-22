package cmd

import (
	"github.com/spf13/cobra"
	"nginx/app"
)

var (
	Runner     CommandLine = &command{}
	configFile             = ""
	debug      bool
)

type CommandLine interface {
	RootCmd() *cobra.Command
	Seed(cmd *cobra.Command, args []string)
}

type command struct {
}

// rootCmd will run the log streamer
var rootCmd = cobra.Command{
	Use:  "micro",
	Long: "A usecase that will validate restful transactions and send them to stripe.",
	Run: func(cmd *cobra.Command, args []string) {
		app.Start()
	},
}

// RootCmd will add flags and subcommands to the different commands
func (c *command) RootCmd() *cobra.Command {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "The configuration file")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "The usecase debug(true is production - false is dev)")

	// add more commands
	rootCmd.AddCommand(&seedCMD)
	return &rootCmd
}
