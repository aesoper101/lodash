package arrayx

import "math/rand"

func Sample[T any](arr []T) T {
	return arr[rand.Intn(len(arr))]
}

func SampleSize[T any](arr []T, n int) []T {
	size, result := len(arr), make([]T, 0)

	if size < n || n < 0 {
		return result
	}

	for i := 0; i < size && i < n; i++ {
		lenResult := size - i

		index := rand.Intn(lenResult)
		result = append(result, arr[index])

	}

	return result
}
