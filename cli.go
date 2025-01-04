package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

type Exercise struct {
	Name string `yaml:"name"`
	Sets string `yaml:"sets"`
}

type Workout struct {
	Day       string     `yaml:"day"`
	Targets   []string   `yaml:"targets"`
	Exercises []Exercise `yaml:"exercises"`
}

type Program struct {
	Name     string    `yaml:"name"`
	Workouts []Workout `yaml:"workouts"`
}

type model struct {
	Program Program
	cursor  int
}

var (
	titleStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF69B4"))
	cursorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	workoutStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#1E90FF"))
)

func initialModel(filePath string) model {
	// Read and parse the YAML file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var program Program
	err = yaml.Unmarshal(data, &program)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return model{
		Program: program,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.Program.Workouts)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render(m.Program.Name+":") + "\n"

	for i, workout := range m.Program.Workouts {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}
		s += fmt.Sprintf("%s %s\n", cursorStyle.Render(cursor), workoutStyle.Render(workout.Day))
		if m.cursor == i {
			for _, exercise := range workout.Exercises {
				s += fmt.Sprintf("  - %s: %s\n", exercise.Name, exercise.Sets)
			}
		}
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("error: no program file provided")
	}
	filePath := os.Args[1]
	m := initialModel(filePath)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
