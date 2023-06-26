package state

import "todo/models"

func ToogleTodo(todo *models.Todo) {
	(*todo).Done = !(*todo).Done
}

func AddNewTodo(todos *[]models.Todo, msg string) {
	newTodo := models.Todo{Text: msg, Done: false}
	*todos = append(*todos, newTodo)
}
