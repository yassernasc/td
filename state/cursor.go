package state

import "github.com/yassernasc/td/models"

func UpdateCursor(cursor *int, limit int, direction string) {
	if direction == "up" {
		if *cursor == 0 {
			*cursor = limit - 1
		} else {
			*cursor--
		}
	}

	if direction == "down" {
		if *cursor == limit-1 {
			*cursor = 0
		} else {
			*cursor++
		}
	}
}

func CursorFocusOnLast(cursor *int, todos []models.Todo) {
	*cursor = len(todos) - 1
}

func ShiftUp(cursor *int, todos *[]models.Todo) {
	if *cursor == 0 {
		*todos = append((*todos)[1:], (*todos)[:1]...)
		*cursor = len(*todos) - 1
	} else {
		(*todos)[*cursor], (*todos)[*cursor-1] = (*todos)[*cursor-1], (*todos)[*cursor]
		*cursor--
	}
}

func ShiftDown(cursor *int, todos *[]models.Todo) {
	limit := len(*todos) - 1
	if *cursor == limit {
		*todos = append((*todos)[limit:], (*todos)[:limit]...)
		*cursor = 0
	} else {
		(*todos)[*cursor], (*todos)[*cursor+1] = (*todos)[*cursor+1], (*todos)[*cursor]
		*cursor++
	}
}

func EnsureCursorIsVisible(cursor *int, todos []models.Todo) {
	limit := len(todos) - 1
	if *cursor > limit {
		*cursor = limit
	}
}
