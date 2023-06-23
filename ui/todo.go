package ui

import (
	"fmt"
	"todo/models"
)

func Todo(todo models.Todo, focused bool) string {
	cursor := " " // no cursor
	if focused {
		cursor = ">"
	}

	checked := " "
	if todo.Done {
		checked = "x" // selected!
	}

	return fmt.Sprintf("%s [%s] %s\n", cursor, checked, todo.Text)
}
