package commands

import (
	"fmt"
	"log"
	"sync"

	"github.com/Beeram12/gomonitor/pkg/checker"
	"github.com/Beeram12/gomonitor/pkg/config"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Performs a one-time status check on all monitored websites and prints a result table",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		var wg sync.WaitGroup

		fmt.Println("Checking Websites...")
		for _, i := range cfg.URLs {
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				status := checker.CheckEndPoint(u)
				if status.IsUp {
					fmt.Printf("%s is UP\n", i)
				} else {
					fmt.Printf("%s is DOWN\n", i)
				}
			}(i)
		}
		wg.Wait()
		fmt.Println("----Check Compelete----")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
