package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"todo/models"
	"todo/state"
	"todo/ui"
)

type model struct {
	todos  []models.Todo
	cursor int
}

func initialModel() model {
	return model{
		todos: []models.Todo{
			{Text: "jojo pose", Done: false},
			{Text: "do magic", Done: false},
			{Text: "buy winrar license", Done: false},
			{Text: "praise the sun", Done: false},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			m.cursor = state.UpdateCursor(m.cursor, len(m.todos), "up")

		case "down":
			m.cursor = state.UpdateCursor(m.cursor, len(m.todos), "down")

		case "enter", " ":
			m.todos = state.ToogleTodo(m.todos, m.cursor)
		}
	}

	return m, nil
}

func (m model) View() string {
	s := ""

	for i, todo := range m.todos {
		s += ui.Todo(todo, m.cursor == i)
	}

	return s
}

func main() {
	tea.NewProgram(initialModel()).Run()
}
