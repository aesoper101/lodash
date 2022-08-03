package arrayx

// ForEach Iterates over elements of collectionx, executing the callback for each element.
func ForEach[T any](array []T, iteratee func(T, int)) {
	for i, v := range array {
		iteratee(v, i)
	}
}

// ForEachRight Iterates over elements of collectionx, executing the callback for each element in reverse order.
func ForEachRight[T any](array []T, iteratee func(T, int)) {
	for i := len(array) - 1; i >= 0; i-- {
		iteratee(array[i], i)
	}
}
