package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF69B4"))
	cursorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	workoutStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#1E90FF"))
	exerciseStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700"))
	daysOfWeek    = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
)

func main() {
	viewProgram := flag.String("view-program", "", "View the program")
	viewSession := flag.String("view-session", "", "View the session")
	flag.Parse()

	if *viewProgram != "" {
		m := createProgramModel(*viewProgram)
		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	} else if *viewSession != "" {
		m := createSessionModel(*viewSession)
		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Please specify either --view-program or --view-session")
		os.Exit(1)
	}
}
