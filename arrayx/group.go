package arrayx

func GroupBy[T any, G comparable](arr []T, fn func(T) G) map[G][]T {
	result := make(map[G][]T)
	for _, v := range arr {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}
