package list

import (
	"fmt"
	"log"

	"github.com/Beeram12/gomonitor/pkg/config"
)

func ListUrls() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	if len(cfg.URLs) == 0 {
		fmt.Println("No URLs are being monitored. Use 'gomonitor add <url>' to add one.")
		return
	}
	fmt.Println("---- Monitored Websites ----")
	for i, url := range cfg.URLs {
		fmt.Printf("%d: %s\n", i, url)
	}
	fmt.Println("--------------------------")
}
