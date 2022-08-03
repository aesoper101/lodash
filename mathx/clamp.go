package mathx

func Clamp[T Number](n, min, max T) T {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}
