package tui

import (
	"time"

	"github.com/Beeram12/gomonitor/pkg/checker"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
)

type viewState int

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

func initModel(urls []string) model {
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

type item string

func (i item) FilterValue() string { return "" }
func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
