package arrayx

// Tail gets all but the first element of `arrayx`. if length is 1 or less, returns an empty arrayx.
func Tail[T any](arr []T) []T {
	if len(arr) <= 1 {
		return []T{}
	}

	return arr[1:]
}

// Take creates a slice of `arrayx` with `n` elements taken from the beginning.
// If `n` is less than or equal to 0, an empty arrayx is returned.
// If `n` is greater than the length of arrayx, arrayx is returned.
func Take[T any](array []T, n int) []T {
	if n <= 0 {
		return []T{}
	}

	if n >= len(array) {
		return array
	}

	return array[:n]
}

// TakeRight creates a slice of `arrayx` with `n` elements taken from the end.
func TakeRight[T any](array []T, n int) []T {
	if n <= 0 {
		return []T{}
	}

	if n >= len(array) {
		return array
	}

	return array[len(array)-n:]
}

// TakeWhile creates a slice of `arrayx` with elements taken from the beginning until the predicate returns false.
// The predicate is invoked with three arguments: (value, index, arrayx).
func TakeWhile[T any](array []T, predicate func(T) bool) []T {
	for i, v := range array {
		if !predicate(v) {
			return array[:i]
		}
	}

	return array
}

// TakeRightWhile creates a slice of `arrayx` with elements taken from the end until the predicate returns false.
// The predicate is invoked with three arguments: (value, index, arrayx).
func TakeRightWhile[T any](array []T, predicate func(T) bool) []T {
	for i := len(array) - 1; i >= 0; i-- {
		if !predicate(array[i]) {
			return array[i+1:]
		}
	}

	return array
}
