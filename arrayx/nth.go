package arrayx

import "math"

//Nth Get the element at index `n` of `arrayx`. If `n` is negative, the nth element from the end is returned.
func Nth[T any](array []T, n int) T {
	if n < 0 {
		n += len(array)
	}

	var nilValue T

	if n < 0 {
		return nilValue
	}
	if math.Abs(float64(n)) >= float64(len(array)) {
		return nilValue
	}
	return array[n]
}

func First[T any](array []T) (T, bool) {
	var result T
	if len(array) == 0 {
		return result, false
	}
	return array[0], true
}

func Last[T any](array []T) (T, bool) {
	var result T
	if len(array) == 0 {
		return result, false
	}
	return array[len(array)-1], true
}
