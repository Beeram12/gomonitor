package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [number]",
	Short: "Removes a monitored URL by its list number (see 'list' command)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		urlToRemove := args[0]

		fmt.Printf("Removing %s from the monitoring list...\n", urlToRemove)

		/*
				TODO:
			1. Convert numberStr to an integer.
			2. Load the list of URLs from the config file.
			3. Remove the URL at the specified index.
			4. Save the updated list back to the config file.
		*/
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
