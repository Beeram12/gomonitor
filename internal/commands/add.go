package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)


var addCmd = &cobra.Command{
	Use: "Add [url]",
	Short: "Add a new URL to monitor",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		urlToAdd := args[0]
		fmt.Printf("Adding %s to monitoring list...\n",urlToAdd)
		//  Here, you would call your core logic to save the URL to your config file.
	},
}

func init(){
	rootCmd.AddCommand(addCmd)
}