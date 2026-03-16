package soda

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	initializing state = iota
	ready
)

type Model struct {
	clock  clockModel
	state  state
	width  int
	height int
}

func NewModel(opts ...Option) Model {
	m := Model{clock: clockModel{tps: 24}}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil

	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	// case tea.MouseMsg:
	// 	# do something

	case tea.WindowSizeMsg:
		if m.state == initializing {
			m.width = msg.Width
			m.height = msg.Height
			m.state = ready
			m.clock, cmd = m.clock.Start()
		}

	case tickMsg:
		m.clock, cmd = m.clock.Update(msg)
	}

	return m, cmd
}

func (m Model) View() string {
	if m.state == initializing {
		return "Loading..."
	}
	return fmt.Sprintf(
		"Window size of width %d, height %d, FPS of %.02f",
		m.width,
		m.height,
		1 / m.clock.deltat,
	)
}

func (m Model) Run() {
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
