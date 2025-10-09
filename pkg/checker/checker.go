package checker

import (
	"net/http"
	"time"
)

type Status struct {
	URL        string
	IsUp       bool
	StatusCode int
}

func CheckEndPoint(url string) Status {
	client := http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Get(url)
	if err != nil {
		return Status{
			URL:        url,
			IsUp:       false,
			StatusCode: 0,
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return Status{
			URL:        url,
			IsUp:       true,
			StatusCode: resp.StatusCode,
		}
	}

	return Status{
		URL:        url,
		IsUp:       false,
		StatusCode: resp.StatusCode,
	}
}
