package arrayx

// Filter  Creates a slice of arrayx with all elements that pass the predicate truth test.
func Filter[T any](array []T, predicate func(T, int) bool) []T {
	result := make([]T, 0)
	for i, v := range array {
		if predicate(v, i) {
			result = append(result, v)
		}
	}
	return result
}
