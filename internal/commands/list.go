package commands

import (
	"github.com/Beeram12/gomonitor/pkg/list"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists URL's monitored",
	Run: func(cmd *cobra.Command, args []string) {
		list.ListUrls()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
