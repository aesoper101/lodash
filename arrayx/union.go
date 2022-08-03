package arrayx

// Union Creates an arrayx of unique values, in order, from all given arrays using SameValueZero for equality comparisons.
func Union[T any](arrays ...[]T) []T {
	if len(arrays) == 0 {
		return []T{}
	}
	result := arrays[0]
	for _, array := range arrays {
		for _, item := range array {
			if !Contains(result, item) {
				result = append(result, item)
			}
		}
	}
	return result
}
