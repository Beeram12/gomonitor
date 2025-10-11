package tui

import (
	"log"

	"github.com/Beeram12/gomonitor/pkg/config"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.currentView {
	case addView:
		return m.updateAdd(msg)
	case removeView:
		return m.updateRemove(msg)
	default:
		return m.updateDashboard(msg)
	}
}

func (m model) updateDashboard(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			m.ticker.Stop()
			return m, tea.Quit
		case "a":
			m.currentView = addView
			return m, m.textInput.Focus()
		case "r":
			m.currentView = removeView
			return m, nil
		}
	case tickMsg:
		return m, tea.Batch(waitForTick(m.ticker), runChecks(m.urls))
	case statusUpdateMsg:
		m.statuses[msg.Status.URL] = msg.Status
		return m, nil
	}
	return m, nil
}

func (m model) updateAdd(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			urlToAdd := m.textInput.Value()
			if urlToAdd != "" {
				if err := config.AddURL(urlToAdd); err != nil {
					log.Printf("Could not add URL: %v", err)
				}
				cfg, _ := config.LoadConfig()
				return InitialModel(cfg.URLs), tea.Batch(runChecks(cfg.URLs), waitForTick(m.ticker))
			}
		case tea.KeyEsc:
			m.currentView = dashboardView
			m.textInput.Blur()
			m.textInput.SetValue("")
			return m, nil
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) updateRemove(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":

			selectedIndex := m.list.Index()

			if err := config.RemoveURLByValue(selectedIndex); err != nil {
				log.Printf("Could not remove URL: %v", err)
			}

			cfg, _ := config.LoadConfig()
			return InitialModel(cfg.URLs), tea.Batch(runChecks(cfg.URLs), waitForTick(m.ticker))

		case "esc":
			m.currentView = dashboardView
			return m, nil
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
