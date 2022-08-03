package mathx

func InRange[T Number](n T, start, end T) bool {
	return n >= start && n <= end
}
