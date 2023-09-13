package selector

import (
	"bufio"
	"fmt"
	"log"
	"os"

	files "github.com/JasonBoyett/terminal-background-tool/internal/files"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
}

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
		cursor:   0,
		choices:  options,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.QuitMsg:
		return m, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
    case "enter":
      files.SetBg(m.choices[m.cursor])
      return m, tea.Quit
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	standardStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E0E0E0"))
	selectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#b007ed"))
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
			line := choice

			if i == m.cursor {
				line = selectedStyle.Render(line)
			} else {
				line = standardStyle.Render(line)
			}

			str += line + "\n"
		}
	}

	return str + "\nPress q to quit.\n"
}

func Setup() error {
  
  var path string 

  fmt.Println("Let's get your image folder set up")
  fmt.Println("Where would you like your image folder to be?")
  fmt.Println("Please provide a path")

  reader := bufio.NewScanner(os.Stdin)
  if reader.Scan(){
    path = reader.Text()
  }
  if err := files.SaveConfig(path); err != nil {
    return err
  }

  return nil
}
