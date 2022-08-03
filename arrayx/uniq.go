package arrayx

func Uniq[T comparable](array []T) []T {
	if len(array) == 0 {
		return []T{}
	}
	result := []T{array[0]}
	for i := 1; i < len(array); i++ {
		if !Contains(result, array[i]) {
			result = append(result, array[i])
		}
	}
	return result
}

func UniqBy[T any, V comparable](array []T, iteratee func(T) V) []T {
	result := make([]T, 0)
	seen := make(map[V]bool)

	for _, value := range array {
		if _, ok := seen[iteratee(value)]; !ok {
			result = append(result, value)
			seen[iteratee(value)] = true
		}
	}

	return result
}

func UniqWith[T comparable](array []T, comparator func(T, T) bool) []T {
	result := make([]T, 0)

	for i := 0; i < len(array); i++ {
		if !ContainsWith(result, array[i], comparator) {
			result = append(result, array[i])
		}
	}

	return result
}

func Without[T comparable](array []T, values ...T) []T {
	result := make([]T, 0)

	for _, value := range array {
		if !Contains(values, value) {
			result = append(result, value)
		}
	}

	return result
}
