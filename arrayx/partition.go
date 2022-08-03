package arrayx

func Partition[T any](arr []T, predicate func(T) bool) [][]T {
	var truthy, falsy []T
	for _, v := range arr {
		if predicate(v) {
			truthy = append(truthy, v)
		} else {
			falsy = append(falsy, v)
		}
	}
	return [][]T{truthy, falsy}
}
