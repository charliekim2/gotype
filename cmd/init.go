package cmd

import (
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (m model) initGameState(filename string) {
	dat, err := os.ReadFile(filename)
	check(err)

	words := strings.Split(string(dat), "\n")
	m.state = GAME
}

func initialModel(width int, height int, numWords int) model {
	return model{
		width:  width,
		height: height,

		theme: Theme{
			background: 0,
		},
		state: MENU,
		settings: Settings{
			numWords: numWords,
		},
	}
}
