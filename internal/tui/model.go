package tui

import (
	files "github.com/JasonBoyett/terminal-background-tool/internal/files"
	helpers "github.com/JasonBoyett/terminal-background-tool/internal/helpers"
)

type searchState int

const (
	start searchState = iota
	searching
	selecting
	postRun
)

type model struct {
	config           files.Config
	cursor           int
	originalChoices  []string
	choices          []string
	selected         map[int]struct{}
	state            searchState
	searchPattern    string
	postRunContainer string
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
	m.state = start
}

func (m *model) resetChoices() {
	if m.state == start {
		return
	}
	m.searchPattern = ""
	m.cursor = 0
	m.choices = m.originalChoices
	m.state = start
}
