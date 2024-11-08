package tui

import (
	helpers "github.com/JasonBoyett/terminal-background-tool/internal/helpers"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

func searchingUpdate(model model, msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.QuitMsg:
		return model, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return model, tea.Quit
		case "esc":
			model.resetChoices()
			break
		case "enter":
			model.cursor = 0
			model.state = selecting
			break
		case "backspace":
			if len(model.searchPattern) > 0 {
				model.searchPattern = model.searchPattern[:len(model.searchPattern)-1]
			} else {
				break
			}
			break
		default:
			// if the key pressed is a special key other than enter, backspace, or escape
			// then we don't want to add it to the pattern so it will be ignored
			if len(msg.String()) > 1 {
				break
			}
			model.searchPattern += msg.String()
			filteredChoices, err := helpers.FilterByRegexp(
				model.originalChoices,
				model.searchPattern,
			)
			if err != nil {
				model.choices = model.originalChoices
				break
			}
			model.choices = filteredChoices
		}
	}
	return model, nil
}

func searchingView(m model) string {
	standardStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#05e2ff"))
	searchStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#f7f496"))
	highlightedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#f402c0"))

	var str string
	if m.searchPattern == "" {
		str += searchStyle.Render("Searching for: ") + m.searchPattern + "\n"
		str += "search is case sensitive and supports regex\n"

		str += "_________________\n"
		str += "Press " + highlightedStyle.Render("escape") + " to exit search mode.\n"
		str += "Press " + highlightedStyle.Render("control + c") + " to exit program.\n"
		return str + "\n"
	}

	viewLen := 5
	start := m.cursor - viewLen
	end := m.cursor + viewLen

	if start < 0 {
		start = 0
	}
	if end >= len(m.choices) {
		end = len(m.choices) - 1
	}

	str += "Searching for: " + m.searchPattern + "\n"

	for i, choice := range m.choices {

		if i >= start && i <= end {
			line, err := helpers.TrimFileExtension(choice)
			if err != nil {
				line = choice
			}
			str += standardStyle.Render(line)
			str += "\n"
		}
	}

	str += "_________________\n"
	str += "Press " + highlightedStyle.Render("enter") + " to enter selection mode.\n"
	str += "Press " + highlightedStyle.Render("escape") + " to exit search mode.\n"
	str += "Press " + highlightedStyle.Render("control + c") + " to exit program.\n"

	return str
}
