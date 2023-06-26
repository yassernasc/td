package ui

import (
	"github.com/charmbracelet/lipgloss"
	"todo/models"
)

var baseStyle = lipgloss.NewStyle()
var focusedStyle = lipgloss.NewStyle().Underline(true)
var selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))

func computeStyle(focused bool, done bool) lipgloss.Style {
	lipgloss.JoinVertical(lipgloss.Left)

	if focused && done {
		return baseStyle.Copy().Inherit(focusedStyle).Inherit(selectedStyle)
	}

	if focused {
		return baseStyle.Copy().Inherit(focusedStyle)
	}

	if done {
		return baseStyle.Copy().Inherit(selectedStyle)
	}

	return baseStyle
}

func Todo(todo models.Todo, focused bool) string {
	style := computeStyle(focused, todo.Done)
	return style.Render(todo.Text)
}

func Todos(todos []models.Todo, cursor int) []string {
	var todoRows []string
	for i, todo := range todos {
		row := Todo(todo, cursor == i)
		todoRows = append(todoRows, row)
	}

	return todoRows
}
