package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const cliBanner = `
██████╗ ██╗   ██╗██╗     ███████╗███████╗               
██╔══██╗██║   ██║██║     ██╔════╝██╔════╝               
██████╔╝██║   ██║██║     ███████╗█████╗                 
██╔═══╝ ██║   ██║██║     ╚════██║██╔══╝                 
██║     ╚██████╔╝███████╗███████║███████╗               
╚═╝      ╚═════╝ ╚══════╝╚══════╝╚══════╝               
                                                        
 ██████╗██╗  ██╗███████╗ ██████╗██╗  ██╗███████╗██████╗ 
██╔════╝██║  ██║██╔════╝██╔════╝██║ ██╔╝██╔════╝██╔══██╗
██║     ███████║█████╗  ██║     █████╔╝ █████╗  ██████╔╝
██║     ██╔══██║██╔══╝  ██║     ██╔═██╗ ██╔══╝  ██╔══██╗
╚██████╗██║  ██║███████╗╚██████╗██║  ██╗███████╗██║  ██║
 ╚═════╝╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝

			PulseCheck CLI
		Go-based Uptime Monitoring`


// rootCmd is base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "pulsecheck",
	Short: "PulseCheck is a website uptime monitoring tool.",
	Long: `A fast and concurrent uptime monitoring tool built with Go,
complete with a CLI, API, and a web dashboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cliBanner)
	},
}

func Execute(){
	err := rootCmd.Execute()
	if err != nil{
		os.Exit(1)
	}
}