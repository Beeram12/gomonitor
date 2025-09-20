package main

import (
	"fmt"
	"os"
	"github.com/Beeram12/gomonitor/internal/checker"
	"github.com/spf13/cobra"
)

const cliBanner = `
██████╗  ██████╗     ███╗   ███╗ ██████╗ ███╗   ██╗██╗████████╗ ██████╗ ██████╗ 
██╔════╝ ██╔═══██╗    ████╗ ████║██╔═══██╗████╗  ██║██║╚══██╔══╝██╔═══██╗██╔══██╗
██║  ███╗██║   ██║    ██╔████╔██║██║   ██║██╔██╗ ██║██║   ██║   ██║   ██║██████╔╝
██║   ██║██║   ██║    ██║╚██╔╝██║██║   ██║██║╚██╗██║██║   ██║   ██║   ██║██╔══██╗
╚██████╔╝╚██████╔╝    ██║ ╚═╝ ██║╚██████╔╝██║ ╚████║██║   ██║   ╚██████╔╝██║  ██║
 ╚═════╝  ╚═════╝     ╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝
                                                                                 
                  🚀 Lightweight Website Status Checker in Go 🚀`

type listCommands struct {
	urls []string
}

var rootCmd = &cobra.Command{
	Use:   "gomonitor",
	Short: "GoMonitor is a website uptime monitoring tool.",
	Long: `A fast and concurrent uptime monitoring tool built with Go,
complete with a CLI, API, and a web dashboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cliBanner)
	},
}

// Subcommands

var checkCmd = &cobra.Command{
	Use: "check [url]",
	Short: "Check a website",
	Run: func(cmd *cobra.Command, args []string){
		checker.CheckStatus(args)
	},
}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
