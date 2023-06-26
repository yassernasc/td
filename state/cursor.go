package state

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
