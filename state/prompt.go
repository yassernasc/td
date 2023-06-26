package state

import "github.com/charmbracelet/bubbles/textinput"

func CreatePrompt() textinput.Model {
	input := textinput.New()
	input.Focus()
	return input
}

func ExitPrompt(input *textinput.Model, insertMode *bool) {
	(*input).Reset()
	*insertMode = false
}
