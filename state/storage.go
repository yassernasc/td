package state

import (
	"encoding/json"
	"github.com/adrg/xdg"
	"github.com/yassernasc/td/models"
	"os"
	"path/filepath"
)

var (
	dirPath  = filepath.Join(xdg.DataHome, "td")
	filePath = filepath.Join(dirPath, "todos.json")
)

func Load() []models.Todo {
	var todos []models.Todo
	data, err := os.ReadFile(filePath)
	if err == nil {
		json.Unmarshal(data, &todos)
	}
	return todos
}

func Save(todos []models.Todo) {
	// create dir if it not exists by ignoring errors
	os.Mkdir(dirPath, os.ModePerm)

	json, _ := json.MarshalIndent(todos, "", "  ")
	os.WriteFile(filePath, json, 0666)
}
