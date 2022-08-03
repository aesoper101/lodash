package arrayx

// FindIndex is like Find, except that it returns the index of the first element
func FindIndex[T any](arr []T, predicate func(T) bool) int {
	for i, v := range arr {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// Find last element of arrayx passing predicate truth test.
// The predicate is invoked with three arguments: (value, index, arrayx).
func Find[T any](arr []T, predicate func(T) bool) T {
	var result T
	for _, v := range arr {
		if predicate(v) {
			result = v
			break
		}
	}
	return result
}

// FindLastIndex is like FindIndex, except that it iterates over elements of arrayx from right to left.
func FindLastIndex[T any](arr []T, predicate func(T) bool) (T, int, bool) {
	for i := len(arr) - 1; i >= 0; i-- {
		if predicate(arr[i]) {
			return arr[i], i, true
		}
	}
	var result T
	return result, -1, false
}

func FindIndexOf[T any](arr []T, predicate func(T) bool) (T, int, bool) {
	for i, v := range arr {
		if predicate(v) {
			return v, i, true
		}
	}
	var result T
	return result, -1, false
}

func FindOrCreate[T any](arr []T, predicate func(T) bool, creator func() T) (T, int, bool) {
	for i, v := range arr {
		if predicate(v) {
			return v, i, true
		}
	}
	var result = creator()
	arr = append(arr, result)
	return result, len(arr) - 1, false
}

func FindLast[T any](arr []T, predicate func(T) bool) T {
	var result T
	for i := len(arr) - 1; i >= 0; i-- {
		if predicate(arr[i]) {
			result = arr[i]
			break
		}
	}
	return result
}
