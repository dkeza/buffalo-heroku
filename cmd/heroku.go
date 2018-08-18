package cmd

import (
	"github.com/spf13/cobra"
)

// herokuCmd represents the heroku command
var herokuCmd = &cobra.Command{
	Use:   "heroku",
	Short: "description about this plugin",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(herokuCmd)
}