package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"todo/models"
	"todo/state"
	"todo/ui"
)

type model struct {
	cursor     int
	exiting    bool
	input      textinput.Model
	insertMode bool
	todos      []models.Todo
}

func initialModel() model {
	return model{
		input: state.CreatePrompt(),
		todos: []models.Todo{
			{Text: "buy winrar license", Done: false},
			{Text: "praise the sun", Done: false},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.insertMode {
		m.input, _ = m.input.Update(msg)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				if m.input.Value() != "" {
					state.AddNewTodo(&m.todos, m.input.Value())
					m.input.Reset()
				}
			case "esc":
				m.cursor = len(m.todos) - 1
				state.ExitPrompt(&m.input, &m.insertMode)
			}
		}
	} else {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			case "esc", "ctrl+c", "q":
				m.exiting = true
				return m, tea.Quit

			case "up":
				state.UpdateCursor(&m.cursor, len(m.todos), "up")

			case "down":
				state.UpdateCursor(&m.cursor, len(m.todos), "down")

			case "enter", " ":
				state.ToogleTodo(&m.todos[m.cursor])

			case "a", "i":
				m.cursor = -1
				m.insertMode = true
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.exiting {
		return ui.Pending(m.todos)
	}

	list := ui.List(m.todos, m.cursor)

	if m.insertMode {
		return fmt.Sprintf("%s\n%s", list, m.input.View())
	}

	return list
}

func main() {
	tea.NewProgram(initialModel()).Run()
}
