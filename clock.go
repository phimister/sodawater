package soda

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

type clockModel struct {
	tps      int
	deltat   float64
	count    uint64
	lastTime time.Time
}

func (c clockModel) tick() tea.Cmd {
	return tea.Tick(
		time.Second/time.Duration(c.tps),
		func(t time.Time) tea.Msg {
			return tickMsg(t)
		})
}

func (c clockModel) Start() (clockModel, tea.Cmd) {
	c.lastTime = time.Now()
	return c, c.tick()
}

func (c clockModel) Update(t tickMsg) (clockModel, tea.Cmd) {
	now := time.Time(t)
	c.deltat = now.Sub(c.lastTime).Seconds()
	c.lastTime = now
	c.count++
	return c, c.tick()
}

