package tui

import helpers "github.com/JasonBoyett/terminal-background-tool/internal/helpers"

type searchState int

const (
	off searchState = iota
	searching
	selecting
)

type model struct {
	cursor          int
	originalChoices []string
	choices         []string
	selected        map[int]struct{}
	searchState     searchState
	searchPattern   string
}

func (m *model) selectChoice() string {
	return m.choices[m.cursor]
}

func (m *model) filterChoices(pattern string) {
	filteredChoices, err := helpers.FilterByRegexp(m.originalChoices, pattern)
	if err != nil {
		filteredChoices = m.originalChoices
	}
	m.choices = filteredChoices
	m.searchState = off
}

func (m *model) resetChoices() {
	if m.searchState == off {
		return
	}
	m.searchPattern = ""
	m.cursor = 0
	m.choices = m.originalChoices
	m.searchState = off
}
