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
	todos := state.Load()

	var initialInsertMode bool
	if len(todos) == 0 {
		initialInsertMode = true
	}

	return model{
		input:      state.CreatePrompt(),
		todos:      todos,
		insertMode: initialInsertMode,
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
				state.AddNewTodo(&m.todos, m.input.Value())
				m.input.Reset()

				// cursor shows up when starts with no item and then one is added
				m.cursor = -1

			case "esc":
				state.CursorFocusOnLast(&m.cursor, m.todos)
				state.ExitPrompt(&m.input, &m.insertMode)

				// nothing to show, exit
				if len(m.todos) == 0 {
					return m, tea.Quit
				}
			}
		}
	} else if m.editMode {
		m.input, _ = m.input.Update(msg)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				state.EditTodo(&m.todos[m.cursor], m.input.Value())
				state.ExitPrompt(&m.input, &m.editMode)

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

			case "shift+up":
				state.ShiftUp(&m.cursor, &m.todos)

			case "shift+down":
				state.ShiftDown(&m.cursor, &m.todos)

			case "enter", " ":
				state.ToogleTodo(&m.todos[m.cursor])

			case "d", "r", "x":
				state.MarkTodo(&m.todos[m.cursor])

			case "c":
				state.CleanTodos(&m.todos)
				state.EnsureCursorIsVisible(&m.cursor, m.todos)

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
		pending := state.FilterPendingTodos(m.todos)
		state.Save(pending)
		return ui.Pending(pending)
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
