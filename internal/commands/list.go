package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists URL's monitored",
	Run: func(cmd *cobra.Command, args []string) {
		/*
			Listing URL's logic is implemented
		*/
		fmt.Printf("This lists all the URL's being monitored")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
