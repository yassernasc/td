package state

func UpdateCursor(cursor int, limit int, direction string) int {
	if direction == "up" {
		if cursor == 0 {
			return limit - 1
		}

		return cursor - 1
	}

	if direction == "down" {
		if cursor == limit-1 {
			return 0
		}

		return cursor + 1
	}

	return cursor
}
