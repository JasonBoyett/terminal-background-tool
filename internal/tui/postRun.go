package tui

import (
	"github.com/JasonBoyett/terminal-background-tool/internal/files"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

func postRunUpdate(model model, msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.QuitMsg:
		return model, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			model.postRunContainer = ""
			model.state = start
			return model, nil
		case "esc":
			model.postRunContainer = ""
			model.state = start
			return model, nil
		case "enter":
			files.SaveConfig(model.config.BgDirectory, model.postRunContainer)
			model.state = start
			return model, nil
		case "backspace":
			if len(model.postRunContainer) > 0 {
				model.postRunContainer = model.postRunContainer[:len(model.postRunContainer)-1]
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
			model.postRunContainer += msg.String()
		}
	}
	return model, nil
}

func postRunView(m model) string {
	var (
		highlightedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f402c0"))
	)

	var str string
	if m.postRunContainer == "" {
		str += "enter your post run script here\n"

		str += "_________________\n"
		str += "Use " + highlightedStyle.Render("%T") + " in your script to represent your png image\n"
		str += "Use " + highlightedStyle.Render("%t") + " in your script to represent your png image\n"
		str += "Press " + highlightedStyle.Render("escape") + " to go back without changing your post run script.\n"
		str += "Press " + highlightedStyle.Render("control + c") + " to exit without changing your post run script.\n"
		str += "Press " + highlightedStyle.Render("enter") + " to set your post run script\n"
		return str + "\n"
	}

	str += "enter your post run script here\n"
	str += m.postRunContainer + "\n"
	str += "Use " + highlightedStyle.Render("%T") + " in your script to represent your png image\n"
	str += "Use " + highlightedStyle.Render("%t") + " in your script to represent your png image\n"
	str += "Press " + highlightedStyle.Render("escape or control + c") + " to exit without changing your post run script.\n"
	str += "Press " + highlightedStyle.Render("enter") + " to set your post run script\n"
	return str + "\n"
}
