package arrayx

import "reflect"

// IndexOf  Get the index at which the first occurrence of `value` is found in `arrayx`,
// If `value` is not found, -1 is returned.
// If `fromIndex` is provided and the index is outside of the arrayx bounds, -1 is returned.
func IndexOf[T any](array []T, value T, fromIndex ...int) int {
	index := 0
	if len(fromIndex) > 0 {
		if fromIndex[0] > len(array)-1 {
			return -1
		} else if fromIndex[0] > 0 {
			index = fromIndex[0]
		}
	}

	for i := index; i < len(array); i++ {
		if reflect.DeepEqual(array[i], value) {
			return i
		}
	}

	return -1
}

func LastIndexOf[T any](array []T, value T, fromIndex ...int) int {
	index := len(array) - 1
	if len(fromIndex) > 0 {
		if fromIndex[0] < 0 {
			index = 0
		} else if fromIndex[0] < len(array) {
			index = fromIndex[0]
		}
	}

	for i := index; i >= 0; i-- {
		if reflect.DeepEqual(array[i], value) {
			return i
		}
	}

	return -1
}
