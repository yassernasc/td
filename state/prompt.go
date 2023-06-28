package state

import (
	"github.com/charmbracelet/bubbles/textinput"
	"todo/models"
)

func CreatePrompt() textinput.Model {
	input := textinput.New()
	input.Focus()

	return input
}

func SetPromptByTodo(input *textinput.Model, todo models.Todo) {
	(*input).SetValue(todo.Text)
}

func ExitPrompt(input *textinput.Model, mode *bool) {
	(*input).Reset()
	*mode = false
}
