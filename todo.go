package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"todo/models"
	"todo/state"
	"todo/ui"
)

type model struct {
	cursor     int
	editMode   bool
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
				state.CursorFocusOnLast(&m.cursor, m.todos)
				state.ExitPrompt(&m.input, &m.insertMode)
			}
		}
	} else if m.editMode {
		m.input, _ = m.input.Update(msg)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				if m.input.Value() != "" {
					state.EditTodo(&m.todos[m.cursor], m.input.Value())
					state.ExitPrompt(&m.input, &m.editMode)
				}
			case "esc":
				state.ExitPrompt(&m.input, &m.editMode)
			}
		}
	} else {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			case "esc", "q", "ctrl+c":
				m.exiting = true
				return m, tea.Quit

			case "up", "down":
				state.UpdateCursor(&m.cursor, len(m.todos), msg.String())

			case "enter", " ":
				if !m.todos[m.cursor].Marked {
					state.ToogleTodo(&m.todos[m.cursor])
				}

			case "d", "r":
				if !m.todos[m.cursor].Done {
					state.MarkTodo(&m.todos[m.cursor])
				}

			case "x":
				state.RemoveMarkedTodos(&m.todos)

			case "a", "i":
				m.cursor = -1
				m.insertMode = true

			case "e":
				m.editMode = true
				state.SetPromptByTodo(&m.input, m.todos[m.cursor])
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.exiting {
		return ui.Pending(m.todos)
	}

	todos := ui.Todos(m.todos, m.cursor)
	prompt := m.input.View()

	if m.editMode {
		todos[m.cursor] = prompt
	}

	if m.insertMode {
		todos = append(todos, prompt)
	}

	return ui.List(todos)
}

func main() {
	tea.NewProgram(initialModel()).Run()
}
