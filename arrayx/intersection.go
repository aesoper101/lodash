package arrayx

import "reflect"

func InterSection[T any](arrays ...[]T) []T {
	results := make([]T, 0)
	if len(arrays) == 0 {
		return results
	}

	results = arrays[0]

	for i := 1; i < len(arrays); i++ {
		for j := 0; j < len(results); j++ {
			if !Contains(arrays[i], results[j]) {
				results = Remove(results, func(value T) bool {
					return reflect.DeepEqual(value, results[j])
				})
				j--
			}
		}
	}

	return results
}
