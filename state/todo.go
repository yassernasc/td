package state

import "todo/models"

func ToogleTodo(todo *models.Todo) {
	(*todo).Done = !(*todo).Done
}

func MarkTodo(todo *models.Todo) {
	(*todo).Marked = !(*todo).Marked
}

func AddNewTodo(todos *[]models.Todo, msg string) {
	newTodo := models.Todo{Text: msg, Done: false}
	*todos = append(*todos, newTodo)
}

func EditTodo(todo *models.Todo, newText string) {
	(*todo).Text = newText
}

func RemoveMarkedTodos(todos *[]models.Todo) {
	var validTodos []models.Todo
	for _, todo := range *todos {
		if !todo.Marked {
			validTodos = append(validTodos, todo)
		}
	}

	*todos = validTodos
}
