package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type Game struct {
	words   []string
	time    timer.Model
	wpm     float32
	started bool
}

func (g Game) Init() tea.Cmd {
	return nil
}

func InitGame(words []string, duration int) *Game {
	game := Game{
		words:   words,
		time:    timer.New(time.Second * time.Duration(duration)),
		wpm:     0,
		started: false,
	}

	return &game
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		g.time, cmd = g.time.Update(msg)
		return g, cmd

	case timer.TimeoutMsg:
		return g, tea.Quit

	case tea.KeyMsg:
		if msg.String() == "q" {
			return g, tea.Quit
		} else if !g.started {
			g.started = true
			return g, g.time.Init()
		}
	}

	return g, nil
}

func (g Game) View() string {
	currentWords := strings.Join(g.words, " ")
	currentTime := g.time.View()
	currentWpm := g.wpm

	view := fmt.Sprintf("%s\nWPM: %.2f\n", currentWords, currentWpm) + currentTime
	return view
}
