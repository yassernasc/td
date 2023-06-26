package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"todo/models"
)

func List(todos []models.Todo, cursor int) string {
	var todoRows []string
	for i, todo := range todos {
		row := Todo(todo, cursor == i)
		todoRows = append(todoRows, row)
	}

	return lipgloss.NewStyle().MarginLeft(2).Render(lipgloss.JoinVertical(lipgloss.Left, todoRows...))
}

func Pending(todos []models.Todo) string {
	var pendingOnly []string
	for _, todo := range todos {
		if !todo.Done {
			pendingOnly = append(pendingOnly, todo.Text)
		}
	}

	if len(pendingOnly) == 0 {
		return "all done, congrats\n"
	}

	return fmt.Sprintf("%s\n", lipgloss.JoinVertical(lipgloss.Left, pendingOnly...))
}
