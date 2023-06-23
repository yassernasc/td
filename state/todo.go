package state

import "todo/models"

func ToogleTodo(todos []models.Todo, cursor int) []models.Todo {
	todos[cursor].Done = !todos[cursor].Done
	return todos
}
