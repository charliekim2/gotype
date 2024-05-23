package main

import (
	"fmt"
	"os"

	"github.com/charliekim2/gotype/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	words := []string{"Hello", "world!"}
	g := cmd.InitGame(words, 15)

	if _, err := tea.NewProgram(g).Run(); err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
}
