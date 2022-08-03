package arrayx

// Reverse arrayx so that the first element becomes the last, the second element becomes the second to last, and so on.
func Reverse[T any](array []T) []T {
	for left, right := 0, len(array)-1; left < right; left, right = left+1, right-1 {
		array[left], array[right] = array[right], array[left]
	}
	return array
}
