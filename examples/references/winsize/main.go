package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	log string
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return tea.Tick(1, //1*time.Second,
	func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.log += fmt.Sprintf("[WindowResizeMsg] width: %d, height %d\n", msg.Width, msg.Height)

	case tickMsg:
		m.log += fmt.Sprintf("[tickMsg] time: %d\n", time.Time(msg).Second())

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("\n%s\n", m.log)
}

func main() {
	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Printf("[error] %v", err)
		os.Exit(1)
	}
}
