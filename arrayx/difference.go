package arrayx

import "reflect"

// Difference creates a new arrayx with all of the elements that are not present in the other arrays.
func Difference[T any](array1, array2 []T) []T {
	result := make([]T, 0)
	for _, value := range array1 {
		if !Contains(array2, value) {
			result = append(result, value)
		}
	}
	return result
}

func DifferenceBy[T any](array1, array2 []T, iteratee func(T) any) []T {
	result := make([]T, 0)
	for _, value := range array1 {
		if !ContainsBy(array2, value, iteratee) {
			result = append(result, value)
		}
	}
	return result
}

func DifferenceWith[T any](array1, array2 []T, comparator func(T, T) bool) []T {
	result := make([]T, 0)
	for _, value := range array1 {
		if !ContainsWith(array2, value, comparator) {
			result = append(result, value)
		}
	}
	return result
}

// Contains determines whether or not a given value is present in a given arrayx.
func Contains[T any](array []T, value T) bool {
	for _, v := range array {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

func ContainsBy[T any](array []T, value T, iteratee func(T) any) bool {
	for _, v := range array {
		if reflect.DeepEqual(iteratee(v), iteratee(value)) {
			return true
		}
	}
	return false
}

func ContainsWith[T any](array []T, value T, comparator func(T, T) bool) bool {
	for _, v := range array {
		if comparator(v, value) {
			return true
		}
	}
	return false
}
