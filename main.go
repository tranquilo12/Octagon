package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {

	prog := progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))
	p := tea.NewProgram(initialModel(prog))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
