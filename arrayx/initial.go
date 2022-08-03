package arrayx

func Initial[T any](array []T) []T {
	if len(array) == 0 {
		return array
	}
	if len(array) == 1 {
		return array[:0]
	}
	return array[:len(array)-1]
}
