package arrayx

// Reject returns the elements of arrayx that `predicate` does not return truthy for.
func Reject[T any](array []T, predicate func(T, int) bool) []T {
	result := make([]T, 0)
	for i, v := range array {
		if !predicate(v, i) {
			result = append(result, v)
		}
	}
	return result
}
