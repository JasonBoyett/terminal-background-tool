package tui

import (
	"log"
	"os"

	files "github.com/JasonBoyett/terminal-background-tool/internal/files"
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
	config, err := files.LoadConfig()
	if err != nil {
		config = files.Config{}
	}
	return model{
		cursor:           0,
		config:           config,
		originalChoices:  options,
		choices:          options,
		selected:         make(map[int]struct{}),
		postRunContainer: config.PostRun,
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case start:
		return searchOffUpdate(m, msg)
	case searching:
		return searchingUpdate(m, msg)
	case selecting:
		return selectingUpdate(m, msg)
	case postRun:
		return postRunUpdate(m, msg)
	}
	return m, nil
}

func (m model) View() string {
	switch m.state {
	case start:
		return searchOffView(m)
	case searching:
		return searchingView(m)
	case selecting:
		return selectingView(m)
	case postRun:
		return postRunView(m)
	}
	return ""
}
