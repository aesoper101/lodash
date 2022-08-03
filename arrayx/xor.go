package arrayx

func Xor[T comparable](array1 []T, array2 []T) []T {
	result := make([]T, 0)

	for _, value := range array1 {
		if !Contains(array2, value) {
			result = append(result, value)
		}
	}

	for _, value := range array2 {
		if !Contains(array1, value) {
			result = append(result, value)
		}
	}

	return result
}

func XorBy[T comparable](array1 []T, array2 []T, iteratee func(T) T) []T {
	result := make([]T, 0)

	for _, value := range array1 {
		if !Contains(array2, iteratee(value)) {
			result = append(result, value)
		}
	}

	for _, value := range array2 {
		if !Contains(array1, iteratee(value)) {
			result = append(result, value)
		}
	}

	return result
}

func XorWith[T comparable](array1 []T, array2 []T, comparator func(T, T) bool) []T {
	result := make([]T, 0)

	for _, value := range array1 {
		if !ContainsWith(array2, value, comparator) {
			result = append(result, value)
		}
	}

	for _, value := range array2 {
		if !ContainsWith(array1, value, comparator) {
			result = append(result, value)
		}
	}

	return result
}
