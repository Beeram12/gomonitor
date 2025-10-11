package tui

import (
	"sync"
	"time"

	"github.com/Beeram12/gomonitor/pkg/checker"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type viewState int
type statusUpdateMsg struct{ Status checker.Status }
type tickMsg time.Time

const (
	dashboardView viewState = iota
	addView
	removeView
)

type model struct {
	currentView viewState
	urls        []string
	statuses    map[string]checker.Status
	textInput   textinput.Model
	list        list.Model
	ticker      *time.Ticker
	quitting    bool
}

func InitialModel(urls []string) model {
	ti := textinput.New()
	ti.Placeholder = "https://example.com"
	ti.Focus()

	items := make([]list.Item, len(urls))
	for i, u := range urls {
		items[i] = item(u)
	}

	li := list.New(items, list.NewDefaultDelegate(), 0, 0)
	li.Title = "Which website do you want to remove?"

	return model{
		currentView: dashboardView,
		urls:        urls,
		statuses:    make(map[string]checker.Status),
		textInput:   ti,
		list:        li,
		ticker:      time.NewTicker(100 * time.Second),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		runChecks(m.urls),
		waitForTick(m.ticker),
	)

}

func runChecks(urls []string) tea.Cmd {
	return func() tea.Msg {
		var wg sync.WaitGroup
		for _, url := range urls {
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				status := checker.CheckEndPoint(url)
				Program.Send(statusUpdateMsg{Status: status})
			}(url)
		}
		wg.Wait()
		return nil
	}
}

func waitForTick(ticker *time.Ticker) tea.Cmd {
	return func() tea.Msg {
		<-ticker.C
		return tickMsg{}
	}
}

type item string

func (i item) FilterValue() string { return "" }
func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
