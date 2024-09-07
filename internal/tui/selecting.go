package tui

import (
	files "github.com/JasonBoyett/terminal-background-tool/internal/files"
	helpers "github.com/JasonBoyett/terminal-background-tool/internal/helpers"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

func selectingUpdate(model model, msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.QuitMsg:
		return model, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return model, tea.Quit
		case "up", "k":
			if model.cursor > 0 {
				model.cursor--
			}
		case "enter":
			err := files.SetBg(model.selectChoice())
			if err != nil {
				return model, tea.Quit
			}
			return model, tea.Quit
		case "r":
			err := files.RandomBgFromOpts(model.choices)
			if err != nil {
				return model, tea.Quit
			}
			return model, tea.Quit
		case "down", "j":
			if model.cursor < len(model.choices)-1 {
				model.cursor++
			}
		case "esc":
			model.resetChoices()
		case "/":
			model.cursor = 0
			model.searchState = searching
		}
	}

	return model, nil
}

func selectingView(m model) string {
	standardStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#05e2ff"))
	selectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffce1f"))
	highlightedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#f402c0"))
	viewLen := 5

	str := "Which background would you like to use?\n\n"

	start := m.cursor - viewLen
	end := m.cursor + viewLen

	if start < 0 {
		start = 0
	}

	if end >= len(m.choices) {
		end = len(m.choices) - 1
	}

	for i, choice := range m.choices {
		if i >= start && i <= end {
			line, err := helpers.TrimFileExtension(choice)
			if err != nil {
				line = choice
			}

			if i == m.cursor {
				line = selectedStyle.Render(line)
			} else {
				line = standardStyle.Render(line)
			}

			str += line + "\n"
		}
	}

	str += "_________________\n"
	str += "Press " + highlightedStyle.Render("enter") + " to select.\n"
	str += "Press " + highlightedStyle.Render("r") + " to select a random option.\n"
	str += "Press " + highlightedStyle.Render("/") + " to change your search.\n"
	str += "Press " + highlightedStyle.Render("esc") + " to exit search mode.\n"

	return str + "Press " + highlightedStyle.Render("q") + " to quit.\n"
}
