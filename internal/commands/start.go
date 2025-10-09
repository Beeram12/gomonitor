package commands

import "github.com/spf13/cobra"

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the interactive monitoring dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		/*
			Code for interacting with TUI
		*/
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
