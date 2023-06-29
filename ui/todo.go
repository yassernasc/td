package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/yassernasc/td/models"
)

var (
	greenColor   = lipgloss.AdaptiveColor{Light: "40", Dark: "43"}
	redColor     = lipgloss.AdaptiveColor{Light: "203", Dark: "210"}
	grayColor    = lipgloss.AdaptiveColor{Light: "240", Dark: "245"}
	focusedStyle = lipgloss.NewStyle().Underline(true)
	doneStyle    = lipgloss.NewStyle().Foreground(grayColor).Strikethrough(true)
	markedStyle  = lipgloss.NewStyle().Foreground(redColor)
	checkMark    = lipgloss.NewStyle().SetString("✓").Foreground(greenColor).PaddingRight(1).String()
	xMark        = lipgloss.NewStyle().SetString("✗").Foreground(redColor).PaddingRight(1).String()
)

func applyStyle(base *lipgloss.Style, flavor lipgloss.Style) {
	*base = (*base).Inherit(flavor)
}

func getStyle(todo models.Todo, focused bool) lipgloss.Style {
	style := lipgloss.NewStyle()

	if todo.Done {
		applyStyle(&style, doneStyle)
	}

	if focused {
		applyStyle(&style, focusedStyle)
	}

	if todo.Marked {
		applyStyle(&style, markedStyle)
	}

	return style
}

func getPrefix(todo models.Todo) string {
	if todo.Done {
		return checkMark
	}

	if todo.Marked {
		return xMark
	}

	return lipgloss.NewStyle().PaddingLeft(2).String()
}

func Todo(todo models.Todo, focused bool) string {
	style := getStyle(todo, focused)
	prefix := getPrefix(todo)
	return prefix + style.Render(todo.Text)
}

func Todos(todos []models.Todo, cursor int) []string {
	var todoRows []string
	for i, todo := range todos {
		row := Todo(todo, cursor == i)
		todoRows = append(todoRows, row)
	}

	return todoRows
}
