package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"gopkg.in/yaml.v3"
)

type Exercise struct {
	Name string `yaml:"name"`
	Sets string `yaml:"sets"`
}

type Workout struct {
	Name      string     `yaml:"name"`
	Weekday   string     `yaml:"weekday"`
	Exercises []Exercise `yaml:"exercises"`
}

type Program struct {
	Name     string    `yaml:"name"`
	StartDay string    `yaml:"start_day"`
	Workouts []Workout `yaml:"workouts"`
}

type programModel struct {
	Program Program
	cursor  int
}

func createProgramModel(filePath string) programModel {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var program Program
	err = yaml.Unmarshal(data, &program)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var startDayIndex int
	for i, day := range daysOfWeek {
		if day == program.StartDay {
			startDayIndex = i
			break
		}
	}
	for i, workout := range program.Workouts {
		workout.Weekday = daysOfWeek[(startDayIndex+i)%7]
		program.Workouts[i] = workout
	}

	return programModel{
		Program: program,
	}
}

func (m programModel) Init() tea.Cmd {
	return nil
}

func (m programModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m programModel) View() string {
	s := titleStyle.Render(m.Program.Name+":") + "\n"

	for i, workout := range m.Program.Workouts {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}
		s += fmt.Sprintf(
			"%s %s - %s\n",
			cursorStyle.Render(cursor),
			workoutStyle.Render(workout.Weekday),
			workoutStyle.Render(workout.Name),
		)
		if m.cursor == i {
			for _, exercise := range workout.Exercises {
				s += fmt.Sprintf("  - %s: %s\n", exercise.Name, exercise.Sets)
			}
		}
	}

	s += "\nPress q to quit.\n"
	return s
}
