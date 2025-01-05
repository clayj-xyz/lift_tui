package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"gopkg.in/yaml.v3"
)

type CompletedSet struct {
	Reps   int `yaml:"reps"`
	Weight int `yaml:"weight"`
}

type SessionExercise struct {
	Name string         `yaml:"name"`
	Sets []CompletedSet `yaml:"sets"`
}

type Session struct {
	Date      time.Time         `yaml:"date"`
	Program   string            `yaml:"program"`
	Workout   int               `yaml:"workout"`
	Exercises []SessionExercise `yaml:"exercises"`
}

type sessionModel struct {
	Session Session
}

func createSessionModel(filePath string) sessionModel {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var session Session
	err = yaml.Unmarshal(data, &session)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return sessionModel{
		Session: session,
	}
}

func (m sessionModel) Init() tea.Cmd {
	return nil
}

func (m sessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m sessionModel) View() string {
	s := titleStyle.Render("Date: "+m.Session.Date.Format("2006-01-02")+":") + "\n"
	s += titleStyle.Render(fmt.Sprintf("Program: %s\n", m.Session.Program)) + "\n"

	for _, exercise := range m.Session.Exercises {
		s += workoutStyle.Render(fmt.Sprintf("%s:", exercise.Name)) + "\n"
		for _, set := range exercise.Sets {
			s += fmt.Sprintf("  %d reps @ %d lbs\n", set.Reps, set.Weight)
		}
	}

	s += "\nPress q to quit.\n"
	return s
}
