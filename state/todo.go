package state

import "github.com/yassernasc/td/models"

func ToogleTodo(todo *models.Todo) {
	if !todo.Marked {
		(*todo).Done = !(*todo).Done
	}
}

func MarkTodo(todo *models.Todo) {
	if !todo.Done {
		(*todo).Marked = !(*todo).Marked
	}
}

func AddNewTodo(todos *[]models.Todo, msg string) {
	if msg != "" {
		newTodo := models.Todo{Text: msg, Done: false}
		*todos = append(*todos, newTodo)
	}
}

func EditTodo(todo *models.Todo, newText string) {
	if newText != "" {
		(*todo).Text = newText
	}
}

func FilterPendingTodos(todos []models.Todo) []models.Todo {
	var pendingOnly []models.Todo
	for _, todo := range todos {
		if !todo.Done && !todo.Marked {
			pendingOnly = append(pendingOnly, todo)
		}
	}
	return pendingOnly
}

func CleanTodos(todos *[]models.Todo) {
	*todos = FilterPendingTodos(*todos)
}
