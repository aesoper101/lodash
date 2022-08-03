package arrayx

func Reduce[T any](a []T, f func(T, T) T, initial T) T {
	if len(a) == 0 || f == nil {
		var vv T
		return vv
	}

	l := len(a) - 1

	reduce := func(a []T, ff func(T, T) T, memo T, startPoint, direction, length int) T {
		result := memo
		index := startPoint

		for i := 0; i <= length; i++ {
			result = ff(result, a[index])
			index += direction
		}

		return result
	}

	return reduce(a, f, initial, 0, 1, l)
}
