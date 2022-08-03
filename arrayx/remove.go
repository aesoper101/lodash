package arrayx

// Remove removes all elements from `arrayx` that `predicate` returns truthy for and returns an arrayx of
// the removed elements. The predicate is invoked with three arguments: (value, index, arrayx).
func Remove[T any](arr []T, predicate func(T) bool) []T {
	for i := 0; i < len(arr); {
		if predicate(arr[i]) {
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			i++
		}
	}
	return arr
}

func RemoveAt[T any](arr []T, index int) []T {
	if index < 0 || index >= len(arr) {
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}
