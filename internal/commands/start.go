package commands

import (
	"fmt"
	"log"

	"github.com/Beeram12/gomonitor/internal/tui"
	"github.com/Beeram12/gomonitor/pkg/config"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the interactive monitoring dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Could not load config: %v", err)
		}

		initialModel := tui.InitialModel(cfg.URLs)
		p := tea.NewProgram(initialModel)
		tui.Program = p
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running program: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
