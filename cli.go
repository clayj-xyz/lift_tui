package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	yaml "gopkg.in/yaml.v3"
)

type Exercise struct {
	Name string `yaml:"name"`
	Sets string `yaml:"sets"`
}

type Routine struct {
	Day       string     `yaml:"day"`
	Targets   []string   `yaml:"targets"`
	Exercises []Exercise `yaml:"exercises"`
}

type Program struct {
	Routine []Routine `yaml:"routine"`
}

type model struct {
	routines []Routine
	cursor   int
}

func initialModel() model {
	// Read and parse the YAML file
	data, err := os.ReadFile("program.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var program Program
	err = yaml.Unmarshal(data, &program)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return model{
		routines: program.Routine,
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
			if m.cursor < len(m.routines)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Exercise Routines:\n\n"

	for i, routine := range m.routines {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}
		s += fmt.Sprintf("%s %s\n", cursor, routine.Day)
		if m.cursor == i {
			for _, exercise := range routine.Exercises {
				s += fmt.Sprintf("  - %s: %s\n", exercise.Name, exercise.Sets)
			}
		}
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
