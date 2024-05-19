package cmd

import tea "github.com/charmbracelet/bubbletea"

// Nested model holds menu, game
// Menu is genealized model? can be Main, Settings, Pause, etc.
// Not sure if results should be nested or just launch from the Game update
type model struct {
	width  int
	height int

	theme    Theme
	state    State
	settings Settings

	menu    tea.Model
	game    tea.Model
	restuls tea.Model
}

// Placeholder for now - use termenv for themeing (colors, font style, etc)
type Theme struct {
	background int
}

type Menu struct {
	selected MenuItem
}

type Game struct {
	words       []string
	currentWord int
	time        int
	wpm         float32
}

type Results struct {
	time int
	wpm  float32
}

type Settings struct {
	numWords int
}

type MenuItem int

const (
	PLAY MenuItem = iota
	QUIT
)

func (m MenuItem) String() string {
	return [...]string{"Play the game", "Quit to terminal"}[m]
}

type State int

const (
	MENU State = iota
	GAME
	RESULTS
)
