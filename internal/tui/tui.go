package tui

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// The entry point for the TUI that selects the user background.
func EnterTui(options []string) {
	p := tea.NewProgram(initialModel(options))
	if _, err := p.Run(); err != nil {
		log.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}

func initialModel(options []string) model {
	return model{
		cursor:          0,
		originalChoices: options,
		choices:         options,
		selected:        make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.searchState {
	case off:
		return searchOffUpdate(m, msg)
	case searching:
		return searchingUpdate(m, msg)
	case selecting:
		return selectingUpdate(m, msg)
	}
	return m, nil
}

func (m model) View() string {
	switch m.searchState {
	case off:
		return searchOffView(m)
	case searching:
		return searchingView(m)
	case selecting:
		return selectingView(m)
	}
	return ""
}
