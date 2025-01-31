package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	// Add fields to hold the state of your application.
}

type msg struct {
	// Add fields to represent the different types of messages your application can handle.
}

func (m model) Init() tea.Cmd {
	// Initialize your application and return any commands you want to run.
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle the different types of messages and return the updated model and any commands you want to run.
	fmt.Println("hello"
	return m, nil
}

func (m model) View() string {
	// Return a string representation of your model.
	return ""
}

func main() {
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		// Handle any errors that occurred and exit.
		os.Exit(1)
	}
}
