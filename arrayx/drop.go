package arrayx

// Drop Creates a slice of arrayx with n elements dropped from the beginning.
// num — The number of elements to drop. default — 1. if len(num) > 1, will drop num[1] elements from the beginning.
func Drop[T any](array []T, num ...uint) []T {
	n := 1
	if len(num) > 0 {
		n = int(num[0])
	}

	if n >= len(array) {
		return make([]T, 0)
	}

	results := make([]T, len(array)-n)
	for i := n; i < len(array); i++ {
		results[i-n] = array[i]
	}

	return results
}

// DropRight Creates a slice of arrayx with n elements dropped from the end.
func DropRight[T any](array []T, num ...uint) []T {
	n := 1
	if len(num) > 0 {
		n = int(num[0])
	}

	if n >= len(array) {
		return make([]T, 0)
	}

	results := make([]T, len(array)-n)
	for i := 0; i < len(array)-n; i++ {
		results[i] = array[i]
	}

	return results
}

// DropRightWhile Creates a slice of arrayx excluding elements dropped from the end.
// Elements are dropped until predicate returns false. The predicate is invoked with
// three arguments: (value, index, arrayx).
func DropRightWhile[T any](array []T, predicate func(T, uint) bool) []T {
	i := len(array) - 1
	for i >= 0 && predicate(array[i], uint(i)) {
		i--
	}

	result := make([]T, i+1)
	for j := 0; j <= i; j++ {
		result[j] = array[j]
	}

	return result
}

// DropWhile Creates a slice of arrayx excluding elements dropped from the beginning.
// Elements are dropped until predicate returns false.
// The predicate is invoked with three arguments: (value, index, arrayx).
func DropWhile[T any](array []T, predicate func(T, uint) bool) []T {
	i := 0
	for i < len(array) {
		if !predicate(array[i], uint(i)) {
			break
		}
		i++
	}

	result := make([]T, len(array)-i)
	for j := i; j < len(array); j++ {
		result[j-i] = array[j]
	}

	return result
}
