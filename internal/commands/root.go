package commands

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomonitor",
	Short: "A lightweight and concurrent website uptime monitoring service.",
	Long:  `Will think about it later`,
	Run: func(cmd *cobra.Command, args []string) {
		// I will do something here but dont know for now
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
