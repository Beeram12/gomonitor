package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.quitting {
		return "Shutting down...\n"
	}
	switch m.currentView {
	case addView:
		return m.viewAdd()
	case removeView:
		return m.viewRemove()
	default:
		return m.viewDashboard()
	}
}

func (m model) viewDashboard() string {
	var b strings.Builder


	header := lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("%-50s | %s", "URL", "STATUS"))

	b.WriteString(header + "\n")
	b.WriteString(strings.Repeat("-", 65) + "\n")

	for _, url := range m.urls {
		status, ok := m.statuses[url]
		statusStr := "Checking...."
		style := lipgloss.NewStyle().Foreground(lipgloss.Color("245"))

		if ok {
			if status.IsUp {
				statusStr = fmt.Sprintf("UP (Code: %d)", status.StatusCode)
				style = lipgloss.NewStyle().Foreground(lipgloss.Color("40"))
			} else {
				statusStr = fmt.Sprintf("DOWN (Code: %d)", status.StatusCode)
				style = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
			}
		}
		b.WriteString(fmt.Sprintf("%-50s | %s\n", url, style.Render(statusStr)))
	}

	help := "\n'a' to add | 'r' to remove | 'q' to quit\n"
	b.WriteString(help)
	return b.String()
}

func (m model) viewAdd() string {
	return fmt.Sprintf(
		"Enter the new website URL to monitor:\n\n%s\n\n(esc to cancel)",
		m.textInput.View(),
	)
}

func (m model) viewRemove() string {
	return "\n" + m.list.View()
}
