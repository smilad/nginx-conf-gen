package cmd

import (
	"github.com/spf13/cobra"
)

var (
	seedCMD = cobra.Command{
		Use:  "seed database",
		Long: "seed database strucutures. This will seed tables",
		Run:  Runner.Seed,
	}
)

// seed database with fake data
func (c *command) Seed(cmd *cobra.Command, args []string) {

}
