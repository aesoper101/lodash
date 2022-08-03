package arrayx

import "reflect"

// Pull Remove all the given values from array using SameValueZero for equality comparisons.
func Pull[T any](array []T, values ...T) []T {
	for _, value := range values {
		array = Remove(array, func(val T) bool {
			return reflect.DeepEqual(value, val)
		})
	}
	return array
}

// PullAll expacts that it accepts an array of values to remove.
func PullAll[T any](array []T, values []T) []T {
	for _, value := range values {
		array = Remove(array, func(val T) bool {
			return reflect.DeepEqual(value, val)
		})
	}

	return array
}

// PullAllBy excepts that it accepts an iteratee which is invoked for each element of array
// and values to generate the criterion by which they're compared. The iteratee is invoked with one argument: (value).
func PullAllBy[T any](array []T, values []T, iteratee func(T) any) []T {
	for _, value := range values {
		array = Remove(array, func(val T) bool {
			return reflect.DeepEqual(iteratee(value), iteratee(val))
		})
	}

	return array
}

// PullAllWith excepts that it accepts comparator which is invoked to compare elements of array to values.
// The comparator is invoked with two arguments: (arrVal, othVal).
func PullAllWith[T any](array []T, values []T, comparator func(T, T) bool) []T {
	for _, value := range values {
		array = Remove(array, func(val T) bool {
			return comparator(value, val)
		})
	}

	return array
}

// PullAt Remove elements from array corresponding to indexes and returns an array of removed elements.
// indexes is a slice of indexes of elements to remove.
func PullAt[T any](array []T, indexes ...int) []T {
	var result []T

	l := len(array)

	for _, v := range indexes {
		if v >= 0 && v < l {
			result = append(result, array[v])
		}
	}

	return result
}
