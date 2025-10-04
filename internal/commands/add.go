package commands

import (
	"fmt"
	"log"

	"github.com/Beeram12/gomonitor/pkg/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [url]",
	Short: "Add a new URL to monitor",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		urlToAdd := args[0]

		err := config.AddURL(urlToAdd)
		if err != nil {
			log.Fatalf("Error adding URL: %v", err)
		}
		fmt.Printf("Successfully added %s to the monitoring list.\n", urlToAdd)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
