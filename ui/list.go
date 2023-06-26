package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"todo/models"
)

func List(rows []string) string {
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
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
