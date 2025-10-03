package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Performs a one-time status check on all monitored websites and prints a result table",
	Run: func(cmd *cobra.Command, args []string) {
		/*
			The logic for this is comming

		*/
		fmt.Printf("This command will check all the monitored websites")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
