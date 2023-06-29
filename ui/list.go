package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"td/models"
)

func List(rows []string) string {
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func Pending(todos []models.Todo) string {
	var texts []string
	for _, todo := range todos {
		texts = append(texts, todo.Text)
	}

	if len(texts) == 0 {
		return ""
	}

	return fmt.Sprintf("%s\n", lipgloss.JoinVertical(lipgloss.Left, texts...))
}
